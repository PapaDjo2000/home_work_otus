package book

import (
	"github.com/google/uuid"
)

type Usertype int

type Compare struct {
	Usertype Usertype
}

const (
	identical          = "одинаковые"
	Year      Usertype = iota
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

func (b Book) GetID() uuid.UUID {
	return b.id
}

func (b *Book) SetID(id uuid.UUID) {
	b.id = id
}

func (b Book) GetTitle() string {
	return b.title
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b Book) GetAuthor() string {
	return b.author
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b Book) GetYear() int {
	return b.year
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b Book) GetSize() int {
	return b.size
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b Book) GetRate() float32 {
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

func (c Compare) CompareBook(book1, book2 Book) (t1 any) {
	switch c.Usertype {
	case Year:
		switch {
		case book1.year > book2.year:
			return true
		case book1.year < book2.year:
			return false
		default:
			return identical
		}
	case Size:
		switch {
		case book1.size > book2.size:
			return true
		case book1.size < book2.size:
			return false
		default:
			return identical
		}
	case Rate:
		switch {
		case book1.rate > book2.rate:
			return true
		case book1.rate < book2.rate:
			return false
		default:
			return identical
		}
	}
	return "ошибка"
}
