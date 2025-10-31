package ies

import "rrc/utils"

// MeasAndMobParametersXDD-Diff-handoverLTE-5GC ::= ENUMERATED
type MeasandmobparametersxddDiffHandoverlte5gc struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersxddDiffHandoverlte5gcEnumeratedNothing = iota
	MeasandmobparametersxddDiffHandoverlte5gcEnumeratedSupported
)
