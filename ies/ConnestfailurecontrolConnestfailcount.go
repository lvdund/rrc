package ies

import "rrc/utils"

// ConnEstFailureControl-connEstFailCount ::= ENUMERATED
type ConnestfailurecontrolConnestfailcount struct {
	Value utils.ENUMERATED
}

const (
	ConnestfailurecontrolConnestfailcountEnumeratedNothing = iota
	ConnestfailurecontrolConnestfailcountEnumeratedN1
	ConnestfailurecontrolConnestfailcountEnumeratedN2
	ConnestfailurecontrolConnestfailcountEnumeratedN3
	ConnestfailurecontrolConnestfailcountEnumeratedN4
)
