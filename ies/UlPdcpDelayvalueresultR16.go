package ies

import "rrc/utils"

// UL-PDCP-DelayValueResult-r16 ::= SEQUENCE
// Extensible
type UlPdcpDelayvalueresultR16 struct {
	DrbIdR16        DrbIdentity
	AveragedelayR16 utils.INTEGER `lb:0,ub:10000`
}
