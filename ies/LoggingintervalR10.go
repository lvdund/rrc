package ies

import "rrc/utils"

// LoggingInterval-r10 ::= ENUMERATED
type LoggingintervalR10 struct {
	Value utils.ENUMERATED
}

const (
	LoggingintervalR10Ms1280  = 0
	LoggingintervalR10Ms2560  = 1
	LoggingintervalR10Ms5120  = 2
	LoggingintervalR10Ms10240 = 3
	LoggingintervalR10Ms20480 = 4
	LoggingintervalR10Ms30720 = 5
	LoggingintervalR10Ms40960 = 6
	LoggingintervalR10Ms61440 = 7
)
