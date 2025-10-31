package ies

import "rrc/utils"

// PDSCH-Config-pdsch-HARQ-ACK-OneShotFeedbackDCI-1-2-r17 ::= ENUMERATED
type PdschConfigPdschHarqAckOneshotfeedbackdci12R17 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigPdschHarqAckOneshotfeedbackdci12R17EnumeratedNothing = iota
	PdschConfigPdschHarqAckOneshotfeedbackdci12R17EnumeratedEnabled
)
