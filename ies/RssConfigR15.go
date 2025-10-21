package ies

import "rrc/utils"

// RSS-Config-r15 ::= SEQUENCE
type RssConfigR15 struct {
	DurationR15     utils.ENUMERATED
	FreqlocationR15 utils.INTEGER
	PeriodicityR15  utils.ENUMERATED
	PowerboostR15   utils.ENUMERATED
	TimeoffsetR15   utils.INTEGER
}
