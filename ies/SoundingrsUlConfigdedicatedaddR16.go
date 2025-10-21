package ies

import "rrc/utils"

// SoundingRS-UL-ConfigDedicatedAdd-r16 ::= SEQUENCE
type SoundingrsUlConfigdedicatedaddR16 struct {
	SrsConfigindexapR16      utils.INTEGER
	SrsConfigapdciFormat4R16 *interface{}
	SrsActivateapR13         *interface{}
}
