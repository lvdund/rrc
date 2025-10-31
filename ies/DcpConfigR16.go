package ies

import "rrc/utils"

// DCP-Config-r16 ::= SEQUENCE
type DcpConfigR16 struct {
	PsRntiR16                     RntiValue
	PsOffsetR16                   utils.INTEGER `lb:0,ub:120`
	Sizedci26R16                  utils.INTEGER `lb:0,ub:maxDCI26SizeR16`
	PsPositiondci26R16            utils.INTEGER `lb:0,ub:maxDCI26Size1R16`
	PsWakeupR16                   *utils.BOOLEAN
	PsTransmitperiodicl1RsrpR16   *utils.BOOLEAN
	PsTransmitotherperiodiccsiR16 *utils.BOOLEAN
}
