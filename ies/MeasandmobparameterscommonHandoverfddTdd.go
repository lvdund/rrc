package ies

import "rrc/utils"

// MeasAndMobParametersCommon-handoverFDD-TDD ::= ENUMERATED
type MeasandmobparameterscommonHandoverfddTdd struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonHandoverfddTddEnumeratedNothing = iota
	MeasandmobparameterscommonHandoverfddTddEnumeratedSupported
)
