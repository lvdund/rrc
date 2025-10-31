package ies

import "rrc/utils"

// SIB12-r16-segmentType-r16 ::= ENUMERATED
type Sib12R16SegmenttypeR16 struct {
	Value utils.ENUMERATED
}

const (
	Sib12R16SegmenttypeR16EnumeratedNothing = iota
	Sib12R16SegmenttypeR16EnumeratedNotlastsegment
	Sib12R16SegmenttypeR16EnumeratedLastsegment
)
