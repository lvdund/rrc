package ies

import "rrc/utils"

// BWP-UplinkCommon ::= SEQUENCE
// Extensible
type BwpUplinkcommon struct {
	Genericparameters                   Bwp
	RachConfigcommon                    *utils.Setuprelease[RachConfigcommon]
	PuschConfigcommon                   *utils.Setuprelease[PuschConfigcommon]
	PucchConfigcommon                   *utils.Setuprelease[PucchConfigcommon]
	RachConfigcommoniabR16              *utils.Setuprelease[RachConfigcommon]            `ext`
	UseinterlacepucchPuschR16           *BwpUplinkcommonUseinterlacepucchPuschR16        `ext`
	MsgaConfigcommonR16                 *utils.Setuprelease[MsgaConfigcommonR16]         `ext`
	EnableraPrioritizationforslicingR17 *utils.BOOLEAN                                   `ext`
	AdditionalrachConfiglistR17         *utils.Setuprelease[AdditionalrachConfiglistR17] `ext`
	RsrpThresholdmsg3R17                *RsrpRange                                       `ext`
	Numberofmsg3RepetitionslistR17      *[]Numberofmsg3RepetitionsR17                    `lb:4,ub:4,ext`
	McsMsg3RepetitionsR17               *[]utils.INTEGER                                 `lb:8,ub:8,ext`
}
