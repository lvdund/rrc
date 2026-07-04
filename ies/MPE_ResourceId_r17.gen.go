// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MPE-ResourceId-r17 (line 12509).
// MPE-ResourceId-r17 ::=      INTEGER (1..maxMPE-Resources-r17)

var mPEResourceIdR17Constraints = per.Constrained(1, common.MaxMPE_Resources_r17)

type MPE_ResourceId_r17 struct {
	Value int64
}

func (v *MPE_ResourceId_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, mPEResourceIdR17Constraints)
}

func (v *MPE_ResourceId_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(mPEResourceIdR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
