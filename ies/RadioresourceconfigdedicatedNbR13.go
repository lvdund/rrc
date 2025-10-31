package ies

// RadioResourceConfigDedicated-NB-r13 ::= SEQUENCE
// Extensible
type RadioresourceconfigdedicatedNbR13 struct {
	SrbToaddmodlistR13         *SrbToaddmodlistNbR13
	DrbToaddmodlistR13         *DrbToaddmodlistNbR13
	DrbToreleaselistR13        *DrbToreleaselistNbR13
	MacMainconfigR13           *RadioresourceconfigdedicatedNbR13MacMainconfigR13
	PhysicalconfigdedicatedR13 *PhysicalconfigdedicatedNbR13
	RlfTimersandconstantsR13   *RlfTimersandconstantsNbR13
}
