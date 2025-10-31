package ies

// RRCConnectionReconfiguration-v13c0-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV13c0 struct {
	RadioresourceconfigdedicatedV13c0 *RadioresourceconfigdedicatedV13c0
	ScelltoaddmodlistV13c0            *ScelltoaddmodlistV13c0
	ScelltoaddmodlistextV13c0         *ScelltoaddmodlistextV13c0
	ScgConfigurationV13c0             *ScgConfigurationV13c0
	Noncriticalextension              *RrcconnectionreconfigurationV13c0IesNoncriticalextension
}
