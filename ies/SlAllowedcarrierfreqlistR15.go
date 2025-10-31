package ies

// SL-AllowedCarrierFreqList-r15 ::= SEQUENCE
type SlAllowedcarrierfreqlistR15 struct {
	Allowedcarrierfreqset1 []ArfcnValueeutraR9 `lb:1,ub:maxFreqV2XR14`
	Allowedcarrierfreqset2 []ArfcnValueeutraR9 `lb:1,ub:maxFreqV2XR14`
}
