// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MBS-FSAI-r17 (line 4607).
// MBS-FSAI-r17 ::= OCTET STRING (SIZE (3))

var mBSFSAIR17Constraints = per.FixedSize(3)

type MBS_FSAI_r17 struct {
	Value []byte
}

func (v *MBS_FSAI_r17) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(v.Value, mBSFSAIR17Constraints)
}

func (v *MBS_FSAI_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeOctetString(mBSFSAIR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
