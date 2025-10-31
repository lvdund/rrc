package ies

import "rrc/utils"

// SRS-AntennaPort ::= ENUMERATED
type SrsAntennaport struct {
	Value utils.ENUMERATED
}

const (
	SrsAntennaportEnumeratedNothing = iota
	SrsAntennaportEnumeratedAn1
	SrsAntennaportEnumeratedAn2
	SrsAntennaportEnumeratedAn4
	SrsAntennaportEnumeratedSpare1
)
