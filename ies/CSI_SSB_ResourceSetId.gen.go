// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-SSB-ResourceSetId (line 7628).
// CSI-SSB-ResourceSetId ::=           INTEGER (0..maxNrofCSI-SSB-ResourceSets-1)

var cSISSBResourceSetIdConstraints = per.Constrained(0, common.MaxNrofCSI_SSB_ResourceSets_1)

type CSI_SSB_ResourceSetId struct {
	Value int64
}

func (v *CSI_SSB_ResourceSetId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, cSISSBResourceSetIdConstraints)
}

func (v *CSI_SSB_ResourceSetId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(cSISSBResourceSetIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
