package ies

import "rrc/utils"

// SystemInformationBlockType25-r15 ::= SEQUENCE
// Extensible
type Systeminformationblocktype25R15 struct {
	UacBarringforcommonR15    *UacBarringpercatlistR15
	UacBarringperplmnListR15  *UacBarringperplmnListR15
	UacBarringinfosetlistR15  UacBarringinfosetlistR15
	UacAc1SelectassistinfoR15 *Systeminformationblocktype25R15UacAc1SelectassistinfoR15
	Latenoncriticalextension  *utils.OCTETSTRING
}
