package ies

import "rrc/utils"

// MeasCycleSCell-r10 ::= ENUMERATED
type MeascyclescellR10 struct {
	Value utils.ENUMERATED
}

const (
	MeascyclescellR10EnumeratedNothing = iota
	MeascyclescellR10EnumeratedSf160
	MeascyclescellR10EnumeratedSf256
	MeascyclescellR10EnumeratedSf320
	MeascyclescellR10EnumeratedSf512
	MeascyclescellR10EnumeratedSf640
	MeascyclescellR10EnumeratedSf1024
	MeascyclescellR10EnumeratedSf1280
	MeascyclescellR10EnumeratedSpare1
)
