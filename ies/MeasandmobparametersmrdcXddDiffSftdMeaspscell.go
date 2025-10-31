package ies

import "rrc/utils"

// MeasAndMobParametersMRDC-XDD-Diff-sftd-MeasPSCell ::= ENUMERATED
type MeasandmobparametersmrdcXddDiffSftdMeaspscell struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersmrdcXddDiffSftdMeaspscellEnumeratedNothing = iota
	MeasandmobparametersmrdcXddDiffSftdMeaspscellEnumeratedSupported
)
