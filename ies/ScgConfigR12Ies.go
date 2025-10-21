package ies

import "rrc/utils"

// SCG-Config-r12-IEs ::= SEQUENCE
type ScgConfigR12Ies struct {
	ScgRadioconfigR12    *ScgConfigpartscgR12
	Noncriticalextension *ScgConfigV12i0aIes
}
