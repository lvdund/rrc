package ies

import "rrc/utils"

// GeneralParametersMRDC-XDD-Diff-srb3 ::= ENUMERATED
type GeneralparametersmrdcXddDiffSrb3 struct {
	Value utils.ENUMERATED
}

const (
	GeneralparametersmrdcXddDiffSrb3EnumeratedNothing = iota
	GeneralparametersmrdcXddDiffSrb3EnumeratedSupported
)
