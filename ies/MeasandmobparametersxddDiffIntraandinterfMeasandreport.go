package ies

import "rrc/utils"

// MeasAndMobParametersXDD-Diff-intraAndInterF-MeasAndReport ::= ENUMERATED
type MeasandmobparametersxddDiffIntraandinterfMeasandreport struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersxddDiffIntraandinterfMeasandreportEnumeratedNothing = iota
	MeasandmobparametersxddDiffIntraandinterfMeasandreportEnumeratedSupported
)
