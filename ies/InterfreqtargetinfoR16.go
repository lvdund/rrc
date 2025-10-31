package ies

// InterFreqTargetInfo-r16 ::= SEQUENCE
type InterfreqtargetinfoR16 struct {
	DlCarrierfreqR16 ArfcnValuenr
	CelllistR16      *[]Physcellid `lb:1,ub:32`
}
