package ies

import "rrc/utils"

// T-StatusProhibit-v1610 ::= ENUMERATED
type TStatusprohibitV1610 struct {
	Value utils.ENUMERATED
}

const (
	TStatusprohibitV1610EnumeratedNothing = iota
	TStatusprohibitV1610EnumeratedMs1
	TStatusprohibitV1610EnumeratedMs2
	TStatusprohibitV1610EnumeratedMs3
	TStatusprohibitV1610EnumeratedMs4
	TStatusprohibitV1610EnumeratedSpare4
	TStatusprohibitV1610EnumeratedSpare3
	TStatusprohibitV1610EnumeratedSpare2
	TStatusprohibitV1610EnumeratedSpare1
)
