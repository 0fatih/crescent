package amm

import (
	"fmt"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ OrderSource = (*OrderBook)(nil)

// OrderBook is an order book.
type OrderBook struct {
	buys, sells orderBookTicks
}

// NewOrderBook returns a new OrderBook.
func NewOrderBook(orders ...Order) *OrderBook {
	ob := &OrderBook{}
	ob.Add(orders...)
	return ob
}

// Add adds orders to the order book.
func (ob *OrderBook) Add(orders ...Order) {
	for _, order := range orders {
		switch order.GetDirection() {
		case Buy:
			ob.buys.add(order)
		case Sell:
			ob.sells.add(order)
		}
	}
}

// HighestBuyPrice returns the highest buy price in the order book.
func (ob *OrderBook) HighestBuyPrice() (sdk.Dec, bool) {
	return ob.buys.highestPrice()
}

// LowestSellPrice returns the lowest sell price in the order book.
func (ob *OrderBook) LowestSellPrice() (sdk.Dec, bool) {
	return ob.sells.lowestPrice()
}

// BuyAmountOver returns the amount of buy orders in the order book
// for price greater or equal than given price.
func (ob *OrderBook) BuyAmountOver(price sdk.Dec) sdk.Int {
	return ob.buys.amountOver(price)
}

// SellAmountUnder returns the amount of sell orders in the order book
// for price less or equal than given price.
func (ob *OrderBook) SellAmountUnder(price sdk.Dec) sdk.Int {
	return ob.sells.amountUnder(price)
}

// BuyOrdersOver returns buy orders in the order book for price greater
// or equal than given price.
func (ob *OrderBook) BuyOrdersOver(price sdk.Dec) []Order {
	return ob.buys.ordersOver(price)
}

// SellOrdersUnder returns sell orders in the order book for price less
// or equal than given price.
func (ob *OrderBook) SellOrdersUnder(price sdk.Dec) []Order {
	return ob.sells.ordersUnder(price)
}

// orderBookTicks represents a list of orderBookTick.
// This type is used for both buy/sell sides of OrderBook.
type orderBookTicks []*orderBookTick

func (ticks orderBookTicks) findPrice(price sdk.Dec) (i int, exact bool) {
	i = sort.Search(len(ticks), func(i int) bool {
		return ticks[i].price.LTE(price)
	})
	if i < len(ticks) && ticks[i].price.Equal(price) {
		exact = true
	}
	return
}

func (ticks *orderBookTicks) add(order Order) {
	i, exact := ticks.findPrice(order.GetPrice())
	if exact {
		(*ticks)[i].add(order)
	} else {
		if i < len(*ticks) {
			// Insert a new order book tick at index i.
			*ticks = append((*ticks)[:i], append([]*orderBookTick{newOrderBookTick(order)}, (*ticks)[i:]...)...)
		} else {
			// Append a new order book tick at the end.
			*ticks = append(*ticks, newOrderBookTick(order))
		}
	}
}

func (ticks orderBookTicks) highestPrice() (sdk.Dec, bool) {
	if len(ticks) == 0 {
		return sdk.Dec{}, false
	}
	for _, tick := range ticks {
		if TotalOpenAmount(tick.orders).IsPositive() {
			return tick.price, true
		}
	}
	return sdk.Dec{}, false
}

func (ticks orderBookTicks) lowestPrice() (sdk.Dec, bool) {
	if len(ticks) == 0 {
		return sdk.Dec{}, false
	}
	for i := len(ticks) - 1; i >= 0; i-- {
		if TotalOpenAmount(ticks[i].orders).IsPositive() {
			return ticks[i].price, true
		}
	}
	return sdk.Dec{}, false
}

func (ticks orderBookTicks) amountOver(price sdk.Dec) sdk.Int {
	i, exact := ticks.findPrice(price)
	if !exact {
		i--
	}
	amt := sdk.ZeroInt()
	for ; i >= 0; i-- {
		amt = amt.Add(TotalOpenAmount(ticks[i].orders))
	}
	return amt
}

func (ticks orderBookTicks) amountUnder(price sdk.Dec) sdk.Int {
	i, _ := ticks.findPrice(price)
	amt := sdk.ZeroInt()
	for ; i < len(ticks); i++ {
		amt = amt.Add(TotalOpenAmount(ticks[i].orders))
	}
	return amt
}

func (ticks orderBookTicks) ordersOver(price sdk.Dec) []Order {
	i, exact := ticks.findPrice(price)
	if !exact {
		i--
	}
	var orders []Order
	for ; i >= 0; i-- {
		orders = append(orders, ticks[i].orders...)
	}
	return orders
}

func (ticks orderBookTicks) ordersUnder(price sdk.Dec) []Order {
	i, _ := ticks.findPrice(price)
	var orders []Order
	for ; i < len(ticks); i++ {
		orders = append(orders, ticks[i].orders...)
	}
	return orders
}

// orderBookTick represents a tick in OrderBook.
type orderBookTick struct {
	price  sdk.Dec
	orders []Order
}

func newOrderBookTick(order Order) *orderBookTick {
	return &orderBookTick{
		price:  order.GetPrice(),
		orders: []Order{order},
	}
}

func (tick *orderBookTick) add(order Order) {
	if !order.GetPrice().Equal(tick.price) {
		panic(fmt.Sprintf("order price %q != tick price %q", order.GetPrice(), tick.price))
	}
	if first := tick.orders[0]; first.GetDirection() != order.GetDirection() {
		panic(fmt.Sprintf("order direction %q != tick direction %q", order.GetDirection(), first.GetDirection()))
	}
	tick.orders = append(tick.orders, order)
}
