package ies

import "rrc/utils"

// PUR-Config-NB-r16 ::= SEQUENCE
// Extensible
type PurConfigNbR16 struct {
	PurConfigidR16             *PurConfigidNbR16
	PurTimealignmenttimerR16   *utils.INTEGER `lb:0,ub:8`
	PurNrsrpChangethresholdR16 *SetupreleasePurNrsrpChangethresholdR16
	PurImplicitreleaseafterR16 *PurConfigNbR16PurImplicitreleaseafterR16
	PurRntiR16                 *CRnti
	PurResponsewindowtimerR16  *PurConfigNbR16PurResponsewindowtimerR16
	PurStarttimeparametersR16  *PurConfigNbR16PurStarttimeparametersR16
	PurNumoccasionsR16         PurConfigNbR16PurNumoccasionsR16
	PurPhysicalconfigR16       *PurConfigNbR16PurPhysicalconfigR16
}
