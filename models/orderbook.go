package models

import (
	"hash"
	"hash/crc32"

	"github.com/shopspring/decimal"
)

// The bids and asks are formatted like so:
// [[best price, size at price], [next next best price, size at price], ...]
//
// Checksum
// Every message contains a signed 32-bit integer checksum of the orderbook.
// You can run the same checksum on your client orderbook state and compare it to checksum field.
// If they are the same, your client's state is correct.
// If not, you have likely lost or mishandled a packet and should re-subscribe to receive the initial snapshot.
//
// The checksum operates on a string that represents the first 100 orders on the orderbook on either side. The format of the string is:
//
// <best_bid_price>:<best_bid_size>:<best_ask_price>:<best_ask_size>:<second_best_bid_price>:<second_best_ask_price>:...
// For example, if the orderbook was comprised of the following two bids and asks:
//
// bids: [[5000.5, 10], [4995.0, 5]]
// asks: [[5001.0, 6], [5002.0, 7]]
// The string would be '5005.5:10:5001.0:6:4995.0:5:5002.0:7'
//
// If there are more orders on one side of the book than the other, then simply omit the information about orders that don't exist.
//
// For example, if the orderbook had the following bids and asks:
//
// bids: [[5000.5, 10], [4995.0, 5]]
// asks: [[5001.0, 6]]
// The string would be '5005.5:10:5001.0:6:4995.0:5'
//
// The final checksum is the crc32 value of this string.
type OrderBook struct {
	Asks     [][]decimal.Decimal `json:"asks"`
	Bids     [][]decimal.Decimal `json:"bids"`
	Checksum int64               `json:"checksum,omitempty"`
	Time     FTXTime             `json:"time"`
}

func (a *OrderBook) Update(b *OrderBook) int64 {
	a.Asks = insertAll(a.Asks, b.Asks, 1)
	a.Bids = insertAll(a.Bids, b.Bids, -1)
	return checksum(*a)
}

func insertAll(data [][]decimal.Decimal, elements [][]decimal.Decimal, direction int) [][]decimal.Decimal {
	for _, e := range elements {
		data = insert(data, e, direction)
	}
	if len(data) > 100 {
		data = data[:100]
	}
	return data
}

func insert(data [][]decimal.Decimal, element []decimal.Decimal, direction int) [][]decimal.Decimal {
	i, shift := find(data, element, direction)
	if i == -1 {
		data = append(data, element)
	} else if element[1].IsZero() {
		data = append(data[:i], data[i+1:]...)
	} else {
		if shift {
			data = append(data[:i+1], data[i:]...)
		}
		data[i] = element
	}
	return data
}

func find(data [][]decimal.Decimal, element []decimal.Decimal, direction int) (int, bool) {
	for i, d := range data {
		cmp := d[0].Cmp(element[0])
		if cmp == 0 {
			return i, false
		} else if cmp == direction {
			return i, true
		}
	}
	return -1, false
}

var (
	semicolon = []byte(":")
	dot_zero  = []byte(".0")
)

func checksum(a OrderBook) int64 {
	size := size(len(a.Asks), len(a.Bids))
	if size == 0 {
		return 0
	}
	crc := crc32.NewIEEE()
	writePair(crc, a.Bids[0])
	crc.Write(semicolon)
	writePair(crc, a.Asks[0])
	for i := 1; i < size; i++ {
		crc.Write(semicolon)
		writePair(crc, a.Bids[i])
		crc.Write(semicolon)
		writePair(crc, a.Asks[i])
	}
	return int64(crc.Sum32())
}

func writePair(crc hash.Hash32, p []decimal.Decimal) {
	writeDecimal(crc, p[0])
	crc.Write(semicolon)
	writeDecimal(crc, p[1])
}

func writeDecimal(crc hash.Hash32, d decimal.Decimal) {
	crc.Write([]byte(d.String()))
	if d.Exponent() == 0 {
		crc.Write(dot_zero)
	}
}

func size(a, b int) int {
	if a < b {
		return a
	}
	return b
}
