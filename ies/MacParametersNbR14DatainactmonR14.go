package ies

import "rrc/utils"

// MAC-Parameters-NB-r14-dataInactMon-r14 ::= ENUMERATED
type MacParametersNbR14DatainactmonR14 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersNbR14DatainactmonR14EnumeratedNothing = iota
	MacParametersNbR14DatainactmonR14EnumeratedSupported
)
