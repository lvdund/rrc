package ies

import "rrc/utils"

// BandCombinationParameters-v1130-simultaneousRx-Tx-r11 ::= ENUMERATED
type BandcombinationparametersV1130SimultaneousrxTxR11 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationparametersV1130SimultaneousrxTxR11EnumeratedNothing = iota
	BandcombinationparametersV1130SimultaneousrxTxR11EnumeratedSupported
)
