// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NCR-AperiodicFwdTimeResourceId-r18 (line 10299).
// NCR-AperiodicFwdTimeResourceId-r18 ::= INTEGER (0..maxNrofAperiodicFwdTimeResource-1-r18)

var nCRAperiodicFwdTimeResourceIdR18Constraints = per.Constrained(0, common.MaxNrofAperiodicFwdTimeResource_1_r18)

type NCR_AperiodicFwdTimeResourceId_r18 struct {
	Value int64
}

func (v *NCR_AperiodicFwdTimeResourceId_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, nCRAperiodicFwdTimeResourceIdR18Constraints)
}

func (v *NCR_AperiodicFwdTimeResourceId_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(nCRAperiodicFwdTimeResourceIdR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
