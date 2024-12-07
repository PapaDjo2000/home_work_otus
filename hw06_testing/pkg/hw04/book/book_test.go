package book

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	title          = "Властелин колец"
	author         = "Толкин"
	year           = 1954
	size           = 254
	rate   float32 = 4.5
)

func TestBook_NewBook(t *testing.T) {
	book := NewBook(title, author, year, size, rate)

	assert.NotNil(t, book.id)
	assert.Equal(t, author, book.author)
	assert.Equal(t, title, book.title)
	assert.Equal(t, year, book.year)
	assert.Equal(t, size, book.size)
	assert.Equal(t, rate, book.rate)
}

func TestBook_SetID(t *testing.T) {
	book := NewBook(title, author, year, size, rate)
	testSetID := uuid.New()

	book.SetID(testSetID)
	assert.Equal(t, testSetID, book.id)
}

func TestBook_SetTitle(t *testing.T) {
	book := NewBook(title, author, year, size, rate)
	testSetTitle := "Властелин колец 2"
	book.SetTitle(testSetTitle)
	assert.Equal(t, testSetTitle, book.title)
}

func TestBook_SetAuthor(t *testing.T) {
	book := NewBook(title, author, year, size, rate)
	testSetAuthor := "Толстой"
	book.SetAuthor(testSetAuthor)
	assert.Equal(t, testSetAuthor, book.author)
}
func TestBook_SetYear(t *testing.T) {
	book := NewBook(title, author, year, size, rate)
	testSetYear := 1923
	book.SetYear(testSetYear)
	assert.Equal(t, testSetYear, book.year)
}

func TestBook_SetSize(t *testing.T) {
	book := NewBook(title, author, year, size, rate)
	testSetSize := 1923
	book.SetSize(testSetSize)
	assert.Equal(t, testSetSize, book.size)
}
func TestBook_SetRate(t *testing.T) {
	book := NewBook(title, author, year, size, rate)
	var testSetRate float32 = 324.42
	book.SetRate(testSetRate)
	assert.Equal(t, testSetRate, book.rate)
}

func TestBook_Title(t *testing.T) {
	book := NewBook(title, author, year, size, rate)
	testTitle := book.Title()
	assert.Equal(t, testTitle, book.title)
}

func TestBook_Author(t *testing.T) {
	book := NewBook(title, author, year, size, rate)
	testAuthor := book.Author()
	assert.Equal(t, testAuthor, book.author)
}

func TestBook_Year(t *testing.T) {
	book := NewBook(title, author, year, size, rate)
	testYear := book.Year()
	assert.Equal(t, testYear, book.year)
}

func TestBook_Size(t *testing.T) {
	book := NewBook(title, author, year, size, rate)
	testSize := book.Size()
	assert.Equal(t, testSize, book.size)
}

func TestBook_Rate(t *testing.T) {
	book := NewBook(title, author, year, size, rate)
	testRate := book.Rate()
	assert.Equal(t, testRate, book.rate)
}

func TestBook_NewCompare(t *testing.T) {
	testNewCompare := NewCompare(Year)
	assert.Equal(t, testNewCompare.Usertype, Year)
}

func TestBook_CompareBook(t *testing.T) {
	var def Usertype = 5
	newbook1 := NewBook("Властелин колец 1", "Толкин", 1954, 1125, 154.32)
	newbook2 := NewBook("Властелин колец 2", "Толкин", 1955, 954, 154.32)

	Compare := NewCompare(Year)
	testCompare := Compare.CompareBook(*newbook1, *newbook2)
	assert.True(t, true, testCompare)

	Compare = NewCompare(Size)
	testCompare = Compare.CompareBook(*newbook1, *newbook2)
	assert.True(t, true, testCompare)

	Compare = NewCompare(Rate)
	testCompare = Compare.CompareBook(*newbook1, *newbook2)
	assert.True(t, true, testCompare)

	Compare = NewCompare(def)
	testCompare = Compare.CompareBook(*newbook1, *newbook2)
	assert.False(t, false, testCompare)

	///ДВа варианта !

	testCases := []struct {
		name     string
		usertype Usertype
	}{
		{name: "Year",
			usertype: Year,
		},
		{name: "Size",
			usertype: Size,
		},
		{name: "Rate",
			usertype: Rate,
		},
		{name: "def",
			usertype: def,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			if tC.usertype != def {
				Compare = NewCompare(tC.usertype)
				testCompare = Compare.CompareBook(*newbook1, *newbook2)
				assert.True(t, true, testCompare)
			} else {
				Compare = NewCompare(def)
				testCompare = Compare.CompareBook(*newbook1, *newbook2)
				assert.False(t, false, testCompare)
			}
		})
	}
}
