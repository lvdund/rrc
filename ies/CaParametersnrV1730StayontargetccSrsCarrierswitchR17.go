package ies

import "rrc/utils"

// CA-ParametersNR-v1730-stayOnTargetCC-SRS-CarrierSwitch-r17 ::= ENUMERATED
type CaParametersnrV1730StayontargetccSrsCarrierswitchR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1730StayontargetccSrsCarrierswitchR17EnumeratedNothing = iota
	CaParametersnrV1730StayontargetccSrsCarrierswitchR17EnumeratedSupported
)
