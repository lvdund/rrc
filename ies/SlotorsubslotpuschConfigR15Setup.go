package ies

import "rrc/utils"

// SlotOrSubslotPUSCH-Config-r15-setup ::= SEQUENCE
// Extensible
type SlotorsubslotpuschConfigR15Setup struct {
	BetaoffsetslotAckIndexR15     *utils.INTEGER   `lb:0,ub:15`
	Betaoffset2slotAckIndexR15    *utils.INTEGER   `lb:0,ub:15`
	BetaoffsetsubslotAckIndexR15  *[]utils.INTEGER `lb:1,ub:2`
	Betaoffset2subslotAckIndexR15 *[]utils.INTEGER `lb:1,ub:2`
	BetaoffsetslotRiIndexR15      *utils.INTEGER   `lb:0,ub:15`
	BetaoffsetsubslotRiIndexR15   *[]utils.INTEGER `lb:1,ub:2`
	BetaoffsetslotCqiIndexR15     *utils.INTEGER   `lb:0,ub:15`
	BetaoffsetsubslotCqiIndexR15  *utils.INTEGER   `lb:0,ub:15`
	Enable256qamSlotorsubslotR15  *Enable256qamR14
	ResourceallocationoffsetR15   *utils.INTEGER `lb:0,ub:2`
	UlDmrsIfdmaSlotorsubslotR15   utils.BOOLEAN
}
