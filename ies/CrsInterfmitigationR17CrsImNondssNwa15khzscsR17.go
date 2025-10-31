package ies

import "rrc/utils"

// CRS-InterfMitigation-r17-crs-IM-nonDSS-NWA-15kHzSCS-r17 ::= ENUMERATED
type CrsInterfmitigationR17CrsImNondssNwa15khzscsR17 struct {
	Value utils.ENUMERATED
}

const (
	CrsInterfmitigationR17CrsImNondssNwa15khzscsR17EnumeratedNothing = iota
	CrsInterfmitigationR17CrsImNondssNwa15khzscsR17EnumeratedSupported
)
