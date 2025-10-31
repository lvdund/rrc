package ies

import "rrc/utils"

// ServingCellConfigCommon-channelAccessMode2-r17 ::= ENUMERATED
type ServingcellconfigcommonChannelaccessmode2R17 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigcommonChannelaccessmode2R17EnumeratedNothing = iota
	ServingcellconfigcommonChannelaccessmode2R17EnumeratedEnabled
)
