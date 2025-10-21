package ies

import "rrc/utils"

// RRCConnectionRelease-NB-v1550-IEs ::= SEQUENCE
type RrcconnectionreleaseNbV1550Ies struct {
	RedirectedcarrierinfoV1550 *RedirectedcarrierinfoNbV1550
	Noncriticalextension       *RrcconnectionreleaseNbV15b0Ies
}
