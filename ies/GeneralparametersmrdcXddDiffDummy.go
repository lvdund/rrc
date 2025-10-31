package ies

import "rrc/utils"

// GeneralParametersMRDC-XDD-Diff-dummy ::= ENUMERATED
type GeneralparametersmrdcXddDiffDummy struct {
	Value utils.ENUMERATED
}

const (
	GeneralparametersmrdcXddDiffDummyEnumeratedNothing = iota
	GeneralparametersmrdcXddDiffDummyEnumeratedSupported
)
