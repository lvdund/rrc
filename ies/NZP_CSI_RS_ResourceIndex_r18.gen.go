// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NZP-CSI-RS-ResourceIndex-r18 (line 7328).
// NZP-CSI-RS-ResourceIndex-r18 ::=    INTEGER (0..maxNrofNZP-CSI-RS-ResourcesPerSet-1-r18)

var nZPCSIRSResourceIndexR18Constraints = per.Constrained(0, common.MaxNrofNZP_CSI_RS_ResourcesPerSet_1_r18)

type NZP_CSI_RS_ResourceIndex_r18 struct {
	Value int64
}

func (v *NZP_CSI_RS_ResourceIndex_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, nZPCSIRSResourceIndexR18Constraints)
}

func (v *NZP_CSI_RS_ResourceIndex_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(nZPCSIRSResourceIndexR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
