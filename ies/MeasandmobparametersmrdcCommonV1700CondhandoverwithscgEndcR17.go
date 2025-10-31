package ies

import "rrc/utils"

// MeasAndMobParametersMRDC-Common-v1700-condHandoverWithSCG-ENDC-r17 ::= ENUMERATED
type MeasandmobparametersmrdcCommonV1700CondhandoverwithscgEndcR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersmrdcCommonV1700CondhandoverwithscgEndcR17EnumeratedNothing = iota
	MeasandmobparametersmrdcCommonV1700CondhandoverwithscgEndcR17EnumeratedSupported
)
