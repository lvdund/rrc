package ies

import "rrc/utils"

// DC-Parameters-v1310-pdcp-TransferSplitUL-r13 ::= ENUMERATED
type DcParametersV1310PdcpTransfersplitulR13 struct {
	Value utils.ENUMERATED
}

const (
	DcParametersV1310PdcpTransfersplitulR13EnumeratedNothing = iota
	DcParametersV1310PdcpTransfersplitulR13EnumeratedSupported
)
