// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NCR-PeriodicFwdResourceId-r18 (line 10404).
// NCR-PeriodicFwdResourceId-r18 ::= INTEGER (0..maxNrofPeriodicFwdResource-1-r18)

var nCRPeriodicFwdResourceIdR18Constraints = per.Constrained(0, common.MaxNrofPeriodicFwdResource_1_r18)

type NCR_PeriodicFwdResourceId_r18 struct {
	Value int64
}

func (v *NCR_PeriodicFwdResourceId_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, nCRPeriodicFwdResourceIdR18Constraints)
}

func (v *NCR_PeriodicFwdResourceId_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(nCRPeriodicFwdResourceIdR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
