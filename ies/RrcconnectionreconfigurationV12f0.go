package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v12f0-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV12f0 struct {
	ScgConfigurationV12f0    *ScgConfigurationV12f0
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionreconfigurationV1370
}
