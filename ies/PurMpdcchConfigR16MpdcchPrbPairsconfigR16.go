package ies

import "rrc/utils"

// PUR-MPDCCH-Config-r16-mpdcch-PRB-PairsConfig-r16 ::= SEQUENCE
type PurMpdcchConfigR16MpdcchPrbPairsconfigR16 struct {
	NumberprbPairsR16          PurMpdcchConfigR16MpdcchPrbPairsconfigR16NumberprbPairsR16
	ResourceblockassignmentR16 utils.BITSTRING `lb:4,ub:4`
}
