package ies

import "rrc/utils"

// SystemInformationBlockType15-NB-r14 ::= SEQUENCE
// Extensible
type Systeminformationblocktype15NbR14 struct {
	MbmsSaiIntrafreqR14      *MbmsSaiListR11
	MbmsSaiInterfreqlistR14  *MbmsSaiInterfreqlistNbR14
	Latenoncriticalextension *utils.OCTETSTRING
}
