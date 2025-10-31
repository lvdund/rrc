package ies

import "rrc/utils"

// CSI-RS-ForTracking ::= SEQUENCE
type CsiRsFortracking struct {
	Maxburstlength                   utils.INTEGER `lb:0,ub:2`
	Maxsimultaneousresourcesetspercc utils.INTEGER `lb:0,ub:8`
	Maxconfiguredresourcesetspercc   utils.INTEGER `lb:0,ub:64`
	Maxconfiguredresourcesetsallcc   utils.INTEGER `lb:0,ub:256`
}
