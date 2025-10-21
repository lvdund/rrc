package ies

import "rrc/utils"

// LAA-SCellConfiguration-v1430 ::= SEQUENCE
type LaaScellconfigurationV1430 struct {
	CrosscarrierschedulingconfigUlR14         *interface{}
	LbtConfigR14                              *LbtConfigR14
	PdcchConfiglaaR14                         *PdcchConfiglaaR14
	AbsenceofanyothertechnologyR14            *utils.ENUMERATED
	SoundingrsUlConfigdedicatedaperiodicV1430 *SoundingrsUlConfigdedicatedaperiodicV1430
}
