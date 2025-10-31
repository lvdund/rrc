package ies

import "rrc/utils"

// MaxRS-IndexCellQualNR-r15 ::= utils.INTEGER (1..maxRS-IndexCellQual-r15)
type MaxrsIndexcellqualnrR15 struct {
	Value utils.INTEGER `lb:0,ub:maxRSIndexcellqualR15`
}
