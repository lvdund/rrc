package ies

import "rrc/utils"

// SL-PreconfigGeneral-r16 ::= SEQUENCE
// Extensible
type SlPreconfiggeneralR16 struct {
	SlTddConfigurationR16 *TddUlDlConfigcommon
	ReservedbitsR16       *utils.BITSTRING `lb:2,ub:2`
}
