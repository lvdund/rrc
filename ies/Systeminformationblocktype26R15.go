package ies

import "rrc/utils"

// SystemInformationBlockType26-r15 ::= SEQUENCE
// Extensible
type Systeminformationblocktype26R15 struct {
	V2xInterfreqinfolistR15       *SlInterfreqinfolistv2xR14
	CbrPsschTxconfiglistR15       *SlCbrPpppTxconfiglistR15
	V2xPacketduplicationconfigR15 *SlV2xPacketduplicationconfigR15
	SyncfreqlistR15               *SlV2xSyncfreqlistR15
	SlssTxmultifreqR15            *bool
	V2xFreqselectionconfiglistR15 *SlV2xFreqselectionconfiglistR15
	ThreshsRssiCbrR15             *utils.INTEGER `lb:0,ub:45`
	Latenoncriticalextension      *utils.OCTETSTRING
}
