package ies

import "rrc/utils"

// UL-GapFR2-Config-r17-ugrp-r17 ::= ENUMERATED
type UlGapfr2ConfigR17UgrpR17 struct {
	Value utils.ENUMERATED
}

const (
	UlGapfr2ConfigR17UgrpR17EnumeratedNothing = iota
	UlGapfr2ConfigR17UgrpR17EnumeratedMs5
	UlGapfr2ConfigR17UgrpR17EnumeratedMs20
	UlGapfr2ConfigR17UgrpR17EnumeratedMs40
	UlGapfr2ConfigR17UgrpR17EnumeratedMs160
)
