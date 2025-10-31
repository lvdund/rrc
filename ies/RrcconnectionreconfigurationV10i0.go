package ies

// RRCConnectionReconfiguration-v10i0-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV10i0 struct {
	AntennainfodedicatedpcellV10i0 *AntennainfodedicatedV10i0
	Noncriticalextension           *RrcconnectionreconfigurationV10l0
}
