package ies

import "rrc/utils"

// FR2-2-AccessParamsPerBand-r17-type2-ChannelAccess-FR2-2-r17 ::= ENUMERATED
type Fr22AccessparamsperbandR17Type2ChannelaccessFr22R17 struct {
	Value utils.ENUMERATED
}

const (
	Fr22AccessparamsperbandR17Type2ChannelaccessFr22R17EnumeratedNothing = iota
	Fr22AccessparamsperbandR17Type2ChannelaccessFr22R17EnumeratedSupported
)
