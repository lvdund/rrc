package ies

import "rrc/utils"

// UE-NR-Capability-v1700-ntn-ScenarioSupport-r17 ::= ENUMERATED
type UeNrCapabilityV1700NtnScenariosupportR17 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1700NtnScenariosupportR17EnumeratedNothing = iota
	UeNrCapabilityV1700NtnScenariosupportR17EnumeratedGso
	UeNrCapabilityV1700NtnScenariosupportR17EnumeratedNgso
)
