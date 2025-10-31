package ies

import "rrc/utils"

// ServingCellConfig-channelAccessMode2-r17 ::= ENUMERATED
type ServingcellconfigChannelaccessmode2R17 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigChannelaccessmode2R17EnumeratedNothing = iota
	ServingcellconfigChannelaccessmode2R17EnumeratedEnabled
)
