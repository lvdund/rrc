package ies

import "rrc/utils"

// BandCombinationParameters-r11-simultaneousRx-Tx-r11 ::= ENUMERATED
type BandcombinationparametersR11SimultaneousrxTxR11 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationparametersR11SimultaneousrxTxR11EnumeratedNothing = iota
	BandcombinationparametersR11SimultaneousrxTxR11EnumeratedSupported
)
