package ies

import "rrc/utils"

// MeasObjectWLAN-r13 ::= SEQUENCE
// Extensible
type MeasobjectwlanR13 struct {
	CarrierfreqR13      *interface{}
	WlanToaddmodlistR13 *WlanIdListR13
	WlanToremovelistR13 *WlanIdListR13
}
