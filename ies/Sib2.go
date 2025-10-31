package ies

import "rrc/utils"

// SIB2 ::= SEQUENCE
// Extensible
type Sib2 struct {
	Cellreselectioninfocommon      *Sib2Cellreselectioninfocommon
	Cellreselectionservingfreqinfo *Sib2Cellreselectionservingfreqinfo
	Intrafreqcellreselectioninfo   *Sib2Intrafreqcellreselectioninfo `ext`
	RelaxedmeasurementR16          *Sib2RelaxedmeasurementR16        `ext`
	CellequivalentsizeR17          *utils.INTEGER                    `lb:0,ub:16,ext`
	RelaxedmeasurementR17          *Sib2RelaxedmeasurementR17        `ext`
}
