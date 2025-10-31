package ies

// MeasAndMobParameters ::= SEQUENCE
type Measandmobparameters struct {
	Measandmobparameterscommon  *Measandmobparameterscommon
	MeasandmobparametersxddDiff *MeasandmobparametersxddDiff
	MeasandmobparametersfrxDiff *MeasandmobparametersfrxDiff
}
