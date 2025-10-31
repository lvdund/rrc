package ies

import "rrc/utils"

// PosGapConfig-r17 ::= SEQUENCE
// Extensible
type PosgapconfigR17 struct {
	MeaspospreconfiggapidR17 MeaspospreconfiggapidR17
	GapoffsetR17             utils.INTEGER `lb:0,ub:159`
	MglR17                   PosgapconfigR17MglR17
	MgrpR17                  PosgapconfigR17MgrpR17
	MgtaR17                  PosgapconfigR17MgtaR17
	GaptypeR17               PosgapconfigR17GaptypeR17
}
