package ies

import "rrc/utils"

// SystemInformationBlockType15-r11 ::= SEQUENCE
// Extensible
type Systeminformationblocktype15R11 struct {
	MbmsSaiIntrafreqR11      *MbmsSaiListR11
	MbmsSaiInterfreqlistR11  *MbmsSaiInterfreqlistR11
	Latenoncriticalextension *utils.OCTETSTRING
}
