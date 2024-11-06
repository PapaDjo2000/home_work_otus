package book

import (
	"github.com/google/uuid"
)

type Usertype int

type Compare struct {
	Usertype Usertype
}

const (
	Year Usertype = iota
	Size
	Rate
)

type Book struct {
	id     uuid.UUID
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func NewBook(title, author string, year, size int, rate float32) *Book {
	id, _ := uuid.NewRandom()
	return &Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

func (b *Book) ID() uuid.UUID {
	return b.id
}

func (b *Book) SetID(id uuid.UUID) {
	b.id = id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) Rate() float32 {
	return b.rate
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

func NewCompare(usertype Usertype) *Compare {
	return &Compare{
		Usertype: usertype,
	}
}

func (c Compare) CompareBook(book1, book2 Book) bool {
	switch c.Usertype {
	case Year:
		return book1.year > book2.year
	case Size:
		return book1.size > book2.size
	case Rate:
		return book1.rate > book2.rate
	default:
		return false
	}
}
