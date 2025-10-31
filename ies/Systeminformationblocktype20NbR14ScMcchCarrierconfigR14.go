package ies

import "rrc/utils"

// SystemInformationBlockType20-NB-r14-sc-mcch-CarrierConfig-r14 ::= CHOICE
const (
	Systeminformationblocktype20NbR14ScMcchCarrierconfigR14ChoiceNothing = iota
	Systeminformationblocktype20NbR14ScMcchCarrierconfigR14ChoiceDlCarrierconfigR14
	Systeminformationblocktype20NbR14ScMcchCarrierconfigR14ChoiceDlCarrierindexR14
)

type Systeminformationblocktype20NbR14ScMcchCarrierconfigR14 struct {
	Choice             uint64
	DlCarrierconfigR14 *DlCarrierconfigcommonNbR14
	DlCarrierindexR14  *utils.INTEGER `lb:0,ub:maxNonAnchorCarriersNbR14`
}
