package ies

import "rrc/utils"

// SIB2-cellReselectionInfoCommon ::= SEQUENCE
// Extensible
type Sib2Cellreselectioninfocommon struct {
	NrofssBlockstoaverage          *utils.INTEGER `lb:0,ub:maxNrofSSBlockstoaverage`
	AbsthreshssBlocksconsolidation *Thresholdnr
	Rangetobestcell                *Rangetobestcell
	QHyst                          Sib2CellreselectioninfocommonQHyst
	Speedstatereselectionpars      *Sib2CellreselectioninfocommonSpeedstatereselectionpars
}
