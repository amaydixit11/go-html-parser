package main

type DocumentType struct {
	name     string
	publicId string
	systemId string
}
type DocumentTypeInterface interface {
	GetName() string
	GetPublicId() string
	GetSystemId() string
}
