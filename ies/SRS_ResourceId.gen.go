// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-ResourceId (line 15652).
// SRS-ResourceId ::=                      INTEGER (0..maxNrofSRS-Resources-1)

var sRSResourceIdConstraints = per.Constrained(0, common.MaxNrofSRS_Resources_1)

type SRS_ResourceId struct {
	Value int64
}

func (v *SRS_ResourceId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sRSResourceIdConstraints)
}

func (v *SRS_ResourceId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sRSResourceIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
