package ies

// SystemInformation-r8-IEs ::= SEQUENCE
// Extensible
type SysteminformationR8 struct {
	SibTypeandinfo       []SysteminformationR8IesSibTypeandinfoItem `lb:1,ub:maxSIB`
	Noncriticalextension *SysteminformationV8a0
}
