package ies

import "rrc/utils"

// MeasAndMobParametersXDD-Diff-sftd-MeasNR-Neigh-DRX ::= ENUMERATED
type MeasandmobparametersxddDiffSftdMeasnrNeighDrx struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersxddDiffSftdMeasnrNeighDrxEnumeratedNothing = iota
	MeasandmobparametersxddDiffSftdMeasnrNeighDrxEnumeratedSupported
)
