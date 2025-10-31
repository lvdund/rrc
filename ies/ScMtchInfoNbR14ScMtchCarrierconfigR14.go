package ies

import "rrc/utils"

// SC-MTCH-Info-NB-r14-sc-mtch-CarrierConfig-r14 ::= CHOICE
const (
	ScMtchInfoNbR14ScMtchCarrierconfigR14ChoiceNothing = iota
	ScMtchInfoNbR14ScMtchCarrierconfigR14ChoiceDlCarrierconfigR14
	ScMtchInfoNbR14ScMtchCarrierconfigR14ChoiceDlCarrierindexR14
)

type ScMtchInfoNbR14ScMtchCarrierconfigR14 struct {
	Choice             uint64
	DlCarrierconfigR14 *DlCarrierconfigcommonNbR14
	DlCarrierindexR14  *utils.INTEGER `lb:0,ub:maxNonAnchorCarriersNbR14`
}
