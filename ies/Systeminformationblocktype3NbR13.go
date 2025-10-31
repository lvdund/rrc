package ies

import "rrc/utils"

// SystemInformationBlockType3-NB-r13 ::= SEQUENCE
// Extensible
type Systeminformationblocktype3NbR13 struct {
	CellreselectioninfocommonR13      Systeminformationblocktype3NbR13CellreselectioninfocommonR13
	CellreselectionservingfreqinfoR13 Systeminformationblocktype3NbR13CellreselectionservingfreqinfoR13
	IntrafreqcellreselectioninfoR13   *Systeminformationblocktype3NbR13IntrafreqcellreselectioninfoR13
	FreqbandinfoR13                   *NsPmaxlistNbR13
	MultibandinfolistR13              *[]NsPmaxlistNbR13 `lb:1,ub:maxMultiBands`
	Latenoncriticalextension          *utils.OCTETSTRING
}
