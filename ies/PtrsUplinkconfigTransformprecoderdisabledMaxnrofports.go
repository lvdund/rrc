package ies

import "rrc/utils"

// PTRS-UplinkConfig-transformPrecoderDisabled-maxNrofPorts ::= ENUMERATED
type PtrsUplinkconfigTransformprecoderdisabledMaxnrofports struct {
	Value utils.ENUMERATED
}

const (
	PtrsUplinkconfigTransformprecoderdisabledMaxnrofportsEnumeratedNothing = iota
	PtrsUplinkconfigTransformprecoderdisabledMaxnrofportsEnumeratedN1
	PtrsUplinkconfigTransformprecoderdisabledMaxnrofportsEnumeratedN2
)
