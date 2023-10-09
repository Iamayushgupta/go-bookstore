package models

import (
	"log"
	"github.com/ayush/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Book represents the Book model with gorm.Model embedded for fields like ID, CreatedAt, UpdatedAt, and DeletedAt.
type Book struct {
	gorm.Model
	Name       string `json:"name"`
	Author     string `json:"author"`
	Publication string `json:"publication"`
}

// init initializes the database connection and performs auto migration to create necessary tables.
func init() {
	config.Connect()
	db = config.GetDB()
	// AutoMigrate creates the 'books' table based on the Book model.
	db.AutoMigrate(&Book{})
	log.Println("Database migration completed.")
}

// CreateBook creates a new book record in the database and returns the created book.
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	log.Printf("Book created with ID: %d", b.ID)
	return b
}

// GetAllBooks retrieves all books from the database and returns them as a slice of Book structs.
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	log.Println("Retrieved all books from the database.")
	return Books
}

// GetBookById retrieves a book by its ID from the database and returns the book and the database object.
func GetBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	result := db.Where("ID=?", ID).Find(&getBook)
	if result.RecordNotFound() {
		log.Printf("Book with ID: %d not found.", ID)
		return nil, result
	}
	log.Printf("Retrieved book with ID: %d", ID)
	return &getBook, result
}

// DeleteBook deletes a book by its ID from the database and returns the deleted book.
func DeleteBook(ID int64) *Book {
	book, _ := GetBookById(ID)
	if book == nil {
		log.Printf("Book with ID: %d not found.", ID)
		return nil
	}

	// Delete the fetched book from the database
	result := db.Delete(&book)
	if result.RowsAffected == 0 {
		log.Printf("Book with ID: %d not found.", ID)
		return nil
	}

	log.Printf("Deleted book with ID: %d", ID)
	return book
}