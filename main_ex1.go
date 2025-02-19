package main

// Implement a function that takes a list of words and outputs groups of anagrams.
// For instance, given an input structure of:
// { "relayed", "dog", "baker", "secure", "break", "layered", "rescue", "god", "delayer", "recuse", "cat" },
// the function returns:
// { { "dog", "god" }, { "baker", "break" }, { "secure", "rescue", "recuse" }, {"delayer", "layered", "relayed" }, { "cat" } }

// baker-> abekr
// break -> abekr

import (
	"slices"
	"time"
)

func grupAnagrams(arr []string) map[string][]string {
	groups := make(map[string][]string)
	for _, str := range arr {
		byteArr := []byte(str)
		slices.Sort(byteArr)
		key := string(byteArr)
		groups[key] = append(groups[key], str)
	}
	return groups
}

// func main() {
// 	arr := []string{"relayed", "dog", "baker", "secure", "break", "layered", "rescue", "god", "delayer", "recuse", "cat"}
// 	result := grupAnagrams(arr)
// 	fmt.Println(result)
// }

// Design and implement an HTTP API that implements the backend for a public book library.
// The library needs to keep track of existing stock, lent books (and who they are lent to) and return due dates,
// as well as add new stock and lend books out. The API does not use database persistence, storing everything in memory instead.
// The query routing is less important, what matters is the design of the API and how request handlers manage concurrent access.

// Stock:
// [book ID , units] m1
//

// lentBook:
// [book id, return date, user id] m2

// API:
// BookReceiver(lentBook) m2 WL
// LentBook(lentBook) m2 RL
// NewStock(Stock) m1 WL
// GetBooks()[]Stock m1 RL
// GetBookByID(bookID) (units, [](returnDate, user)) m1 RL

// HTTP 1.1 Delete : /receive_book { bookid: string, returnDate: time, user string} ->
// HTTP 1.1 Puts : /lent_book { bookid: string, returnDate: time, user string} ->
// HTTP 1.1 Insert : /new_stock {bookID: string, units int} ->
// HTTP 1.1 Get : /books -> []{bookID: string, units int}
// HTTP 1.1 Get : /books/id {} ->  {units: int}

type Stock struct {
	BookID string
	Units  uint
}

type LentBookMetadata struct {
	ReturnDate time.Time
	User       string
}

type LentBook struct {
	BookID   string
	Metadata LentBookMetadata
}

type BookLibrary interface {
	ReceiveBook(LentBook)
	LentBook(LentBook)
	NewStock(Stock)
	GetBooks() []Stock
	GetBookByID(bookID string) (uint, []LentBookMetadata)
}

type mockBookLibrary struct {
	stocks    map[string]uint // keyed by bookID
	lentBooks map[string][]LentBookMetadata
}

func (m *mockBookLibrary) ReceiveBook(info LentBook) {
	metadata := info.Metadata
	metArr := m.lentBooks[info.BookID]
	if ind := slices.Index(metArr, metadata); ind != -1 {
		metArr[ind] = metArr[0]
		metArr = metArr[1:]
	}
	m.lentBooks[info.BookID] = metArr
}

func (m *mockBookLibrary) LentBook(info LentBook) {
	m.lentBooks[info.BookID] = append(m.lentBooks[info.BookID], info.Metadata)

}
func (m *mockBookLibrary) NewStock(stock Stock) {
	m.stocks[stock.BookID] += stock.Units
}

func (m *mockBookLibrary) GetBooks() []Stock {
	var stocks []Stock
	for bookID, units := range m.stocks {
		stocks = append(stocks, Stock{BookID: bookID, Units: units})
	}
	return stocks
}

func (m *mockBookLibrary) GetBookByID(bookID string) (uint, []LentBookMetadata) {
	return m.stocks[bookID], m.lentBooks[bookID]
}

func newMockBookLibrary() BookLibrary {
	return &mockBookLibrary{}
}

// type realBookLibrary struct{

// }

func newBookLibrary() BookLibrary {
	return &mockBookLibrary{} // real
}

func main() {
	// b := newBookLibrary()
	// b := newMockBookLibrary()
	// .. b
	//  httpserver := { b = newBookLibrary() }
}
