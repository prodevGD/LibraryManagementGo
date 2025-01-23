package main

import (
        "bufio"
        "fmt"
        "os"
        "strings"
)

// Book struct
type Book struct {
        Title   string
        Author  string
        ISBN    string
        Available bool
}

// Initialize a new book
func (b *Book) Initialize(title, author, isbn string) {
        b.Title = title
        b.Author = author
        b.ISBN = isbn
        b.Available = true
}

// Display book details
func (b *Book) DisplayDetails() {
        fmt.Printf("Title: %s\n", b.Title)
        fmt.Printf("Author: %s\n", b.Author)
        fmt.Printf("ISBN: %s\n", b.ISBN)
        fmt.Printf("Available: %t\n", b.Available)
}

// EBook struct (embeds Book)
type EBook struct {
        Book
        FileSize int
}

// Override DisplayDetails for EBook
func (eb *EBook) DisplayDetails() {
        eb.Book.DisplayDetails()
        fmt.Printf("File Size: %d MB\n", eb.FileSize)
}

// BookInterface
type BookInterface interface {
        DisplayDetails()
}

// Library struct
type Library struct {
        books []BookInterface
}

// Add a book to the library
func (l *Library) AddBook(book BookInterface) error {
        // Check for duplicate ISBN
        for _, existingBook := range l.books {
                switch b := existingBook.(type) {
                case *Book:
                        if b.ISBN == book.(*Book).ISBN {
                                return fmt.Errorf("ISBN already exists")
                        }
                case *EBook:
                        if b.ISBN == book.(*Book).ISBN {
                                return fmt.Errorf("ISBN already exists")
                        }
                }
        }
        l.books = append(l.books, book)
        return nil
}

// Remove a book from the library
func (l *Library) RemoveBook(isbn string) error {
        for i, book := range l.books {
                switch b := book.(type) {
                case *Book:
                        if b.ISBN == isbn {
                                l.books = append(l.books[:i], l.books[i+1:]...)
                                return nil
                        }
                case *EBook:
                        if b.ISBN == isbn {
                                l.books = append(l.books[:i], l.books[i+1:]...)
                                return nil
                        }
                }
        }
        return fmt.Errorf("Book with ISBN %s not found", isbn)
}

// Search for books by title
func (l *Library) SearchBookByTitle(title string) []BookInterface {
        var results []BookInterface
        for _, book := range l.books {
                switch b := book.(type) {
                case *Book:
                        if strings.Contains(b.Title, title) {
                                results = append(results, b)
                        }
                case *EBook:
                        if strings.Contains(b.Title, title) {
                                results = append(results, b)
                        }
                }
        }
        return results
}

// List all books in the library
func (l *Library) ListBooks() {
        for _, book := range l.books {
                book.DisplayDetails()
                fmt.Println("--------------------")
        }
}

func main() {
        var library Library
        reader := bufio.NewReader(os.Stdin)

        for {
                fmt.Println("\nLibrary Management System")
                fmt.Println("1. Add Book/EBook")
                fmt.Println("2. Remove Book/EBook")
                fmt.Println("3. Search Books")
                fmt.Println("4. List All Books/EBooks")
                fmt.Println("5. Exit")
                fmt.Print("Enter your choice: ")

                choice, _ := reader.ReadString('\n')
                choice = strings.TrimSpace(choice)

                switch choice {
                case "1":
                        fmt.Println("Enter Book/EBook type (B/E): ")
                        bookType, _ := reader.ReadString('\n')
                        bookType = strings.TrimSpace(bookType)

                        fmt.Print("Enter Title: ")
                        title, _ := reader.ReadString('\n')
                        title = strings.TrimSpace(title)

                        fmt.Print("Enter Author: ")
                        author, _ := reader.ReadString('\n')
                        author = strings.TrimSpace(author)

                        fmt.Print("Enter ISBN: ")
                        isbn, _ := reader.ReadString('\n')
                        isbn = strings.TrimSpace(isbn)

                        if bookType == "B" {
                                book := &Book{}
                                book.Initialize(title, author, isbn)
                                err := library.AddBook(book)
                                if err != nil {
                                        fmt.Println("Error:", err)
                                } else {
                                        fmt.Println("Book added successfully!")
                                }
                        } else if bookType == "E" {
                                fmt.Print("Enter File Size (MB): ")
                                fileSizeStr, _ := reader.ReadString('\n')
                                fileSizeStr = strings.TrimSpace(fileSizeStr)
                                var fileSize int 
                                _, err := fmt.Sscanf(fileSizeStr, "%d", &fileSize) 
                                if err != nil {
                                        fmt.Println("Invalid file size input.")
                                        continue 
                                }

                                eBook := &EBook{Book: Book{Title: title, Author: author, ISBN: isbn}}
                                eBook.FileSize = fileSize
                                err = library.AddBook(eBook)
                                if err != nil {
                                        fmt.Println("Error:", err)
                                } else {
                                        fmt.Println("EBook added successfully!")
                                }
                        } else {
                                fmt.Println("Invalid book type.")
                        }

                case "2":
                        fmt.Print("Enter ISBN of the book to remove: ")
                        isbn, _ := reader.ReadString('\n')
                        isbn = strings.TrimSpace(isbn)
                        err := library.RemoveBook(isbn)
                        if err != nil {
                                fmt.Println("Error:", err)
                        } else {
                                fmt.Println("Book removed successfully!")
                        }

                case "3":
                        fmt.Print("Enter title to search: ")
                        title, _ := reader.ReadString('\n')
                        title = strings.TrimSpace(title)
                        results := library.SearchBookByTitle(title)
                        if len(results) == 0 {
                                fmt.Println("No books found.")
                        } else {
                                fmt.Println("Search Results:")
                                for _, book := range results {
                                        book.DisplayDetails()
                                        fmt.Println("--------------------")
                                }
                        }

                case "4":
                        fmt.Println("List of all Books/EBooks:")
                        library.ListBooks()

                case "5":
                        fmt.Println("Exiting...")
                        os.Exit(0)

                default:
                        fmt.Println("Invalid choice.")
                }
        }
}