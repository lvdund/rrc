package ies

import "rrc/utils"

// ControlResourceSet-tci-PresentInDCI ::= ENUMERATED
type ControlresourcesetTciPresentindci struct {
	Value utils.ENUMERATED
}

const (
	ControlresourcesetTciPresentindciEnumeratedNothing = iota
	ControlresourcesetTciPresentindciEnumeratedEnabled
)
