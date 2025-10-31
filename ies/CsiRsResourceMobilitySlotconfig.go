package ies

import "rrc/utils"

// CSI-RS-Resource-Mobility-slotConfig ::= CHOICE
const (
	CsiRsResourceMobilitySlotconfigChoiceNothing = iota
	CsiRsResourceMobilitySlotconfigChoiceMs4
	CsiRsResourceMobilitySlotconfigChoiceMs5
	CsiRsResourceMobilitySlotconfigChoiceMs10
	CsiRsResourceMobilitySlotconfigChoiceMs20
	CsiRsResourceMobilitySlotconfigChoiceMs40
)

type CsiRsResourceMobilitySlotconfig struct {
	Choice uint64
	Ms4    *utils.INTEGER `lb:0,ub:31`
	Ms5    *utils.INTEGER `lb:0,ub:39`
	Ms10   *utils.INTEGER `lb:0,ub:79`
	Ms20   *utils.INTEGER `lb:0,ub:159`
	Ms40   *utils.INTEGER `lb:0,ub:319`
}
