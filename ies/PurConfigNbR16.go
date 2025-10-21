package ies

import "rrc/utils"

// PUR-Config-NB-r16 ::= SEQUENCE
// Extensible
type PurConfigNbR16 struct {
	PurConfigidR16             *PurConfigidNbR16
	PurTimealignmenttimerR16   *utils.INTEGER
	PurNrsrpChangethresholdR16 *Setuprelease
	PurImplicitreleaseafterR16 *utils.ENUMERATED
	PurRntiR16                 *CRnti
	PurResponsewindowtimerR16  *utils.ENUMERATED
	PurStarttimeparametersR16  *interface{}
	PurNumoccasionsR16         utils.ENUMERATED
	PurPhysicalconfigR16       *interface{}
}
