package ies

import "rrc/utils"

// DRB-CountInfo ::= SEQUENCE
type DrbCountinfo struct {
	DrbIdentity   DrbIdentity
	CountUplink   utils.INTEGER `lb:0,ub:4294967295`
	CountDownlink utils.INTEGER `lb:0,ub:4294967295`
}
