package ies

import "rrc/utils"

// CRS-InterfMitigation-r17-crs-IM-nonDSS-NWA-30kHzSCS-r17 ::= ENUMERATED
type CrsInterfmitigationR17CrsImNondssNwa30khzscsR17 struct {
	Value utils.ENUMERATED
}

const (
	CrsInterfmitigationR17CrsImNondssNwa30khzscsR17EnumeratedNothing = iota
	CrsInterfmitigationR17CrsImNondssNwa30khzscsR17EnumeratedSupported
)
