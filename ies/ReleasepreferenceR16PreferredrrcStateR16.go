package ies

import "rrc/utils"

// ReleasePreference-r16-preferredRRC-State-r16 ::= ENUMERATED
type ReleasepreferenceR16PreferredrrcStateR16 struct {
	Value utils.ENUMERATED
}

const (
	ReleasepreferenceR16PreferredrrcStateR16EnumeratedNothing = iota
	ReleasepreferenceR16PreferredrrcStateR16EnumeratedIdle
	ReleasepreferenceR16PreferredrrcStateR16EnumeratedInactive
	ReleasepreferenceR16PreferredrrcStateR16EnumeratedConnected
	ReleasepreferenceR16PreferredrrcStateR16EnumeratedOutofconnected
)
