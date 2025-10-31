package ies

// IntraBandCC-CombinationReqList-r17 ::= SEQUENCE
type IntrabandccCombinationreqlistR17 struct {
	ServcellindexlistR17 []Servcellindex             `lb:1,ub:maxNrofServingCells`
	CcCombinationlistR17 []IntrabandccCombinationR17 `lb:1,ub:maxNrofReqComDCLocationR17`
}
