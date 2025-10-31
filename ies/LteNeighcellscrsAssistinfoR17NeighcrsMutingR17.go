package ies

import "rrc/utils"

// LTE-NeighCellsCRS-AssistInfo-r17-neighCRS-muting-r17 ::= ENUMERATED
type LteNeighcellscrsAssistinfoR17NeighcrsMutingR17 struct {
	Value utils.ENUMERATED
}

const (
	LteNeighcellscrsAssistinfoR17NeighcrsMutingR17EnumeratedNothing = iota
	LteNeighcellscrsAssistinfoR17NeighcrsMutingR17EnumeratedEnabled
)
