package ies

import "rrc/utils"

// ReleasePreferenceConfig-r16 ::= SEQUENCE
type ReleasepreferenceconfigR16 struct {
	ReleasepreferenceprohibittimerR16 ReleasepreferenceconfigR16ReleasepreferenceprohibittimerR16
	Connectedreporting                *utils.BOOLEAN
}
