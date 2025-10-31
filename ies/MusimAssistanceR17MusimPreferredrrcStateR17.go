package ies

import "rrc/utils"

// MUSIM-Assistance-r17-musim-PreferredRRC-State-r17 ::= ENUMERATED
type MusimAssistanceR17MusimPreferredrrcStateR17 struct {
	Value utils.ENUMERATED
}

const (
	MusimAssistanceR17MusimPreferredrrcStateR17EnumeratedNothing = iota
	MusimAssistanceR17MusimPreferredrrcStateR17EnumeratedIdle
	MusimAssistanceR17MusimPreferredrrcStateR17EnumeratedInactive
	MusimAssistanceR17MusimPreferredrrcStateR17EnumeratedOutofconnected
)
