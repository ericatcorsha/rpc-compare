// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

type EchoReply struct {
	Message string `json:"message"`
}

type EchoRequest struct {
	Message string `json:"message"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
