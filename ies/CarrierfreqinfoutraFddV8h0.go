package ies

// CarrierFreqInfoUTRA-FDD-v8h0 ::= SEQUENCE
type CarrierfreqinfoutraFddV8h0 struct {
	Multibandinfolist *[]FreqbandindicatorUtraFdd `lb:1,ub:maxMultiBands`
}
