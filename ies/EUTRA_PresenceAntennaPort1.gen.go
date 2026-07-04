// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-PresenceAntennaPort1 (line 26202).
// EUTRA-PresenceAntennaPort1 ::=              BOOLEAN

type EUTRA_PresenceAntennaPort1 struct {
	Value bool
}

func (v *EUTRA_PresenceAntennaPort1) Encode(e *per.Encoder) error {
	return e.EncodeBoolean(v.Value)
}

func (v *EUTRA_PresenceAntennaPort1) Decode(d *per.Decoder) error {
	x, err := d.DecodeBoolean()
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
