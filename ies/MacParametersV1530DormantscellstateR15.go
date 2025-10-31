package ies

import "rrc/utils"

// MAC-Parameters-v1530-dormantSCellState-r15 ::= ENUMERATED
type MacParametersV1530DormantscellstateR15 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1530DormantscellstateR15EnumeratedNothing = iota
	MacParametersV1530DormantscellstateR15EnumeratedSupported
)
