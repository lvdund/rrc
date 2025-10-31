package ies

import "rrc/utils"

// MeasAndMobParametersCommon-periodicEUTRA-MeasAndReport ::= ENUMERATED
type MeasandmobparameterscommonPeriodiceutraMeasandreport struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonPeriodiceutraMeasandreportEnumeratedNothing = iota
	MeasandmobparameterscommonPeriodiceutraMeasandreportEnumeratedSupported
)
