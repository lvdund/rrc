package ies

import "rrc/utils"

// MeasAndMobParametersCommon-eventD1-MeasReportTrigger-r17 ::= ENUMERATED
type MeasandmobparameterscommonEventd1MeasreporttriggerR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonEventd1MeasreporttriggerR17EnumeratedNothing = iota
	MeasandmobparameterscommonEventd1MeasreporttriggerR17EnumeratedSupported
)
