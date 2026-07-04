// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NCR-SemiPersistentFwdResourceSetId-r18 (line 10441).
// NCR-SemiPersistentFwdResourceSetId-r18 ::= INTEGER (0..maxNrofSemiPersistentFwdResourceSet-1-r18)

var nCRSemiPersistentFwdResourceSetIdR18Constraints = per.Constrained(0, common.MaxNrofSemiPersistentFwdResourceSet_1_r18)

type NCR_SemiPersistentFwdResourceSetId_r18 struct {
	Value int64
}

func (v *NCR_SemiPersistentFwdResourceSetId_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, nCRSemiPersistentFwdResourceSetIdR18Constraints)
}

func (v *NCR_SemiPersistentFwdResourceSetId_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(nCRSemiPersistentFwdResourceSetIdR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
