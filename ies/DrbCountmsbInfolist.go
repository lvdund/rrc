package ies

import "rrc/utils"

// DRB-CountMSB-InfoList ::= SEQUENCE OF DRB-CountMSB-Info
// SIZE (1..maxDRB)
type DrbCountmsbInfolist struct {
	Value []DrbCountmsbInfo
}
