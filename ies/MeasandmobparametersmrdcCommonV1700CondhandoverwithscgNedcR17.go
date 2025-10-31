package ies

import "rrc/utils"

// MeasAndMobParametersMRDC-Common-v1700-condHandoverWithSCG-NEDC-r17 ::= ENUMERATED
type MeasandmobparametersmrdcCommonV1700CondhandoverwithscgNedcR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersmrdcCommonV1700CondhandoverwithscgNedcR17EnumeratedNothing = iota
	MeasandmobparametersmrdcCommonV1700CondhandoverwithscgNedcR17EnumeratedSupported
)
