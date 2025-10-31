package ies

import "rrc/utils"

// MAC-Parameters-v1430-dataInactMon-r14 ::= ENUMERATED
type MacParametersV1430DatainactmonR14 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1430DatainactmonR14EnumeratedNothing = iota
	MacParametersV1430DatainactmonR14EnumeratedSupported
)
