package ies

// RRCConnectionReconfiguration-v1310-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV1310 struct {
	ScelltoreleaselistextR13 *ScelltoreleaselistextR13
	ScelltoaddmodlistextR13  *ScelltoaddmodlistextR13
	LwaConfigurationR13      *LwaConfigurationR13
	LwipConfigurationR13     *LwipConfigurationR13
	RclwiConfigurationR13    *RclwiConfigurationR13
	Noncriticalextension     *RrcconnectionreconfigurationV1430
}
