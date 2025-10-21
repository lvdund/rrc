package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v13c0-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV13c0Ies struct {
	RadioresourceconfigdedicatedV13c0 *RadioresourceconfigdedicatedV13c0
	ScelltoaddmodlistV13c0            *ScelltoaddmodlistV13c0
	ScelltoaddmodlistextV13c0         *ScelltoaddmodlistextV13c0
	ScgConfigurationV13c0             *ScgConfigurationV13c0
	Noncriticalextension              *interface{}
}
