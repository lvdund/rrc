package ies

import "rrc/utils"

// SL-HoppingConfigComm-r12-numSubbands-r12 ::= ENUMERATED
type SlHoppingconfigcommR12NumsubbandsR12 struct {
	Value utils.ENUMERATED
}

const (
	SlHoppingconfigcommR12NumsubbandsR12EnumeratedNothing = iota
	SlHoppingconfigcommR12NumsubbandsR12EnumeratedNs1
	SlHoppingconfigcommR12NumsubbandsR12EnumeratedNs2
	SlHoppingconfigcommR12NumsubbandsR12EnumeratedNs4
)
