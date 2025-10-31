package ies

import "rrc/utils"

// TDD-UL-DL-Pattern ::= SEQUENCE
// Extensible
type TddUlDlPattern struct {
	DlUlTransmissionperiodicity      TddUlDlPatternDlUlTransmissionperiodicity
	Nrofdownlinkslots                utils.INTEGER                                   `lb:0,ub:maxNrofSlots`
	Nrofdownlinksymbols              utils.INTEGER                                   `lb:0,ub:maxNrofSymbols1`
	Nrofuplinkslots                  utils.INTEGER                                   `lb:0,ub:maxNrofSlots`
	Nrofuplinksymbols                utils.INTEGER                                   `lb:0,ub:maxNrofSymbols1`
	DlUlTransmissionperiodicityV1530 *TddUlDlPatternDlUlTransmissionperiodicityV1530 `ext`
}
