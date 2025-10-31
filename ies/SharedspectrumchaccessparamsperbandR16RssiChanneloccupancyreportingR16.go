package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-rssi-ChannelOccupancyReporting-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16RssiChanneloccupancyreportingR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16RssiChanneloccupancyreportingR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16RssiChanneloccupancyreportingR16EnumeratedSupported
)
