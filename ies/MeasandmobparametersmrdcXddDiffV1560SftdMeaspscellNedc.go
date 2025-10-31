package ies

import "rrc/utils"

// MeasAndMobParametersMRDC-XDD-Diff-v1560-sftd-MeasPSCell-NEDC ::= ENUMERATED
type MeasandmobparametersmrdcXddDiffV1560SftdMeaspscellNedc struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersmrdcXddDiffV1560SftdMeaspscellNedcEnumeratedNothing = iota
	MeasandmobparametersmrdcXddDiffV1560SftdMeaspscellNedcEnumeratedSupported
)
