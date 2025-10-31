package ies

import "rrc/utils"

// IntraBandContiguousCC-Info-r12-supportedCSI-Proc-r12 ::= ENUMERATED
type IntrabandcontiguousccInfoR12SupportedcsiProcR12 struct {
	Value utils.ENUMERATED
}

const (
	IntrabandcontiguousccInfoR12SupportedcsiProcR12EnumeratedNothing = iota
	IntrabandcontiguousccInfoR12SupportedcsiProcR12EnumeratedN1
	IntrabandcontiguousccInfoR12SupportedcsiProcR12EnumeratedN3
	IntrabandcontiguousccInfoR12SupportedcsiProcR12EnumeratedN4
)
