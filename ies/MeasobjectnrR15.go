package ies

import "rrc/utils"

// MeasObjectNR-r15 ::= SEQUENCE
// Extensible
type MeasobjectnrR15 struct {
	CarrierfreqR15               ArfcnValuenrR15
	RsConfigssbR15               RsConfigssbNrR15
	ThreshrsIndexR15             *ThresholdlistnrR15
	MaxrsIndexcellqualR15        *MaxrsIndexcellqualnrR15
	OffsetfreqR15                QOffsetrangeinterrat
	BlackcellstoremovelistR15    *Cellindexlist
	BlackcellstoaddmodlistR15    *CellstoaddmodlistnrR15
	QuantityconfigsetR15         utils.INTEGER      `lb:0,ub:maxQuantSetsNRR15`
	CellsforwhichtoreportsftdR15 *[]PhyscellidnrR15 `lb:1,ub:maxCellSFTD`
}
