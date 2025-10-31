package ies

import "rrc/utils"

// SystemInformationBlockType2-v10n0-IEs ::= SEQUENCE
type Systeminformationblocktype2V10n0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *Systeminformationblocktype2V13c0
}
