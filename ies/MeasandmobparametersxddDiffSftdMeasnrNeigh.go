package ies

import "rrc/utils"

// MeasAndMobParametersXDD-Diff-sftd-MeasNR-Neigh ::= ENUMERATED
type MeasandmobparametersxddDiffSftdMeasnrNeigh struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersxddDiffSftdMeasnrNeighEnumeratedNothing = iota
	MeasandmobparametersxddDiffSftdMeasnrNeighEnumeratedSupported
)
