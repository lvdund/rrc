package ies

import "rrc/utils"

// RedCap-ConfigCommonSIB-r17 ::= SEQUENCE
// Extensible
type RedcapConfigcommonsibR17 struct {
	HalfduplexredcapallowedR17 *utils.BOOLEAN
	CellbarredredcapR17        *RedcapConfigcommonsibR17CellbarredredcapR17
}
