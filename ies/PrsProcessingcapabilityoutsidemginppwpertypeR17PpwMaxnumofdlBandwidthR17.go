package ies

import "rrc/utils"

// PRS-ProcessingCapabilityOutsideMGinPPWperType-r17-ppw-maxNumOfDL-Bandwidth-r17 ::= CHOICE
const (
	PrsProcessingcapabilityoutsidemginppwpertypeR17PpwMaxnumofdlBandwidthR17ChoiceNothing = iota
	PrsProcessingcapabilityoutsidemginppwpertypeR17PpwMaxnumofdlBandwidthR17ChoiceFr1R17
	PrsProcessingcapabilityoutsidemginppwpertypeR17PpwMaxnumofdlBandwidthR17ChoiceFr2R17
)

type PrsProcessingcapabilityoutsidemginppwpertypeR17PpwMaxnumofdlBandwidthR17 struct {
	Choice uint64
	Fr1R17 *utils.ENUMERATED
	Fr2R17 *utils.ENUMERATED
}
