// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NCR-SemiPersistentFwdResourceId-r18 (line 10436).
// NCR-SemiPersistentFwdResourceId-r18 ::= INTEGER (0..maxNrofSemiPersistentFwdResource-1-r18)

var nCRSemiPersistentFwdResourceIdR18Constraints = per.Constrained(0, common.MaxNrofSemiPersistentFwdResource_1_r18)

type NCR_SemiPersistentFwdResourceId_r18 struct {
	Value int64
}

func (v *NCR_SemiPersistentFwdResourceId_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, nCRSemiPersistentFwdResourceIdR18Constraints)
}

func (v *NCR_SemiPersistentFwdResourceId_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(nCRSemiPersistentFwdResourceIdR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
