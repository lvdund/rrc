package ies

import "rrc/utils"

// ReconfigurationWithSync-t304 ::= ENUMERATED
type ReconfigurationwithsyncT304 struct {
	Value utils.ENUMERATED
}

const (
	ReconfigurationwithsyncT304EnumeratedNothing = iota
	ReconfigurationwithsyncT304EnumeratedMs50
	ReconfigurationwithsyncT304EnumeratedMs100
	ReconfigurationwithsyncT304EnumeratedMs150
	ReconfigurationwithsyncT304EnumeratedMs200
	ReconfigurationwithsyncT304EnumeratedMs500
	ReconfigurationwithsyncT304EnumeratedMs1000
	ReconfigurationwithsyncT304EnumeratedMs2000
	ReconfigurationwithsyncT304EnumeratedMs10000
)
