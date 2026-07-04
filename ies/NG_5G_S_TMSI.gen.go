// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NG-5G-S-TMSI (line 10545).
// NG-5G-S-TMSI ::=                         BIT STRING (SIZE (48))

var nG5GSTMSIConstraints = per.FixedSize(48)

type NG_5G_S_TMSI struct {
	Value per.BitString
}

func (v *NG_5G_S_TMSI) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, nG5GSTMSIConstraints)
}

func (v *NG_5G_S_TMSI) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(nG5GSTMSIConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
