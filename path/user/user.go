package path

import "github.com/AdiKhoironHasan/bookservices-protobank/path"

func GetUserProtectedMethods() path.ProtectedMethods {
	return path.ProtectedMethods{
		"/user.UserService/Ping":   true,
		"/user.UserService/List":   true,
		"/user.UserService/Store":  true,
		"/user.UserService/Detail": true,
		"/user.UserService/Update": true,
		"/user.UserService/Delete": true,
	}
}
