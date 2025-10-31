package ies

import "rrc/utils"

// PDCP-Config-drb ::= SEQUENCE
// Extensible
type PdcpConfigDrb struct {
	Discardtimer         *PdcpConfigDrbDiscardtimer
	PdcpSnSizeul         *PdcpConfigDrbPdcpSnSizeul
	PdcpSnSizedl         *PdcpConfigDrbPdcpSnSizedl
	Headercompression    *PdcpConfigDrbHeadercompression
	Integrityprotection  *PdcpConfigDrbIntegrityprotection
	Statusreportrequired *utils.BOOLEAN
	Outoforderdelivery   *utils.BOOLEAN
}
