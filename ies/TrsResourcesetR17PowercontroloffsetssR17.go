package ies

import "rrc/utils"

// TRS-ResourceSet-r17-powerControlOffsetSS-r17 ::= ENUMERATED
type TrsResourcesetR17PowercontroloffsetssR17 struct {
	Value utils.ENUMERATED
}

const (
	TrsResourcesetR17PowercontroloffsetssR17EnumeratedNothing = iota
	TrsResourcesetR17PowercontroloffsetssR17EnumeratedDb_3
	TrsResourcesetR17PowercontroloffsetssR17EnumeratedDb0
	TrsResourcesetR17PowercontroloffsetssR17EnumeratedDb3
	TrsResourcesetR17PowercontroloffsetssR17EnumeratedDb6
)
