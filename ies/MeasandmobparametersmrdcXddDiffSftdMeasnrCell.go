package ies

import "rrc/utils"

// MeasAndMobParametersMRDC-XDD-Diff-sftd-MeasNR-Cell ::= ENUMERATED
type MeasandmobparametersmrdcXddDiffSftdMeasnrCell struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersmrdcXddDiffSftdMeasnrCellEnumeratedNothing = iota
	MeasandmobparametersmrdcXddDiffSftdMeasnrCellEnumeratedSupported
)
