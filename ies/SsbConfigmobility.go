package ies

import "rrc/utils"

// SSB-ConfigMobility ::= SEQUENCE
// Extensible
type SsbConfigmobility struct {
	SsbTomeasure                       *utils.Setuprelease[SsbTomeasure]
	DerivessbIndexfromcell             utils.BOOLEAN
	SsRssiMeasurement                  *SsRssiMeasurement
	SsbPositionqclCommonR16            *SsbPositionqclRelationR16                     `ext`
	SsbPositionqclCellstoaddmodlistR16 *SsbPositionqclCellstoaddmodlistR16            `ext`
	SsbPositionqclCellstoremovelistR16 *PciList                                       `ext`
	DerivessbIndexfromcellinterR17     *Servcellindex                                 `ext`
	SsbPositionqclCommonR17            *SsbPositionqclRelationR17                     `ext`
	SsbPositionqclCellsR17             *utils.Setuprelease[SsbPositionqclCelllistR17] `ext`
	CcaCellstoaddmodlistR17            *PciList                                       `ext`
	CcaCellstoremovelistR17            *PciList                                       `ext`
}
