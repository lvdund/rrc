package ies

import "rrc/utils"

// CSI-RS-ConfigNZP-Activation-r14-csi-RS-NZP-mode-r14 ::= ENUMERATED
type CsiRsConfignzpActivationR14CsiRsNzpModeR14 struct {
	Value utils.ENUMERATED
}

const (
	CsiRsConfignzpActivationR14CsiRsNzpModeR14EnumeratedNothing = iota
	CsiRsConfignzpActivationR14CsiRsNzpModeR14EnumeratedSemipersistent
	CsiRsConfignzpActivationR14CsiRsNzpModeR14EnumeratedAperiodic
)
