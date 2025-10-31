package ies

import "rrc/utils"

// MAC-ParametersCommon-mg-ActivationCommPRS-Meas-r17 ::= ENUMERATED
type MacParameterscommonMgActivationcommprsMeasR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonMgActivationcommprsMeasR17EnumeratedNothing = iota
	MacParameterscommonMgActivationcommprsMeasR17EnumeratedSupported
)
