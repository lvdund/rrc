package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v1020-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV1020Ies struct {
	ScelltoreleaselistR10 *ScelltoreleaselistR10
	ScelltoaddmodlistR10  *ScelltoaddmodlistR10
	Noncriticalextension  *RrcconnectionreconfigurationV1130Ies
}
