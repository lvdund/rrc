package ies

import "rrc/utils"

// RRCConnectionRelease-NB-v1430-IEs ::= SEQUENCE
type RrcconnectionreleaseNbV1430 struct {
	RedirectedcarrierinfoV1430 *RedirectedcarrierinfoNbV1430
	ExtendedwaittimeCpdataR14  *utils.INTEGER `lb:0,ub:1800`
	Noncriticalextension       *RrcconnectionreleaseNbV1530
}
