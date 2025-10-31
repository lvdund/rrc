package ies

import "rrc/utils"

// SL-V2X-PreconfigCommPool-r14 ::= SEQUENCE
// Extensible
type SlV2xPreconfigcommpoolR14 struct {
	SlOffsetindicatorR14                 *SlOffsetindicatorR12
	SlSubframeR14                        SubframebitmapslR14
	AdjacencypscchPsschR14               utils.BOOLEAN
	SizesubchannelR14                    SlV2xPreconfigcommpoolR14SizesubchannelR14
	NumsubchannelR14                     SlV2xPreconfigcommpoolR14NumsubchannelR14
	StartrbSubchannelR14                 utils.INTEGER  `lb:0,ub:99`
	StartrbPscchPoolR14                  *utils.INTEGER `lb:0,ub:99`
	DatatxparametersR14                  P0SlR12
	ZoneidR14                            *utils.INTEGER `lb:0,ub:7`
	ThreshsRssiCbrR14                    *utils.INTEGER `lb:0,ub:45`
	CbrPsschTxconfiglistR14              *SlCbrPpppTxpreconfiglistR14
	Resourceselectionconfigp2xR14        *SlP2xResourceselectionconfigR14
	SyncallowedR14                       *SlSyncallowedR14
	RestrictresourcereservationperiodR14 *SlRestrictresourcereservationperiodlistR14
}
