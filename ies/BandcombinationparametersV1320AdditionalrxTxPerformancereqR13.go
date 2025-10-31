package ies

import "rrc/utils"

// BandCombinationParameters-v1320-additionalRx-Tx-PerformanceReq-r13 ::= ENUMERATED
type BandcombinationparametersV1320AdditionalrxTxPerformancereqR13 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationparametersV1320AdditionalrxTxPerformancereqR13EnumeratedNothing = iota
	BandcombinationparametersV1320AdditionalrxTxPerformancereqR13EnumeratedSupported
)
