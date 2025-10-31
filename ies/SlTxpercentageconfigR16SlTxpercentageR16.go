package ies

import "rrc/utils"

// SL-TxPercentageConfig-r16-sl-TxPercentage-r16 ::= ENUMERATED
type SlTxpercentageconfigR16SlTxpercentageR16 struct {
	Value utils.ENUMERATED
}

const (
	SlTxpercentageconfigR16SlTxpercentageR16EnumeratedNothing = iota
	SlTxpercentageconfigR16SlTxpercentageR16EnumeratedP20
	SlTxpercentageconfigR16SlTxpercentageR16EnumeratedP35
	SlTxpercentageconfigR16SlTxpercentageR16EnumeratedP50
)
