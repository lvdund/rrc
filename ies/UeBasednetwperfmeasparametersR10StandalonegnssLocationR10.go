package ies

import "rrc/utils"

// UE-BasedNetwPerfMeasParameters-r10-standaloneGNSS-Location-r10 ::= ENUMERATED
type UeBasednetwperfmeasparametersR10StandalonegnssLocationR10 struct {
	Value utils.ENUMERATED
}

const (
	UeBasednetwperfmeasparametersR10StandalonegnssLocationR10EnumeratedNothing = iota
	UeBasednetwperfmeasparametersR10StandalonegnssLocationR10EnumeratedSupported
)
