package ies

import "rrc/utils"

// RSS-Config-r15 ::= SEQUENCE
type RssConfigR15 struct {
	DurationR15     RssConfigR15DurationR15
	FreqlocationR15 utils.INTEGER `lb:0,ub:98`
	PeriodicityR15  RssConfigR15PeriodicityR15
	PowerboostR15   RssConfigR15PowerboostR15
	TimeoffsetR15   utils.INTEGER `lb:0,ub:31`
}
