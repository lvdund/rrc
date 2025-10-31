package ies

import "rrc/utils"

// SL-DiscConfig-r17 ::= SEQUENCE
type SlDiscconfigR17 struct {
	SlRelayueConfigR17  *utils.Setuprelease[SlRelayueConfigR17]
	SlRemoteueConfigR17 *utils.Setuprelease[SlRemoteueConfigR17]
}
