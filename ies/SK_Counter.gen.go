// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SK-Counter (line 15124).
// SK-Counter ::=  INTEGER (0..65535)

var sKCounterConstraints = per.Constrained(0, 65535)

type SK_Counter struct {
	Value int64
}

func (v *SK_Counter) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sKCounterConstraints)
}

func (v *SK_Counter) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sKCounterConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
