package ies

import "rrc/utils"

// MeasAndMobParametersXDD-Diff-handoverLTE-EPC ::= ENUMERATED
type MeasandmobparametersxddDiffHandoverlteEpc struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersxddDiffHandoverlteEpcEnumeratedNothing = iota
	MeasandmobparametersxddDiffHandoverlteEpcEnumeratedSupported
)
