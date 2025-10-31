package ies

import "rrc/utils"

// LTE-NeighCellsCRS-AssistInfo-r17 ::= SEQUENCE
type LteNeighcellscrsAssistinfoR17 struct {
	NeighcarrierbandwidthdlR17      *LteNeighcellscrsAssistinfoR17NeighcarrierbandwidthdlR17
	NeighcarrierfreqdlR17           *utils.INTEGER `lb:0,ub:16383`
	NeighcellidR17                  *EutraPhyscellid
	NeighcrsMutingR17               *LteNeighcellscrsAssistinfoR17NeighcrsMutingR17
	NeighmbsfnSubframeconfiglistR17 *EutraMbsfnSubframeconfiglist
	NeighnrofcrsPortsR17            *LteNeighcellscrsAssistinfoR17NeighnrofcrsPortsR17
	NeighvShiftR17                  *LteNeighcellscrsAssistinfoR17NeighvShiftR17
}
