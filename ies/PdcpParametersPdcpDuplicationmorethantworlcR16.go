package ies

import "rrc/utils"

// PDCP-Parameters-pdcp-DuplicationMoreThanTwoRLC-r16 ::= ENUMERATED
type PdcpParametersPdcpDuplicationmorethantworlcR16 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersPdcpDuplicationmorethantworlcR16EnumeratedNothing = iota
	PdcpParametersPdcpDuplicationmorethantworlcR16EnumeratedSupported
)
