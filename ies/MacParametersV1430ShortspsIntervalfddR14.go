package ies

import "rrc/utils"

// MAC-Parameters-v1430-shortSPS-IntervalFDD-r14 ::= ENUMERATED
type MacParametersV1430ShortspsIntervalfddR14 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1430ShortspsIntervalfddR14EnumeratedNothing = iota
	MacParametersV1430ShortspsIntervalfddR14EnumeratedSupported
)
