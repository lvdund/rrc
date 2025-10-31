package ies

import "rrc/utils"

// AntennaInfoUL-r10-fourAntennaPortActivated-r10 ::= ENUMERATED
type AntennainfoulR10FourantennaportactivatedR10 struct {
	Value utils.ENUMERATED
}

const (
	AntennainfoulR10FourantennaportactivatedR10EnumeratedNothing = iota
	AntennainfoulR10FourantennaportactivatedR10EnumeratedSetup
)
