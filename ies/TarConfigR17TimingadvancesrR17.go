package ies

import "rrc/utils"

// TAR-Config-r17-timingAdvanceSR-r17 ::= ENUMERATED
type TarConfigR17TimingadvancesrR17 struct {
	Value utils.ENUMERATED
}

const (
	TarConfigR17TimingadvancesrR17EnumeratedNothing = iota
	TarConfigR17TimingadvancesrR17EnumeratedEnabled
)
