package ies

import "rrc/utils"

// SidelinkUEInformationNR-v1700-IEs-ue-Type-r17 ::= ENUMERATED
type SidelinkueinformationnrV1700IesUeTypeR17 struct {
	Value utils.ENUMERATED
}

const (
	SidelinkueinformationnrV1700IesUeTypeR17EnumeratedNothing = iota
	SidelinkueinformationnrV1700IesUeTypeR17EnumeratedRelayue
	SidelinkueinformationnrV1700IesUeTypeR17EnumeratedRemoteue
)
