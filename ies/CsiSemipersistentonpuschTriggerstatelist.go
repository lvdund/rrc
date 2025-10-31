package ies

// CSI-SemiPersistentOnPUSCH-TriggerStateList ::= SEQUENCE OF CSI-SemiPersistentOnPUSCH-TriggerState
// SIZE (1..maxNrOfSemiPersistentPUSCH-Triggers)
type CsiSemipersistentonpuschTriggerstatelist struct {
	Value []CsiSemipersistentonpuschTriggerstate `lb:1,ub:maxNrOfSemiPersistentPUSCHTriggers`
}
