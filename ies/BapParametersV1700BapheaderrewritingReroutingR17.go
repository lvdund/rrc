package ies

import "rrc/utils"

// BAP-Parameters-v1700-bapHeaderRewriting-Rerouting-r17 ::= ENUMERATED
type BapParametersV1700BapheaderrewritingReroutingR17 struct {
	Value utils.ENUMERATED
}

const (
	BapParametersV1700BapheaderrewritingReroutingR17EnumeratedNothing = iota
	BapParametersV1700BapheaderrewritingReroutingR17EnumeratedSupported
)
