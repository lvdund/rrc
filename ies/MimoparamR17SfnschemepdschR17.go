package ies

import "rrc/utils"

// MIMOParam-r17-sfnSchemePDSCH-r17 ::= ENUMERATED
type MimoparamR17SfnschemepdschR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoparamR17SfnschemepdschR17EnumeratedNothing = iota
	MimoparamR17SfnschemepdschR17EnumeratedSfnschemea
	MimoparamR17SfnschemepdschR17EnumeratedSfnschemeb
)
