package ies

import "rrc/utils"

// CSI-RS-Config-v1430 ::= SEQUENCE
type CsiRsConfigV1430 struct {
	Dummy                      *CsiRsConfigemimoV1430
	EmimoHybridR14             *CsiRsConfigemimoHybridR14
	AdvancedcodebookenabledR14 *bool
}
