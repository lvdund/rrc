package ies

import "rrc/utils"

// MAC-Parameters-v1430-shortSPS-IntervalTDD-r14 ::= ENUMERATED
type MacParametersV1430ShortspsIntervaltddR14 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1430ShortspsIntervaltddR14EnumeratedNothing = iota
	MacParametersV1430ShortspsIntervaltddR14EnumeratedSupported
)
