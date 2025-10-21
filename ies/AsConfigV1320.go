package ies

import "rrc/utils"

// AS-Config-v1320 ::= SEQUENCE
type AsConfigV1320 struct {
	SourcescellconfiglistR13    *ScelltoaddmodlistextR13
	SourcerclwiConfigurationR13 *RclwiConfigurationR13
}
