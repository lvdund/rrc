package ies

import "rrc/utils"

// CellReselectionInfoCommon-v1460-s-SearchDeltaP-r14 ::= ENUMERATED
type CellreselectioninfocommonV1460SSearchdeltapR14 struct {
	Value utils.ENUMERATED
}

const (
	CellreselectioninfocommonV1460SSearchdeltapR14EnumeratedNothing = iota
	CellreselectioninfocommonV1460SSearchdeltapR14EnumeratedDb6
	CellreselectioninfocommonV1460SSearchdeltapR14EnumeratedDb9
	CellreselectioninfocommonV1460SSearchdeltapR14EnumeratedDb12
	CellreselectioninfocommonV1460SSearchdeltapR14EnumeratedDb15
)
