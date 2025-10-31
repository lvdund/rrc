package ies

import "rrc/utils"

// ServingCellConfig-cbg-TxDiffTBsProcessingType2-r16 ::= ENUMERATED
type ServingcellconfigCbgTxdifftbsprocessingtype2R16 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigCbgTxdifftbsprocessingtype2R16EnumeratedNothing = iota
	ServingcellconfigCbgTxdifftbsprocessingtype2R16EnumeratedEnabled
)
