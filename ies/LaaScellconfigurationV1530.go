package ies

import "rrc/utils"

// LAA-SCellConfiguration-v1530 ::= SEQUENCE
type LaaScellconfigurationV1530 struct {
	AulConfigR15          *AulConfigR15
	PuschModeconfiglaaR15 *PuschModeconfiglaaR15
}
