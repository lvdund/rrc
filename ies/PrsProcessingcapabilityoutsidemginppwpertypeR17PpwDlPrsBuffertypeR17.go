package ies

import "rrc/utils"

// PRS-ProcessingCapabilityOutsideMGinPPWperType-r17-ppw-dl-PRS-BufferType-r17 ::= utils.ENUMERATED // Extensible
type PrsProcessingcapabilityoutsidemginppwpertypeR17PpwDlPrsBuffertypeR17 struct {
	Value utils.ENUMERATED
}

const (
	PrsProcessingcapabilityoutsidemginppwpertypeR17PpwDlPrsBuffertypeR17EnumeratedNothing = iota
	PrsProcessingcapabilityoutsidemginppwpertypeR17PpwDlPrsBuffertypeR17EnumeratedType1
	PrsProcessingcapabilityoutsidemginppwpertypeR17PpwDlPrsBuffertypeR17EnumeratedType2
)
