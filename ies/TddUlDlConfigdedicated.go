package ies

// TDD-UL-DL-ConfigDedicated ::= SEQUENCE
// Extensible
type TddUlDlConfigdedicated struct {
	Slotspecificconfigurationstoaddmodlist  *[]TddUlDlSlotconfig `lb:1,ub:maxNrofSlots`
	Slotspecificconfigurationstoreleaselist *[]TddUlDlSlotindex  `lb:1,ub:maxNrofSlots`
}
