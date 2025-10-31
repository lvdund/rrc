package ies

import "rrc/utils"

// MAC-Parameters-v1530-directSCellHibernation-r15 ::= ENUMERATED
type MacParametersV1530DirectscellhibernationR15 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1530DirectscellhibernationR15EnumeratedNothing = iota
	MacParametersV1530DirectscellhibernationR15EnumeratedSupported
)
