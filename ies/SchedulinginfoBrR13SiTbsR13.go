package ies

import "rrc/utils"

// SchedulingInfo-BR-r13-si-TBS-r13 ::= ENUMERATED
type SchedulinginfoBrR13SiTbsR13 struct {
	Value utils.ENUMERATED
}

const (
	SchedulinginfoBrR13SiTbsR13EnumeratedNothing = iota
	SchedulinginfoBrR13SiTbsR13EnumeratedB152
	SchedulinginfoBrR13SiTbsR13EnumeratedB208
	SchedulinginfoBrR13SiTbsR13EnumeratedB256
	SchedulinginfoBrR13SiTbsR13EnumeratedB328
	SchedulinginfoBrR13SiTbsR13EnumeratedB408
	SchedulinginfoBrR13SiTbsR13EnumeratedB504
	SchedulinginfoBrR13SiTbsR13EnumeratedB600
	SchedulinginfoBrR13SiTbsR13EnumeratedB712
	SchedulinginfoBrR13SiTbsR13EnumeratedB808
	SchedulinginfoBrR13SiTbsR13EnumeratedB936
)
