package ies

import "rrc/utils"

// MAC-ParametersCommon-uplink-Harq-ModeB-r17 ::= ENUMERATED
type MacParameterscommonUplinkHarqModebR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonUplinkHarqModebR17EnumeratedNothing = iota
	MacParameterscommonUplinkHarqModebR17EnumeratedSupported
)
