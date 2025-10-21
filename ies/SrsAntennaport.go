package ies

import "rrc/utils"

// SRS-AntennaPort ::= ENUMERATED
type SrsAntennaport struct {
	Value utils.ENUMERATED
}

const (
	SrsAntennaportAn1    = 0
	SrsAntennaportAn2    = 1
	SrsAntennaportAn4    = 2
	SrsAntennaportSpare1 = 3
)
