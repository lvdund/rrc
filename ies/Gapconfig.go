package ies

import "rrc/utils"

// GapConfig ::= SEQUENCE
// Extensible
type Gapconfig struct {
	Gapoffset                utils.INTEGER `lb:0,ub:159`
	Mgl                      GapconfigMgl
	Mgrp                     GapconfigMgrp
	Mgta                     GapconfigMgta
	Refservcellindicator     *GapconfigRefservcellindicator `ext`
	Reffr2servcellasynccaR16 *Servcellindex                 `ext`
	MglR16                   *GapconfigMglR16               `ext`
}
