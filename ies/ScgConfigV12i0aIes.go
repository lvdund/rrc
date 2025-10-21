package ies

import "rrc/utils"

// SCG-Config-v12i0a-IEs ::= SEQUENCE
type ScgConfigV12i0aIes struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *ScgConfigV13c0Ies
}
