package ies

// CSI-AperiodicTriggerStateList ::= SEQUENCE OF CSI-AperiodicTriggerState
// SIZE (1..maxNrOfCSI-AperiodicTriggers)
type CsiAperiodictriggerstatelist struct {
	Value []CsiAperiodictriggerstate `lb:1,ub:maxNrOfCSIAperiodictriggers`
}
