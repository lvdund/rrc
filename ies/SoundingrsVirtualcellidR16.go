package ies

import "rrc/utils"

// SoundingRS-VirtualCellID-r16 ::= SEQUENCE
type SoundingrsVirtualcellidR16 struct {
	SrsVirtualcellidR16       utils.INTEGER
	SrsVirtualcellidAllsrsR16 bool
}
