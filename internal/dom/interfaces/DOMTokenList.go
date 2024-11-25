package main

type DOMTokenList struct {
	tokens []string
}

type DOMTokenListInterface interface {
	GetLength() uint64
	GetItem(index uint64) string
	Contains(token string) bool
	Add(tokens ...string)
	Remove(tokens ...string)
	Toggle(token string, force bool) bool
	Replace(token string, newToken string) bool
	Supports(token string) bool
	stringify() string
	Iterate() <-chan string
}
