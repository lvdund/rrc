package ies

import "rrc/utils"

// DRB-CountMSB-Info ::= SEQUENCE
type DrbCountmsbInfo struct {
	DrbIdentity      DrbIdentity
	CountmsbUplink   utils.INTEGER `lb:0,ub:33554431`
	CountmsbDownlink utils.INTEGER `lb:0,ub:33554431`
}
