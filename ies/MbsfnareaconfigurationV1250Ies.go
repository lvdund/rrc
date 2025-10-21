package ies

import "rrc/utils"

// MBSFNAreaConfiguration-v1250-IEs ::= SEQUENCE
type MbsfnareaconfigurationV1250Ies struct {
	PmchInfolistextR12   *PmchInfolistextR12
	Noncriticalextension *MbsfnareaconfigurationV1430Ies
}
