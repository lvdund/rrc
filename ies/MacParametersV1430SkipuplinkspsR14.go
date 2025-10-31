package ies

import "rrc/utils"

// MAC-Parameters-v1430-skipUplinkSPS-r14 ::= ENUMERATED
type MacParametersV1430SkipuplinkspsR14 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1430SkipuplinkspsR14EnumeratedNothing = iota
	MacParametersV1430SkipuplinkspsR14EnumeratedSupported
)
