package ies

import "rrc/utils"

// UL-GapFR2-Config-r17-ugl-r17 ::= ENUMERATED
type UlGapfr2ConfigR17UglR17 struct {
	Value utils.ENUMERATED
}

const (
	UlGapfr2ConfigR17UglR17EnumeratedNothing = iota
	UlGapfr2ConfigR17UglR17EnumeratedMs0dot125
	UlGapfr2ConfigR17UglR17EnumeratedMs0dot25
	UlGapfr2ConfigR17UglR17EnumeratedMs0dot5
	UlGapfr2ConfigR17UglR17EnumeratedMs1
)
