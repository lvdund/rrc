package ies

import "rrc/utils"

// BandCombinationParameters-r13-simultaneousRx-Tx-r13 ::= ENUMERATED
type BandcombinationparametersR13SimultaneousrxTxR13 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationparametersR13SimultaneousrxTxR13EnumeratedNothing = iota
	BandcombinationparametersR13SimultaneousrxTxR13EnumeratedSupported
)
