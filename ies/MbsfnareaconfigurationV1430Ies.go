package ies

import "rrc/utils"

// MBSFNAreaConfiguration-v1430-IEs ::= SEQUENCE
type MbsfnareaconfigurationV1430Ies struct {
	CommonsfAllocV1430   CommonsfAllocpatternlistV1430
	Noncriticalextension *MbsfnareaconfigurationV1610Ies
}
