package ies

import "rrc/utils"

// ConfigRestrictInfoSCG-nrdc-PC-mode-FR2-r16 ::= ENUMERATED
type ConfigrestrictinfoscgNrdcPcModeFr2R16 struct {
	Value utils.ENUMERATED
}

const (
	ConfigrestrictinfoscgNrdcPcModeFr2R16EnumeratedNothing = iota
	ConfigrestrictinfoscgNrdcPcModeFr2R16EnumeratedSemi_Static_Mode1
	ConfigrestrictinfoscgNrdcPcModeFr2R16EnumeratedSemi_Static_Mode2
	ConfigrestrictinfoscgNrdcPcModeFr2R16EnumeratedDynamic
)
