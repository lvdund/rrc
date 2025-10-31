package ies

import "rrc/utils"

// PRACH-Config-v1430 ::= SEQUENCE
type PrachConfigV1430 struct {
	RootsequenceindexhighspeedR14         utils.INTEGER `lb:0,ub:837`
	ZerocorrelationzoneconfighighspeedR14 utils.INTEGER `lb:0,ub:12`
	PrachConfigindexhighspeedR14          utils.INTEGER `lb:0,ub:63`
	PrachFreqoffsethighspeedR14           utils.INTEGER `lb:0,ub:94`
}
