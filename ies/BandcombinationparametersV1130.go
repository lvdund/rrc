package ies

import "rrc/utils"

// BandCombinationParameters-v1130 ::= SEQUENCE
// Extensible
type BandcombinationparametersV1130 struct {
	MultipletimingadvanceR11 *utils.ENUMERATED
	SimultaneousrxTxR11      *utils.ENUMERATED
	BandparameterlistR11     *interface{}
}
