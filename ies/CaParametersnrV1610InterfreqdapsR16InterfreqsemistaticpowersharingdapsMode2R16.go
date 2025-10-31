package ies

import "rrc/utils"

// CA-ParametersNR-v1610-interFreqDAPS-r16-interFreqSemiStaticPowerSharingDAPS-Mode2-r16 ::= ENUMERATED
type CaParametersnrV1610InterfreqdapsR16InterfreqsemistaticpowersharingdapsMode2R16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610InterfreqdapsR16InterfreqsemistaticpowersharingdapsMode2R16EnumeratedNothing = iota
	CaParametersnrV1610InterfreqdapsR16InterfreqsemistaticpowersharingdapsMode2R16EnumeratedSupported
)
