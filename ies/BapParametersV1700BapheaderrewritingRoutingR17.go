package ies

import "rrc/utils"

// BAP-Parameters-v1700-bapHeaderRewriting-Routing-r17 ::= ENUMERATED
type BapParametersV1700BapheaderrewritingRoutingR17 struct {
	Value utils.ENUMERATED
}

const (
	BapParametersV1700BapheaderrewritingRoutingR17EnumeratedNothing = iota
	BapParametersV1700BapheaderrewritingRoutingR17EnumeratedSupported
)
