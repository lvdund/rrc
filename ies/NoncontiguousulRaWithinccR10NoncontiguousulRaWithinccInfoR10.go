package ies

import "rrc/utils"

// NonContiguousUL-RA-WithinCC-r10-nonContiguousUL-RA-WithinCC-Info-r10 ::= ENUMERATED
type NoncontiguousulRaWithinccR10NoncontiguousulRaWithinccInfoR10 struct {
	Value utils.ENUMERATED
}

const (
	NoncontiguousulRaWithinccR10NoncontiguousulRaWithinccInfoR10EnumeratedNothing = iota
	NoncontiguousulRaWithinccR10NoncontiguousulRaWithinccInfoR10EnumeratedSupported
)
