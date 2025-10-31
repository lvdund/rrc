package ies

import "rrc/utils"

// ServingCellConfig-cbg-TxDiffTBsProcessingType1-r16 ::= ENUMERATED
type ServingcellconfigCbgTxdifftbsprocessingtype1R16 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigCbgTxdifftbsprocessingtype1R16EnumeratedNothing = iota
	ServingcellconfigCbgTxdifftbsprocessingtype1R16EnumeratedEnabled
)
