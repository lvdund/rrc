package ies

import "rrc/utils"

// RRCEarlyDataComplete-NB-r15-IEs ::= SEQUENCE
type RrcearlydatacompleteNbR15 struct {
	DedicatedinfonasR15         *Dedicatedinfonas
	ExtendedwaittimeR15         *utils.INTEGER `lb:0,ub:1800`
	RedirectedcarrierinfoR15    *RedirectedcarrierinfoNbR13
	RedirectedcarrierinfoextR15 *RedirectedcarrierinfoNbV1430
	Noncriticalextension        *RrcearlydatacompleteNbV1590
}
