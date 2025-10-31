package ies

import "rrc/utils"

// MeasAndMobParametersCommon-condHandoverWithSCG-NRDC-r17 ::= ENUMERATED
type MeasandmobparameterscommonCondhandoverwithscgNrdcR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonCondhandoverwithscgNrdcR17EnumeratedNothing = iota
	MeasandmobparameterscommonCondhandoverwithscgNrdcR17EnumeratedSupported
)
