package ies

import "rrc/utils"

// SL-V2X-ConfigCommon-r14 ::= SEQUENCE
type SlV2xConfigcommonR14 struct {
	V2xCommrxpoolR14              *SlCommrxpoollistv2xR14
	V2xCommtxpoolnormalcommonR14  *SlCommtxpoollistv2xR14
	P2xCommtxpoolnormalcommonR14  *SlCommtxpoollistv2xR14
	V2xCommtxpoolexceptionalR14   *SlCommresourcepoolv2xR14
	V2xSyncconfigR14              *SlSyncconfiglistv2xR14
	V2xInterfreqinfolistR14       *SlInterfreqinfolistv2xR14
	V2xResourceselectionconfigR14 *SlCommtxpoolsensingconfigR14
	ZoneconfigR14                 *SlZoneconfigR14
	TypetxsyncR14                 *SlTypetxsyncR14
	ThresslTxprioritizationR14    *SlPriorityR13
	AnchorcarrierfreqlistR14      *SlAnchorcarrierfreqlistV2xR14
	OffsetdfnR14                  *utils.INTEGER `lb:0,ub:1000`
	CbrCommontxconfiglistR14      *SlCbrCommontxconfiglistR14
}
