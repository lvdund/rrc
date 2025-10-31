package ies

import "rrc/utils"

// UL-GapFR2-Config-r17 ::= SEQUENCE
type UlGapfr2ConfigR17 struct {
	GapoffsetR17             utils.INTEGER `lb:0,ub:159`
	UglR17                   UlGapfr2ConfigR17UglR17
	UgrpR17                  UlGapfr2ConfigR17UgrpR17
	Reffr2ServcellasynccaR17 *Servcellindex
}
