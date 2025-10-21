package ies

import "rrc/utils"

// MBSFNAreaConfiguration-r9 ::= SEQUENCE
type MbsfnareaconfigurationR9 struct {
	CommonsfAllocR9       CommonsfAllocpatternlistR9
	CommonsfAllocperiodR9 utils.ENUMERATED
	PmchInfolistR9        PmchInfolistR9
	Noncriticalextension  *MbsfnareaconfigurationV930Ies
}
