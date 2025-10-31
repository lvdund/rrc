package ies

import "rrc/utils"

// UE-BasedNetwPerfMeasParameters-v1610-ul-PDCP-AvgDelay-r16 ::= ENUMERATED
type UeBasednetwperfmeasparametersV1610UlPdcpAvgdelayR16 struct {
	Value utils.ENUMERATED
}

const (
	UeBasednetwperfmeasparametersV1610UlPdcpAvgdelayR16EnumeratedNothing = iota
	UeBasednetwperfmeasparametersV1610UlPdcpAvgdelayR16EnumeratedSupported
)
