package ies

// LAA-SCellConfiguration-v1430 ::= SEQUENCE
type LaaScellconfigurationV1430 struct {
	CrosscarrierschedulingconfigUlR14         *LaaScellconfigurationV1430CrosscarrierschedulingconfigUlR14
	LbtConfigR14                              *LbtConfigR14
	PdcchConfiglaaR14                         *PdcchConfiglaaR14
	AbsenceofanyothertechnologyR14            *bool
	SoundingrsUlConfigdedicatedaperiodicV1430 *SoundingrsUlConfigdedicatedaperiodicV1430
}
