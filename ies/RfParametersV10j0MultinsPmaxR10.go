package ies

import "rrc/utils"

// RF-Parameters-v10j0-multiNS-Pmax-r10 ::= ENUMERATED
type RfParametersV10j0MultinsPmaxR10 struct {
	Value utils.ENUMERATED
}

const (
	RfParametersV10j0MultinsPmaxR10EnumeratedNothing = iota
	RfParametersV10j0MultinsPmaxR10EnumeratedSupported
)
