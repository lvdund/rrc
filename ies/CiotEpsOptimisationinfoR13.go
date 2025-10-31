package ies

// CIOT-EPS-OptimisationInfo-r13 ::= SEQUENCE OF CIOT-OptimisationPLMN-r13
// SIZE (1.. maxPLMN-r11)
type CiotEpsOptimisationinfoR13 struct {
	Value []CiotOptimisationplmnR13 `lb:1,ub:maxPLMNR11`
}
