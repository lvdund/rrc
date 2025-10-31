package ies

import "rrc/utils"

// RRCEarlyDataComplete-r15-IEs ::= SEQUENCE
type RrcearlydatacompleteR15 struct {
	DedicatedinfonasR15               *Dedicatedinfonas
	ExtendedwaittimeR15               *utils.INTEGER `lb:0,ub:1800`
	IdlemodemobilitycontrolinfoR15    *Idlemodemobilitycontrolinfo
	IdlemodemobilitycontrolinfoextR15 *IdlemodemobilitycontrolinfoV9e0
	RedirectedcarrierinfoR15          *RedirectedcarrierinfoR15
	Noncriticalextension              *RrcearlydatacompleteV1590
}
