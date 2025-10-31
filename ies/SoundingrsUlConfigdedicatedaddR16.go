package ies

import "rrc/utils"

// SoundingRS-UL-ConfigDedicatedAdd-r16 ::= SEQUENCE
type SoundingrsUlConfigdedicatedaddR16 struct {
	SrsConfigindexapR16      utils.INTEGER      `lb:0,ub:31`
	SrsConfigapdciFormat4R16 *[]SrsConfigaddR16 `lb:1,ub:3`
	SrsActivateapR13         *SoundingrsUlConfigdedicatedaddR16SrsActivateapR13
}
