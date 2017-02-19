package main

import (
    "fmt"
)

/*      Defining a Structure
    type struct_variable_type struct {
        member definition;
        member definition;
        ...
        member definition;
    }
*/
type Books struct {
    title string
    author string
    subject string
    book_id int
}

func main() {
    var Book1 Books
    var Book2 Books

    // Book1 specification
    Book1.title = "Go Programming"
    Book1.author = "Mahesh Kumar"
    Book1.subject = "Go Programming Tutorial"
    Book1.book_id = 6495407

    // Book2 specification
    Book2.title = "Java Programming"
    Book2.author = "James Gosling"
    Book2.subject = "Java Programming Tutorial"
    Book2.book_id = 6495701

    /* print Book1 info */
    fmt.Printf( "Book 1 title : %s\n", Book1.title)
    fmt.Printf( "Book 1 author : %s\n", Book1.author)
    fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
    fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

    fmt.Println()

    printBook(Book2)

    //      Pointers to Structures
    var struct_pointer *Books

    struct_pointer = &Book1

    fmt.Println("\nPrint using pointers to structures\n")

    printBookByPointer(struct_pointer)
}

func printBook(book Books) {
    /* print Book2 info */
    fmt.Printf( "Book title : %s\n", book.title)
    fmt.Printf( "Book author : %s\n", book.author)
    fmt.Printf( "Book subject : %s\n", book.subject)
    fmt.Printf( "Book book_id : %d\n", book.book_id)
}

func printBookByPointer(book *Books) {
    fmt.Printf( "Book title : %s\n", book.title);
    fmt.Printf( "Book author : %s\n", book.author);
    fmt.Printf( "Book subject : %s\n", book.subject);
    fmt.Printf( "Book book_id : %d\n", book.book_id);
}