package ies

import "rrc/utils"

// ServingCellConfigCommonSIB-n-TimingAdvanceOffset ::= ENUMERATED
type ServingcellconfigcommonsibNTimingadvanceoffset struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigcommonsibNTimingadvanceoffsetEnumeratedNothing = iota
	ServingcellconfigcommonsibNTimingadvanceoffsetEnumeratedN0
	ServingcellconfigcommonsibNTimingadvanceoffsetEnumeratedN25600
	ServingcellconfigcommonsibNTimingadvanceoffsetEnumeratedN39936
)
