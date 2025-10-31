package ies

import "rrc/utils"

// LTE-NeighCellsCRS-AssistInfo-r17-neighNrofCRS-Ports-r17 ::= ENUMERATED
type LteNeighcellscrsAssistinfoR17NeighnrofcrsPortsR17 struct {
	Value utils.ENUMERATED
}

const (
	LteNeighcellscrsAssistinfoR17NeighnrofcrsPortsR17EnumeratedNothing = iota
	LteNeighcellscrsAssistinfoR17NeighnrofcrsPortsR17EnumeratedN1
	LteNeighcellscrsAssistinfoR17NeighnrofcrsPortsR17EnumeratedN2
	LteNeighcellscrsAssistinfoR17NeighnrofcrsPortsR17EnumeratedN4
)
