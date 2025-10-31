package ies

// RRCConnectionReconfiguration-v1370-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV1370 struct {
	RadioresourceconfigdedicatedV1370 *RadioresourceconfigdedicatedV1370
	ScelltoaddmodlistextV1370         *ScelltoaddmodlistextV1370
	Noncriticalextension              *RrcconnectionreconfigurationV13c0
}
