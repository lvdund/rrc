package ies

import "rrc/utils"

// SystemInformationBlockType3 ::= SEQUENCE
// Extensible
type Systeminformationblocktype3 struct {
	Cellreselectioninfocommon      *interface{}
	Cellreselectionservingfreqinfo *interface{}
	Intrafreqcellreselectioninfo   *interface{}
	Latenoncriticalextension       *utils.OCTETSTRING
}
