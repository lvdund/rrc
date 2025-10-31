package ies

import "rrc/utils"

// SDT-CG-Config-r17 ::= OCTET STRING (CONTAINING SDT-MAC-PHY-CG-Config-r17)
type SdtCgConfigR17 struct {
	Value utils.OCTETSTRING
}
