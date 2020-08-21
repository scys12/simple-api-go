package models

import (
	"github.com/globalsign/mgo/bson"
)

type Books struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Title  string        `bson:"title" json:"title"`
	Genre  string        `bson:"genre" json:"genre"`
	Author string        `bson:"author" json:"author"`
}

const (
	db         = "Books"
	collection = "MovieModel"
)

func (b *Books) InsertBook(book Books) error {
	return Insert(db, collection, book)
}

func (b *Books) FindAllBooks() ([]Books, error) {
	var result []Books
	err := FindAll(db, collection, nil, nil, &result)
	return result, err
}

func (b *Books) FindBookById(id string) (Books, error) {
	var result Books
	err := FindOne(db, collection, bson.M{"_id": bson.ObjectIdHex(id)}, nil, &result)
	return result, err
}

func (b *Books) UpdateBook(book Books) error {
	return Update(db, collection, bson.M{"_id": book.ID}, book)
}

func (b *Books) RemoveBook(id string) error {
	return Remove(db, collection, bson.M{"_id": bson.ObjectIdHex(id)})
}
