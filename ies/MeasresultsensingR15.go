package ies

import "rrc/utils"

// MeasResultSensing-r15 ::= SEQUENCE
type MeasresultsensingR15 struct {
	SlSubframerefR15 utils.INTEGER      `lb:0,ub:10239`
	SensingresultR15 []SensingresultR15 `lb:0,ub:400`
}
