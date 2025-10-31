package ies

import "rrc/utils"

// MAC-MainConfig-phr-Config-setup-dl-PathlossChange ::= ENUMERATED
type MacMainconfigPhrConfigSetupDlPathlosschange struct {
	Value utils.ENUMERATED
}

const (
	MacMainconfigPhrConfigSetupDlPathlosschangeEnumeratedNothing = iota
	MacMainconfigPhrConfigSetupDlPathlosschangeEnumeratedDb1
	MacMainconfigPhrConfigSetupDlPathlosschangeEnumeratedDb3
	MacMainconfigPhrConfigSetupDlPathlosschangeEnumeratedDb6
	MacMainconfigPhrConfigSetupDlPathlosschangeEnumeratedInfinity
)
