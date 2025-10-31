package ies

import "rrc/utils"

// CSI-RS-Resource-Mobility-slotConfig-r17 ::= CHOICE
const (
	CsiRsResourceMobilitySlotconfigR17ChoiceNothing = iota
	CsiRsResourceMobilitySlotconfigR17ChoiceMs4
	CsiRsResourceMobilitySlotconfigR17ChoiceMs5
	CsiRsResourceMobilitySlotconfigR17ChoiceMs10
	CsiRsResourceMobilitySlotconfigR17ChoiceMs20
	CsiRsResourceMobilitySlotconfigR17ChoiceMs40
)

type CsiRsResourceMobilitySlotconfigR17 struct {
	Choice uint64
	Ms4    *utils.INTEGER `lb:0,ub:255`
	Ms5    *utils.INTEGER `lb:0,ub:319`
	Ms10   *utils.INTEGER `lb:0,ub:639`
	Ms20   *utils.INTEGER `lb:0,ub:1279`
	Ms40   *utils.INTEGER `lb:0,ub:2559`
}
