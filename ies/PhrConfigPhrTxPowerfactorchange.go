package ies

import "rrc/utils"

// PHR-Config-phr-Tx-PowerFactorChange ::= ENUMERATED
type PhrConfigPhrTxPowerfactorchange struct {
	Value utils.ENUMERATED
}

const (
	PhrConfigPhrTxPowerfactorchangeEnumeratedNothing = iota
	PhrConfigPhrTxPowerfactorchangeEnumeratedDb1
	PhrConfigPhrTxPowerfactorchangeEnumeratedDb3
	PhrConfigPhrTxPowerfactorchangeEnumeratedDb6
	PhrConfigPhrTxPowerfactorchangeEnumeratedInfinity
)
