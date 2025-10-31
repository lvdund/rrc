package ies

// TDD-UL-DL-ConfigDedicated-IAB-MT-r16 ::= SEQUENCE
// Extensible
type TddUlDlConfigdedicatedIabMtR16 struct {
	SlotspecificconfigurationstoaddmodlistIabMtR16  *[]TddUlDlSlotconfigIabMtR16 `lb:1,ub:maxNrofSlots`
	SlotspecificconfigurationstoreleaselistIabMtR16 *[]TddUlDlSlotindex          `lb:1,ub:maxNrofSlots`
}
