package ies

import "rrc/utils"

// MeasAndMobParametersCommon-independentGapConfigPRS-r17 ::= ENUMERATED
type MeasandmobparameterscommonIndependentgapconfigprsR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonIndependentgapconfigprsR17EnumeratedNothing = iota
	MeasandmobparameterscommonIndependentgapconfigprsR17EnumeratedSupported
)
