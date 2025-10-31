package ies

import "rrc/utils"

// SystemInformationBlockType28-r16-segmentType-r16 ::= ENUMERATED
type Systeminformationblocktype28R16SegmenttypeR16 struct {
	Value utils.ENUMERATED
}

const (
	Systeminformationblocktype28R16SegmenttypeR16EnumeratedNothing = iota
	Systeminformationblocktype28R16SegmenttypeR16EnumeratedNotlastsegment
	Systeminformationblocktype28R16SegmenttypeR16EnumeratedLastsegment
)
