package ies

import "rrc/utils"

// CRS-InterfMitigation-r17-crs-IM-DSS-15kHzSCS-r17 ::= ENUMERATED
type CrsInterfmitigationR17CrsImDss15khzscsR17 struct {
	Value utils.ENUMERATED
}

const (
	CrsInterfmitigationR17CrsImDss15khzscsR17EnumeratedNothing = iota
	CrsInterfmitigationR17CrsImDss15khzscsR17EnumeratedSupported
)
