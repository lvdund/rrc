package ies

import "rrc/utils"

// SystemInformationBlockType3 ::= SEQUENCE
// Extensible
type Systeminformationblocktype3 struct {
	Cellreselectioninfocommon      *Systeminformationblocktype3Cellreselectioninfocommon
	Cellreselectionservingfreqinfo *Systeminformationblocktype3Cellreselectionservingfreqinfo
	Intrafreqcellreselectioninfo   *Systeminformationblocktype3Intrafreqcellreselectioninfo
	Latenoncriticalextension       *utils.OCTETSTRING
}
