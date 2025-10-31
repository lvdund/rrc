package ies

import "rrc/utils"

// MIMOParam-r17-sfnSchemePDCCH-r17 ::= ENUMERATED
type MimoparamR17SfnschemepdcchR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoparamR17SfnschemepdcchR17EnumeratedNothing = iota
	MimoparamR17SfnschemepdcchR17EnumeratedSfnschemea
	MimoparamR17SfnschemepdcchR17EnumeratedSfnschemeb
)
