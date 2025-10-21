package ies

import "rrc/utils"

// RRCConnectionReconfigurationComplete-v1430-IEs ::= SEQUENCE
type RrcconnectionreconfigurationcompleteV1430Ies struct {
	PerccGapindicationlistR14  *PerccGapindicationlistR14
	NumfreqeffectiveR14        *utils.INTEGER
	NumfreqeffectivereducedR14 *utils.INTEGER
	Noncriticalextension       *RrcconnectionreconfigurationcompleteV1510Ies
}
