package ies

import "rrc/utils"

// CE-MultiTB-Parameters-r16-ce-MultiTB-HARQ-AckBundling-r16 ::= ENUMERATED
type CeMultitbParametersR16CeMultitbHarqAckbundlingR16 struct {
	Value utils.ENUMERATED
}

const (
	CeMultitbParametersR16CeMultitbHarqAckbundlingR16EnumeratedNothing = iota
	CeMultitbParametersR16CeMultitbHarqAckbundlingR16EnumeratedSupported
)
