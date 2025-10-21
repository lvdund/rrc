package ies

import "rrc/utils"

// SystemInformation-r8-IEs ::= SEQUENCE
// Extensible
type SysteminformationR8Ies struct {
	SibTypeandinfo       interface{}
	Noncriticalextension *SysteminformationV8a0Ies
}
