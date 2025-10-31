package ies

import "rrc/utils"

// SoundingRS-VirtualCellID-r16 ::= SEQUENCE
type SoundingrsVirtualcellidR16 struct {
	SrsVirtualcellidR16       utils.INTEGER `lb:0,ub:503`
	SrsVirtualcellidAllsrsR16 utils.BOOLEAN
}
