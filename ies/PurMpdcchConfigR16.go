package ies

import "rrc/utils"

// PUR-MPDCCH-Config-r16 ::= SEQUENCE
type PurMpdcchConfigR16 struct {
	MpdcchFreqhoppingR16    utils.BOOLEAN
	MpdcchNarrowbandR16     utils.INTEGER `lb:0,ub:maxAvailNarrowBandsR13`
	MpdcchPrbPairsconfigR16 PurMpdcchConfigR16MpdcchPrbPairsconfigR16
	MpdcchNumrepetitionR16  PurMpdcchConfigR16MpdcchNumrepetitionR16
	MpdcchStartsfUessR16    PurMpdcchConfigR16MpdcchStartsfUessR16
	MpdcchOffsetPurSsR16    PurMpdcchConfigR16MpdcchOffsetPurSsR16
}
