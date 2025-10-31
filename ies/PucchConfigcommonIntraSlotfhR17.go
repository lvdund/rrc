package ies

import "rrc/utils"

// PUCCH-ConfigCommon-intra-SlotFH-r17 ::= ENUMERATED
type PucchConfigcommonIntraSlotfhR17 struct {
	Value utils.ENUMERATED
}

const (
	PucchConfigcommonIntraSlotfhR17EnumeratedNothing = iota
	PucchConfigcommonIntraSlotfhR17EnumeratedFromloweredge
	PucchConfigcommonIntraSlotfhR17EnumeratedFromupperedge
)
