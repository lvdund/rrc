package ies

import "rrc/utils"

// SIB17-r17-segmentType-r17 ::= ENUMERATED
type Sib17R17SegmenttypeR17 struct {
	Value utils.ENUMERATED
}

const (
	Sib17R17SegmenttypeR17EnumeratedNothing = iota
	Sib17R17SegmenttypeR17EnumeratedNotlastsegment
	Sib17R17SegmenttypeR17EnumeratedLastsegment
)
