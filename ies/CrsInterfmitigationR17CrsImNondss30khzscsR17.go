package ies

import "rrc/utils"

// CRS-InterfMitigation-r17-crs-IM-nonDSS-30kHzSCS-r17 ::= ENUMERATED
type CrsInterfmitigationR17CrsImNondss30khzscsR17 struct {
	Value utils.ENUMERATED
}

const (
	CrsInterfmitigationR17CrsImNondss30khzscsR17EnumeratedNothing = iota
	CrsInterfmitigationR17CrsImNondss30khzscsR17EnumeratedSupported
)
