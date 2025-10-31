package ies

import "rrc/utils"

// PUR-Config-r16 ::= SEQUENCE
// Extensible
type PurConfigR16 struct {
	PurConfigidR16             *PurConfigidR16
	PurImplicitreleaseafterR16 *PurConfigR16PurImplicitreleaseafterR16
	PurStarttimeparametersR16  *PurConfigR16PurStarttimeparametersR16
	PurNumoccasionsR16         PurConfigR16PurNumoccasionsR16
	PurRntiR16                 *CRnti
	PurTimealignmenttimerR16   *utils.INTEGER `lb:0,ub:8`
	PurRsrpChangethresholdR16  *SetupreleasePurRsrpChangethresholdR16
	PurResponsewindowtimerR16  *PurConfigR16PurResponsewindowtimerR16
	PurMpdcchConfigR16         *PurMpdcchConfigR16
	PurPdschFreqhoppingR16     utils.BOOLEAN
	PurPucchConfigR16          *PurPucchConfigR16
	PurPuschConfigR16          *PurPuschConfigR16
}
