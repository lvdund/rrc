// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NCR-PeriodicFwdResourceSetId-r18 (line 10409).
// NCR-PeriodicFwdResourceSetId-r18 ::= INTEGER (0..maxNrofPeriodicFwdResourceSet-1-r18)

var nCRPeriodicFwdResourceSetIdR18Constraints = per.Constrained(0, common.MaxNrofPeriodicFwdResourceSet_1_r18)

type NCR_PeriodicFwdResourceSetId_r18 struct {
	Value int64
}

func (v *NCR_PeriodicFwdResourceSetId_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, nCRPeriodicFwdResourceSetIdR18Constraints)
}

func (v *NCR_PeriodicFwdResourceSetId_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(nCRPeriodicFwdResourceSetIdR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
