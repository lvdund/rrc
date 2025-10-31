package ies

import "rrc/utils"

// SL-PHY-MAC-RLC-Config-r16-networkControlledSyncTx-r16 ::= ENUMERATED
type SlPhyMacRlcConfigR16NetworkcontrolledsynctxR16 struct {
	Value utils.ENUMERATED
}

const (
	SlPhyMacRlcConfigR16NetworkcontrolledsynctxR16EnumeratedNothing = iota
	SlPhyMacRlcConfigR16NetworkcontrolledsynctxR16EnumeratedOn
	SlPhyMacRlcConfigR16NetworkcontrolledsynctxR16EnumeratedOff
)
