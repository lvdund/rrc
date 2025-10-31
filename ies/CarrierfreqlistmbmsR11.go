package ies

// CarrierFreqListMBMS-r11 ::= SEQUENCE OF ARFCN-ValueEUTRA-r9
// SIZE (1..maxFreqMBMS-r11)
type CarrierfreqlistmbmsR11 struct {
	Value []ArfcnValueeutraR9 `lb:1,ub:maxFreqMBMSR11`
}
