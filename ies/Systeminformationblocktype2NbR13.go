package ies

import "rrc/utils"

// SystemInformationBlockType2-NB-r13 ::= SEQUENCE
// Extensible
type Systeminformationblocktype2NbR13 struct {
	RadioresourceconfigcommonR13 RadioresourceconfigcommonsibNbR13
	UeTimersandconstantsR13      UeTimersandconstantsNbR13
	FreqinfoR13                  *interface{}
	TimealignmenttimercommonR13  Timealignmenttimer
	MultibandinfolistR13         *interface{}
	Latenoncriticalextension     *utils.OCTETSTRING
}
