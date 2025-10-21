package ies

import "rrc/utils"

// RRCEarlyDataComplete-r15-IEs ::= SEQUENCE
type RrcearlydatacompleteR15Ies struct {
	DedicatedinfonasR15               *Dedicatedinfonas
	ExtendedwaittimeR15               *utils.INTEGER
	IdlemodemobilitycontrolinfoR15    *Idlemodemobilitycontrolinfo
	IdlemodemobilitycontrolinfoextR15 *IdlemodemobilitycontrolinfoV9e0
	RedirectedcarrierinfoR15          *RedirectedcarrierinfoR15Ies
	Noncriticalextension              *RrcearlydatacompleteV1590Ies
}
