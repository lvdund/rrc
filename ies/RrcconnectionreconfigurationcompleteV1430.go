package ies

import "rrc/utils"

// RRCConnectionReconfigurationComplete-v1430-IEs ::= SEQUENCE
type RrcconnectionreconfigurationcompleteV1430 struct {
	PerccGapindicationlistR14  *PerccGapindicationlistR14
	NumfreqeffectiveR14        *utils.INTEGER `lb:0,ub:12`
	NumfreqeffectivereducedR14 *utils.INTEGER `lb:0,ub:12`
	Noncriticalextension       *RrcconnectionreconfigurationcompleteV1510
}
