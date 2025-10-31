package ies

import "rrc/utils"

// MAC-ParametersCommon-mg-ActivationRequestPRS-Meas-r17 ::= ENUMERATED
type MacParameterscommonMgActivationrequestprsMeasR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonMgActivationrequestprsMeasR17EnumeratedNothing = iota
	MacParameterscommonMgActivationrequestprsMeasR17EnumeratedSupported
)
