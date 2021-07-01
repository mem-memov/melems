package melems

import (
	"github.com/mem-memov/mbook"
)

type Storage struct {
	entryLength uint
	book        *mbook.Book
}

func NewStorage(entryLength uint, increment uint) (*Storage, error) {
	book, err := mbook.NewBook(increment * entryLength)
	if err != nil {
		return nil, err
	}

	s := &Storage{
		entryLength: entryLength,
		book:        book,
	}

	return s, nil
}

func (s *Storage) Create() *Entry {
	elements := make([]uint, s.entryLength, s.entryLength)
	return newEntry(elements)
}

func (s *Storage) Read(position uint, entry *Entry) error {
	for i := uint(0); i < s.entryLength; i++ {
		value := s.book.Read(position + i)
		err := entry.Set(i, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) Write(position uint, entry *Entry) error {
	for i := uint(0); i < s.entryLength; i++ {
		value, err := entry.Get(i)
		if err != nil {
			return err
		}
		s.book.Write(position+i, value)
	}

	return nil
}
