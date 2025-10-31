package ies

// SchedulingRequestConfig-v1700 ::= SEQUENCE
type SchedulingrequestconfigV1700 struct {
	SchedulingrequesttoaddmodlistextV1700 *[]SchedulingrequesttoaddmodextV1700 `lb:1,ub:maxNrofSRConfigpercellgroup`
}
