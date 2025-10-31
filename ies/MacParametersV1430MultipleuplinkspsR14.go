package ies

import "rrc/utils"

// MAC-Parameters-v1430-multipleUplinkSPS-r14 ::= ENUMERATED
type MacParametersV1430MultipleuplinkspsR14 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1430MultipleuplinkspsR14EnumeratedNothing = iota
	MacParametersV1430MultipleuplinkspsR14EnumeratedSupported
)
