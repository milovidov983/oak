package models

type IResource interface {
}

type Secret string
type Login string

type Resource struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Secret Secret `json:"secret"`
	Login  Login  `json:"login"`
}

func (r Resource) GetID() string {
	return r.ID
}
