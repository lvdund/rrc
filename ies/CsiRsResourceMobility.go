package ies

import "rrc/utils"

// CSI-RS-Resource-Mobility ::= SEQUENCE
// Extensible
type CsiRsResourceMobility struct {
	CsiRsIndex                  CsiRsIndex
	Slotconfig                  CsiRsResourceMobilitySlotconfig
	Associatedssb               *CsiRsResourceMobilityAssociatedssb
	Frequencydomainallocation   CsiRsResourceMobilityFrequencydomainallocation
	Firstofdmsymbolintimedomain utils.INTEGER                       `lb:0,ub:13`
	Sequencegenerationconfig    utils.INTEGER                       `lb:0,ub:1023`
	SlotconfigR17               *CsiRsResourceMobilitySlotconfigR17 `ext`
}
