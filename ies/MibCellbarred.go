package ies

import "rrc/utils"

// MIB-cellBarred ::= ENUMERATED
type MibCellbarred struct {
	Value utils.ENUMERATED
}

const (
	MibCellbarredEnumeratedNothing = iota
	MibCellbarredEnumeratedBarred
	MibCellbarredEnumeratedNotbarred
)
