package ies

import "rrc/utils"

// PDCCH-ConfigLAA-r14 ::= SEQUENCE
type PdcchConfiglaaR14 struct {
	MaxnumberofschedsubframesFormat0bR14 *utils.ENUMERATED
	MaxnumberofschedsubframesFormat4bR14 *utils.ENUMERATED
	SkipmonitoringdciFormat0aR14         *utils.ENUMERATED
	SkipmonitoringdciFormat4aR14         *utils.ENUMERATED
	PdcchCandidatereductionsFormat0aR14  *PdcchCandidatereductionsR13
	PdcchCandidatereductionsFormat4aR14  *PdcchCandidatereductionslaaUlR14
	PdcchCandidatereductionsFormat0bR14  *PdcchCandidatereductionslaaUlR14
	PdcchCandidatereductionsFormat4bR14  *PdcchCandidatereductionslaaUlR14
}
