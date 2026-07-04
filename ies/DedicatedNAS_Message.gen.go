// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DedicatedNAS-Message (line 7638).
// DedicatedNAS-Message ::=        OCTET STRING

var dedicatedNASMessageConstraints = per.SizeConstraints{}

type DedicatedNAS_Message struct {
	Value []byte
}

func (v *DedicatedNAS_Message) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(v.Value, dedicatedNASMessageConstraints)
}

func (v *DedicatedNAS_Message) Decode(d *per.Decoder) error {
	x, err := d.DecodeOctetString(dedicatedNASMessageConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
