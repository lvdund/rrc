package ies

import "rrc/utils"

// SL-PDCP-ConfigPC5-r16 ::= SEQUENCE
// Extensible
type SlPdcpConfigpc5R16 struct {
	SlPdcpSnSizeR16         *SlPdcpConfigpc5R16SlPdcpSnSizeR16
	SlOutoforderdeliveryR16 *utils.BOOLEAN
}
