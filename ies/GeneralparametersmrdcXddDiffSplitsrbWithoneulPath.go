package ies

import "rrc/utils"

// GeneralParametersMRDC-XDD-Diff-splitSRB-WithOneUL-Path ::= ENUMERATED
type GeneralparametersmrdcXddDiffSplitsrbWithoneulPath struct {
	Value utils.ENUMERATED
}

const (
	GeneralparametersmrdcXddDiffSplitsrbWithoneulPathEnumeratedNothing = iota
	GeneralparametersmrdcXddDiffSplitsrbWithoneulPathEnumeratedSupported
)
