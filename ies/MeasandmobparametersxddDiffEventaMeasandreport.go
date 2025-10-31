package ies

import "rrc/utils"

// MeasAndMobParametersXDD-Diff-eventA-MeasAndReport ::= ENUMERATED
type MeasandmobparametersxddDiffEventaMeasandreport struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersxddDiffEventaMeasandreportEnumeratedNothing = iota
	MeasandmobparametersxddDiffEventaMeasandreportEnumeratedSupported
)
