package ies

import "rrc/utils"

// CRS-InterfMitigation-r17-crs-IM-nonDSS-15kHzSCS-r17 ::= ENUMERATED
type CrsInterfmitigationR17CrsImNondss15khzscsR17 struct {
	Value utils.ENUMERATED
}

const (
	CrsInterfmitigationR17CrsImNondss15khzscsR17EnumeratedNothing = iota
	CrsInterfmitigationR17CrsImNondss15khzscsR17EnumeratedSupported
)
