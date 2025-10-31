package ies

import "rrc/utils"

// AS-Config ::= SEQUENCE
// Extensible
type AsConfig struct {
	Rrcreconfiguration   utils.OCTETSTRING
	SourcerbSnConfig     *utils.OCTETSTRING `ext`
	SourcescgNrConfig    *utils.OCTETSTRING `ext`
	SourcescgEutraConfig *utils.OCTETSTRING `ext`
	SourcescgConfigured  *utils.BOOLEAN     `ext`
	SdtConfigR17         *SdtConfigR17      `ext`
}
