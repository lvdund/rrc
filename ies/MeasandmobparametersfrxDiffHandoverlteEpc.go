package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-handoverLTE-EPC ::= ENUMERATED
type MeasandmobparametersfrxDiffHandoverlteEpc struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffHandoverlteEpcEnumeratedNothing = iota
	MeasandmobparametersfrxDiffHandoverlteEpcEnumeratedSupported
)
