package ies

// PUCCH-ResourceGroup-r16 ::= SEQUENCE
type PucchResourcegroupR16 struct {
	PucchResourcegroupidR16 PucchResourcegroupidR16
	ResourcepergrouplistR16 []PucchResourceid `lb:1,ub:maxNrofPUCCHResourcespergroupR16`
}
