package ies

import "rrc/utils"

// ConfigRestrictInfoSCG-nrdc-PC-mode-FR1-r16 ::= ENUMERATED
type ConfigrestrictinfoscgNrdcPcModeFr1R16 struct {
	Value utils.ENUMERATED
}

const (
	ConfigrestrictinfoscgNrdcPcModeFr1R16EnumeratedNothing = iota
	ConfigrestrictinfoscgNrdcPcModeFr1R16EnumeratedSemi_Static_Mode1
	ConfigrestrictinfoscgNrdcPcModeFr1R16EnumeratedSemi_Static_Mode2
	ConfigrestrictinfoscgNrdcPcModeFr1R16EnumeratedDynamic
)
