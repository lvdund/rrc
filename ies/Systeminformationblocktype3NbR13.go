package ies

import "rrc/utils"

// SystemInformationBlockType3-NB-r13 ::= SEQUENCE
// Extensible
type Systeminformationblocktype3NbR13 struct {
	CellreselectioninfocommonR13      interface{}
	CellreselectionservingfreqinfoR13 interface{}
	IntrafreqcellreselectioninfoR13   *interface{}
	FreqbandinfoR13                   *NsPmaxlistNbR13
	MultibandinfolistR13              *interface{}
	Latenoncriticalextension          *utils.OCTETSTRING
}
