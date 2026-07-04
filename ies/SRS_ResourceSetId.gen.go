// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-ResourceSetId (line 15430).
// SRS-ResourceSetId ::=                   INTEGER (0..maxNrofSRS-ResourceSets-1)

var sRSResourceSetIdConstraints = per.Constrained(0, common.MaxNrofSRS_ResourceSets_1)

type SRS_ResourceSetId struct {
	Value int64
}

func (v *SRS_ResourceSetId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sRSResourceSetIdConstraints)
}

func (v *SRS_ResourceSetId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sRSResourceSetIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
