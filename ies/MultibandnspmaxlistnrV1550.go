package ies

// MultiBandNsPmaxListNR-v1550 ::= SEQUENCE OF NS-PmaxListNR-r15
// SIZE (1.. maxMultiBandsNR-r15)
type MultibandnspmaxlistnrV1550 struct {
	Value []NsPmaxlistnrR15 `lb:1,ub:maxMultiBandsNRR15`
}
