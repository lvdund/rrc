package ies

// PDCCH-ConfigLAA-r14 ::= SEQUENCE
type PdcchConfiglaaR14 struct {
	MaxnumberofschedsubframesFormat0bR14 *PdcchConfiglaaR14MaxnumberofschedsubframesFormat0bR14
	MaxnumberofschedsubframesFormat4bR14 *PdcchConfiglaaR14MaxnumberofschedsubframesFormat4bR14
	SkipmonitoringdciFormat0aR14         *bool
	SkipmonitoringdciFormat4aR14         *bool
	PdcchCandidatereductionsFormat0aR14  *PdcchCandidatereductionsR13
	PdcchCandidatereductionsFormat4aR14  *PdcchCandidatereductionslaaUlR14
	PdcchCandidatereductionsFormat0bR14  *PdcchCandidatereductionslaaUlR14
	PdcchCandidatereductionsFormat4bR14  *PdcchCandidatereductionslaaUlR14
}
