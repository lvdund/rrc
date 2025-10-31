package ies

import "rrc/utils"

// UL-PDCP-DelayResult-r13-qci-Id-r13 ::= ENUMERATED
type UlPdcpDelayresultR13QciIdR13 struct {
	Value utils.ENUMERATED
}

const (
	UlPdcpDelayresultR13QciIdR13EnumeratedNothing = iota
	UlPdcpDelayresultR13QciIdR13EnumeratedQci1
	UlPdcpDelayresultR13QciIdR13EnumeratedQci2
	UlPdcpDelayresultR13QciIdR13EnumeratedQci3
	UlPdcpDelayresultR13QciIdR13EnumeratedQci4
	UlPdcpDelayresultR13QciIdR13EnumeratedSpare4
	UlPdcpDelayresultR13QciIdR13EnumeratedSpare3
	UlPdcpDelayresultR13QciIdR13EnumeratedSpare2
	UlPdcpDelayresultR13QciIdR13EnumeratedSpare1
)
