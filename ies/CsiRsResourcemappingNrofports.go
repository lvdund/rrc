package ies

import "rrc/utils"

// CSI-RS-ResourceMapping-nrofPorts ::= ENUMERATED
type CsiRsResourcemappingNrofports struct {
	Value utils.ENUMERATED
}

const (
	CsiRsResourcemappingNrofportsEnumeratedNothing = iota
	CsiRsResourcemappingNrofportsEnumeratedP1
	CsiRsResourcemappingNrofportsEnumeratedP2
	CsiRsResourcemappingNrofportsEnumeratedP4
	CsiRsResourcemappingNrofportsEnumeratedP8
	CsiRsResourcemappingNrofportsEnumeratedP12
	CsiRsResourcemappingNrofportsEnumeratedP16
	CsiRsResourcemappingNrofportsEnumeratedP24
	CsiRsResourcemappingNrofportsEnumeratedP32
)
