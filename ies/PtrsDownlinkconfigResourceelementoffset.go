package ies

import "rrc/utils"

// PTRS-DownlinkConfig-resourceElementOffset ::= ENUMERATED
type PtrsDownlinkconfigResourceelementoffset struct {
	Value utils.ENUMERATED
}

const (
	PtrsDownlinkconfigResourceelementoffsetEnumeratedNothing = iota
	PtrsDownlinkconfigResourceelementoffsetEnumeratedOffset01
	PtrsDownlinkconfigResourceelementoffsetEnumeratedOffset10
	PtrsDownlinkconfigResourceelementoffsetEnumeratedOffset11
)
