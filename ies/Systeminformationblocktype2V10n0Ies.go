package ies

import "rrc/utils"

// SystemInformationBlockType2-v10n0-IEs ::= SEQUENCE
type Systeminformationblocktype2V10n0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *Systeminformationblocktype2V13c0Ies
}
