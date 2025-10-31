package ies

import "rrc/utils"

// FrequencyComponent-r17 ::= ENUMERATED
type FrequencycomponentR17 struct {
	Value utils.ENUMERATED
}

const (
	FrequencycomponentR17EnumeratedNothing = iota
	FrequencycomponentR17EnumeratedActivecarrier
	FrequencycomponentR17EnumeratedConfiguredcarrier
	FrequencycomponentR17EnumeratedActivebwp
	FrequencycomponentR17EnumeratedConfiguredbwp
)
