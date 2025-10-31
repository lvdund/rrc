package ies

import "rrc/utils"

// MBMS-SessionInfo-r9 ::= SEQUENCE
// Extensible
type MbmsSessioninfoR9 struct {
	TmgiR9                   TmgiR9
	SessionidR9              *utils.OCTETSTRING `lb:1,ub:1`
	LogicalchannelidentityR9 utils.INTEGER      `lb:0,ub:maxSessionPerPMCH1`
}
