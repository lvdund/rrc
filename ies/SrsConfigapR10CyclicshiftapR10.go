package ies

import "rrc/utils"

// SRS-ConfigAp-r10-cyclicShiftAp-r10 ::= ENUMERATED
type SrsConfigapR10CyclicshiftapR10 struct {
	Value utils.ENUMERATED
}

const (
	SrsConfigapR10CyclicshiftapR10EnumeratedNothing = iota
	SrsConfigapR10CyclicshiftapR10EnumeratedCs0
	SrsConfigapR10CyclicshiftapR10EnumeratedCs1
	SrsConfigapR10CyclicshiftapR10EnumeratedCs2
	SrsConfigapR10CyclicshiftapR10EnumeratedCs3
	SrsConfigapR10CyclicshiftapR10EnumeratedCs4
	SrsConfigapR10CyclicshiftapR10EnumeratedCs5
	SrsConfigapR10CyclicshiftapR10EnumeratedCs6
	SrsConfigapR10CyclicshiftapR10EnumeratedCs7
)
