package ies

// CRS-InterfMitigation-r17 ::= SEQUENCE
type CrsInterfmitigationR17 struct {
	CrsImDss15khzscsR17       *CrsInterfmitigationR17CrsImDss15khzscsR17
	CrsImNondss15khzscsR17    *CrsInterfmitigationR17CrsImNondss15khzscsR17
	CrsImNondssNwa15khzscsR17 *CrsInterfmitigationR17CrsImNondssNwa15khzscsR17
	CrsImNondss30khzscsR17    *CrsInterfmitigationR17CrsImNondss30khzscsR17
	CrsImNondssNwa30khzscsR17 *CrsInterfmitigationR17CrsImNondssNwa30khzscsR17
}
