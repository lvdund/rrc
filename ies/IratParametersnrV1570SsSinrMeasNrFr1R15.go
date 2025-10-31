package ies

import "rrc/utils"

// IRAT-ParametersNR-v1570-ss-SINR-Meas-NR-FR1-r15 ::= ENUMERATED
type IratParametersnrV1570SsSinrMeasNrFr1R15 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrV1570SsSinrMeasNrFr1R15EnumeratedNothing = iota
	IratParametersnrV1570SsSinrMeasNrFr1R15EnumeratedSupported
)
