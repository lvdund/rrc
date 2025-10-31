package ies

import "rrc/utils"

// CA-ParametersNR-v1610-interFreqDAPS-r16-interFreqSemiStaticPowerSharingDAPS-Mode1-r16 ::= ENUMERATED
type CaParametersnrV1610InterfreqdapsR16InterfreqsemistaticpowersharingdapsMode1R16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610InterfreqdapsR16InterfreqsemistaticpowersharingdapsMode1R16EnumeratedNothing = iota
	CaParametersnrV1610InterfreqdapsR16InterfreqsemistaticpowersharingdapsMode1R16EnumeratedSupported
)
