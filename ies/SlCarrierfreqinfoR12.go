package ies

// SL-CarrierFreqInfo-r12 ::= SEQUENCE
type SlCarrierfreqinfoR12 struct {
	CarrierfreqR12      ArfcnValueeutraR9
	PlmnIdentitylistR12 *PlmnIdentitylist4R12
}
