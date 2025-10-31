package ies

import "rrc/utils"

// SL-V2X-PreconfigFreqInfo-r14 ::= SEQUENCE
// Extensible
type SlV2xPreconfigfreqinfoR14 struct {
	V2xCommpreconfiggeneralR14    SlPreconfiggeneralR12
	V2xCommpreconfigsyncR14       *SlPreconfigv2xSyncR14
	V2xCommrxpoollistR14          SlPreconfigv2xRxpoollistR14
	V2xCommtxpoollistR14          SlPreconfigv2xTxpoollistR14
	P2xCommtxpoollistR14          SlPreconfigv2xTxpoollistR14
	V2xResourceselectionconfigR14 *SlCommtxpoolsensingconfigR14
	ZoneconfigR14                 *SlZoneconfigR14
	SyncpriorityR14               SlV2xPreconfigfreqinfoR14SyncpriorityR14
	ThresslTxprioritizationR14    *SlPriorityR13
	OffsetdfnR14                  *utils.INTEGER `lb:0,ub:1000`
}
