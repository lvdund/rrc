package ies

import "rrc/utils"

// SchedulingInfo-NB-r13-si-TB-r13 ::= ENUMERATED
type SchedulinginfoNbR13SiTbR13 struct {
	Value utils.ENUMERATED
}

const (
	SchedulinginfoNbR13SiTbR13EnumeratedNothing = iota
	SchedulinginfoNbR13SiTbR13EnumeratedB56
	SchedulinginfoNbR13SiTbR13EnumeratedB120
	SchedulinginfoNbR13SiTbR13EnumeratedB208
	SchedulinginfoNbR13SiTbR13EnumeratedB256
	SchedulinginfoNbR13SiTbR13EnumeratedB328
	SchedulinginfoNbR13SiTbR13EnumeratedB440
	SchedulinginfoNbR13SiTbR13EnumeratedB552
	SchedulinginfoNbR13SiTbR13EnumeratedB680
)
