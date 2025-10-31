package ies

import "rrc/utils"

// PagingRecord-cn-Domain ::= ENUMERATED
type PagingrecordCnDomain struct {
	Value utils.ENUMERATED
}

const (
	PagingrecordCnDomainEnumeratedNothing = iota
	PagingrecordCnDomainEnumeratedPs
	PagingrecordCnDomainEnumeratedCs
)
