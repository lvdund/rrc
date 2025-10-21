package ies

import "rrc/utils"

// RRCConnectionRelease-NB-v1430-IEs ::= SEQUENCE
type RrcconnectionreleaseNbV1430Ies struct {
	RedirectedcarrierinfoV1430 *RedirectedcarrierinfoNbV1430
	ExtendedwaittimeCpdataR14  *utils.INTEGER
	Noncriticalextension       *RrcconnectionreleaseNbV1530Ies
}
