package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-handoverLTE-5GC ::= ENUMERATED
type MeasandmobparametersfrxDiffHandoverlte5gc struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffHandoverlte5gcEnumeratedNothing = iota
	MeasandmobparametersfrxDiffHandoverlte5gcEnumeratedSupported
)
