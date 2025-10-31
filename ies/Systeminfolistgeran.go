package ies

import "rrc/utils"

// SystemInfoListGERAN ::= SEQUENCE OF utils.OCTETSTRING // SIZE (1..maxGERAN-SI)
type Systeminfolistgeran struct {
	Value []utils.OCTETSTRING `lb:1,ub:maxGERANSi`
}
