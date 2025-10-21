package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v1370-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV1370Ies struct {
	RadioresourceconfigdedicatedV1370 *RadioresourceconfigdedicatedV1370
	ScelltoaddmodlistextV1370         *ScelltoaddmodlistextV1370
	Noncriticalextension              *RrcconnectionreconfigurationV13c0Ies
}
