package ies

import "rrc/utils"

// CI-ConfigurationPerServingCell-r16-uplinkCancellationPriority-v1610 ::= ENUMERATED
type CiConfigurationperservingcellR16UplinkcancellationpriorityV1610 struct {
	Value utils.ENUMERATED
}

const (
	CiConfigurationperservingcellR16UplinkcancellationpriorityV1610EnumeratedNothing = iota
	CiConfigurationperservingcellR16UplinkcancellationpriorityV1610EnumeratedEnabled
)
