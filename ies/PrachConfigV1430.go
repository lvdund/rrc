package ies

import "rrc/utils"

// PRACH-Config-v1430 ::= SEQUENCE
type PrachConfigV1430 struct {
	RootsequenceindexhighspeedR14         utils.INTEGER
	ZerocorrelationzoneconfighighspeedR14 utils.INTEGER
	PrachConfigindexhighspeedR14          utils.INTEGER
	PrachFreqoffsethighspeedR14           utils.INTEGER
}
