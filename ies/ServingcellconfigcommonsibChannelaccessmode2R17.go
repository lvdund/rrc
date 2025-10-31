package ies

import "rrc/utils"

// ServingCellConfigCommonSIB-channelAccessMode2-r17 ::= ENUMERATED
type ServingcellconfigcommonsibChannelaccessmode2R17 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigcommonsibChannelaccessmode2R17EnumeratedNothing = iota
	ServingcellconfigcommonsibChannelaccessmode2R17EnumeratedEnabled
)
