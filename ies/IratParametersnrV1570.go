package ies

import "rrc/utils"

// IRAT-ParametersNR-v1570 ::= SEQUENCE
type IratParametersnrV1570 struct {
	SsSinrMeasNrFr1R15 *utils.ENUMERATED
	SsSinrMeasNrFr2R15 *utils.ENUMERATED
}
