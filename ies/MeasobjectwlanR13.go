package ies

// MeasObjectWLAN-r13 ::= SEQUENCE
// Extensible
type MeasobjectwlanR13 struct {
	CarrierfreqR13      *MeasobjectwlanR13CarrierfreqR13
	WlanToaddmodlistR13 *WlanIdListR13
	WlanToremovelistR13 *WlanIdListR13
}
