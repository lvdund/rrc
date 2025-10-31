package ies

import "rrc/utils"

// BandSidelinkEUTRA-r16-gnb-ScheduledMode4SidelinkEUTRA-r16 ::= ENUMERATED
type BandsidelinkeutraR16GnbScheduledmode4sidelinkeutraR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkeutraR16GnbScheduledmode4sidelinkeutraR16EnumeratedNothing = iota
	BandsidelinkeutraR16GnbScheduledmode4sidelinkeutraR16EnumeratedSupported
)
