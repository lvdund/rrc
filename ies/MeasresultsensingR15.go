package ies

import "rrc/utils"

// MeasResultSensing-r15 ::= SEQUENCE
type MeasresultsensingR15 struct {
	SlSubframerefR15 utils.INTEGER
	SensingresultR15 interface{}
}
