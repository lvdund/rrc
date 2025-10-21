package ies

import "rrc/utils"

// PUR-MPDCCH-Config-r16 ::= SEQUENCE
type PurMpdcchConfigR16 struct {
	MpdcchFreqhoppingR16    bool
	MpdcchNarrowbandR16     utils.INTEGER
	MpdcchPrbPairsconfigR16 utils.ENUMERATED
	MpdcchNumrepetitionR16  utils.ENUMERATED
	MpdcchStartsfUessR16    interface{}
	MpdcchOffsetPurSsR16    utils.ENUMERATED
}
