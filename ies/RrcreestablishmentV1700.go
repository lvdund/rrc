package ies

import "rrc/utils"

// RRCReestablishment-v1700-IEs ::= SEQUENCE
type RrcreestablishmentV1700 struct {
	SlL2remoteueConfigR17 *utils.Setuprelease[SlL2remoteueConfigR17]
	Noncriticalextension  *RrcreestablishmentV1700IesNoncriticalextension
}
