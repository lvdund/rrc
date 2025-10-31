package ies

import "rrc/utils"

// GapConfig-r17 ::= SEQUENCE
// Extensible
type GapconfigR17 struct {
	MeasgapidR17             MeasgapidR17
	GaptypeR17               GapconfigR17GaptypeR17
	GapoffsetR17             utils.INTEGER `lb:0,ub:159`
	MglR17                   GapconfigR17MglR17
	MgrpR17                  GapconfigR17MgrpR17
	MgtaR17                  GapconfigR17MgtaR17
	RefservcellindicatorR17  *GapconfigR17RefservcellindicatorR17
	Reffr2ServcellasynccaR17 *Servcellindex
	PreconfigindR17          *utils.BOOLEAN
	NcsgindR17               *utils.BOOLEAN
	GapassociationprsR17     *utils.BOOLEAN
	GapsharingR17            *Measgapsharingscheme
	GappriorityR17           *GappriorityR17
}
