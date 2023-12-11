package model

import "github.com/google/uuid"

type Artifact struct {
	ID          uuid.UUID
	Membership  float64
	Name        string
	Century     string
	Decade      string
	Age         string
	Description string
}

func NewArtefact(century, decade, age string) *Artifact {
	return &Artifact{
		Century: century,
		Decade:  decade,
		Age:     age,
	}
}
