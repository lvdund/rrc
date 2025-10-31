package ies

import "rrc/utils"

// MCCH-Config-r17 ::= SEQUENCE
type McchConfigR17 struct {
	McchRepetitionperiodandoffsetR17 McchRepetitionperiodandoffsetR17
	McchWindowstartslotR17           utils.INTEGER `lb:0,ub:79`
	McchWindowdurationR17            *McchConfigR17McchWindowdurationR17
	McchModificationperiodR17        McchConfigR17McchModificationperiodR17
}
