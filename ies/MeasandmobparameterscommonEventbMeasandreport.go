package ies

import "rrc/utils"

// MeasAndMobParametersCommon-eventB-MeasAndReport ::= ENUMERATED
type MeasandmobparameterscommonEventbMeasandreport struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonEventbMeasandreportEnumeratedNothing = iota
	MeasandmobparameterscommonEventbMeasandreportEnumeratedSupported
)
