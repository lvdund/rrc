package ies

import "rrc/utils"

// SRS-Resources-maxNumberSRS-Ports-PerResource ::= ENUMERATED
type SrsResourcesMaxnumbersrsPortsPerresource struct {
	Value utils.ENUMERATED
}

const (
	SrsResourcesMaxnumbersrsPortsPerresourceEnumeratedNothing = iota
	SrsResourcesMaxnumbersrsPortsPerresourceEnumeratedN1
	SrsResourcesMaxnumbersrsPortsPerresourceEnumeratedN2
	SrsResourcesMaxnumbersrsPortsPerresourceEnumeratedN4
)
