package ies

import "rrc/utils"

// SystemInfoListGERAN ::= SEQUENCE OF OCTET
// SIZE (1..maxGERAN-SI)
type Systeminfolistgeran struct {
	Value []Octet
}
