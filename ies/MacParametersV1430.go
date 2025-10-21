package ies

import "rrc/utils"

// MAC-Parameters-v1430 ::= SEQUENCE
type MacParametersV1430 struct {
	ShortspsIntervalfddR14 *utils.ENUMERATED
	ShortspsIntervaltddR14 *utils.ENUMERATED
	SkipuplinkdynamicR14   *utils.ENUMERATED
	SkipuplinkspsR14       *utils.ENUMERATED
	MultipleuplinkspsR14   *utils.ENUMERATED
	DatainactmonR14        *utils.ENUMERATED
}
