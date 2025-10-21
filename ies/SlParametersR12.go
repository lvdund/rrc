package ies

import "rrc/utils"

// SL-Parameters-r12 ::= SEQUENCE
type SlParametersR12 struct {
	CommsimultaneoustxR12          *utils.ENUMERATED
	CommsupportedbandsR12          *FreqbandindicatorlisteutraR12
	DiscsupportedbandsR12          *SupportedbandinfolistR12
	DiscscheduledresourceallocR12  *utils.ENUMERATED
	DiscUeSelectedresourceallocR12 *utils.ENUMERATED
	DiscSlssR12                    *utils.ENUMERATED
	DiscsupportedprocR12           *utils.ENUMERATED
}
