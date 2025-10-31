package ies

// RRCConnectionReconfiguration-v920-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV920 struct {
	OtherconfigR9        *OtherconfigR9
	FullconfigR9         *bool
	Noncriticalextension *RrcconnectionreconfigurationV1020
}
