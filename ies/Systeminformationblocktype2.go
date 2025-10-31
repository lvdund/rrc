package ies

import "rrc/utils"

// SystemInformationBlockType2 ::= SEQUENCE
// Extensible
type Systeminformationblocktype2 struct {
	AcBarringinfo             *Systeminformationblocktype2AcBarringinfo
	Radioresourceconfigcommon Radioresourceconfigcommonsib
	UeTimersandconstants      UeTimersandconstants
	Freqinfo                  *Systeminformationblocktype2Freqinfo
	MbsfnSubframeconfiglist   *MbsfnSubframeconfiglist
	Timealignmenttimercommon  Timealignmenttimer
	Latenoncriticalextension  *utils.OCTETSTRING
}
