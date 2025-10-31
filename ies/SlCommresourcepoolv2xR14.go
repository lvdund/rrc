package ies

import "rrc/utils"

// SL-CommResourcePoolV2X-r14 ::= SEQUENCE
// Extensible
type SlCommresourcepoolv2xR14 struct {
	SlOffsetindicatorR14                 *SlOffsetindicatorR12
	SlSubframeR14                        SubframebitmapslR14
	AdjacencypscchPsschR14               utils.BOOLEAN
	SizesubchannelR14                    SlCommresourcepoolv2xR14SizesubchannelR14
	NumsubchannelR14                     SlCommresourcepoolv2xR14NumsubchannelR14
	StartrbSubchannelR14                 utils.INTEGER  `lb:0,ub:99`
	StartrbPscchPoolR14                  *utils.INTEGER `lb:0,ub:99`
	RxparametersncellR14                 *SlCommresourcepoolv2xR14RxparametersncellR14
	DatatxparametersR14                  *SlTxparametersR12
	ZoneidR14                            *utils.INTEGER `lb:0,ub:7`
	ThreshsRssiCbrR14                    *utils.INTEGER `lb:0,ub:45`
	PoolreportidR14                      *SlV2xTxpoolreportidentityR14
	CbrPsschTxconfiglistR14              *SlCbrPpppTxconfiglistR14
	Resourceselectionconfigp2xR14        *SlP2xResourceselectionconfigR14
	SyncallowedR14                       *SlSyncallowedR14
	RestrictresourcereservationperiodR14 *SlRestrictresourcereservationperiodlistR14
}
