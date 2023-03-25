package path

import "github.com/AdiKhoironHasan/bookservices-protobank/path"

func GetBookProtectedMethods() path.ProtectedMethods {
	return path.ProtectedMethods{
		"/book.BookService/Ping":   true,
		"/book.BookService/List":   true,
		"/book.BookService/Store":  true,
		"/book.BookService/Detail": true,
		"/book.BookService/Update": true,
		"/book.BookService/Delete": true,
	}
}
