package ies

import "rrc/utils"

// SL-CommResourcePoolV2X-r14 ::= SEQUENCE
// Extensible
type SlCommresourcepoolv2xR14 struct {
	SlOffsetindicatorR14                 *SlOffsetindicatorR12
	SlSubframeR14                        SubframebitmapslR14
	AdjacencypscchPsschR14               bool
	SizesubchannelR14                    utils.ENUMERATED
	NumsubchannelR14                     utils.ENUMERATED
	StartrbSubchannelR14                 utils.INTEGER
	StartrbPscchPoolR14                  *utils.INTEGER
	RxparametersncellR14                 *interface{}
	DatatxparametersR14                  *SlTxparametersR12
	ZoneidR14                            *utils.INTEGER
	ThreshsRssiCbrR14                    *utils.INTEGER
	PoolreportidR14                      *SlV2xTxpoolreportidentityR14
	CbrPsschTxconfiglistR14              *SlCbrPpppTxconfiglistR14
	Resourceselectionconfigp2xR14        *SlP2xResourceselectionconfigR14
	SyncallowedR14                       *SlSyncallowedR14
	RestrictresourcereservationperiodR14 *SlRestrictresourcereservationperiodlistR14
}
