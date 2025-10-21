package ies

import "rrc/utils"

// SystemInformationBlockType2 ::= SEQUENCE
// Extensible
type Systeminformationblocktype2 struct {
	AcBarringinfo             *interface{}
	Radioresourceconfigcommon Radioresourceconfigcommonsib
	UeTimersandconstants      UeTimersandconstants
	Freqinfo                  *interface{}
	MbsfnSubframeconfiglist   *MbsfnSubframeconfiglist
	Timealignmenttimercommon  Timealignmenttimer
	Latenoncriticalextension  *utils.OCTETSTRING
}
