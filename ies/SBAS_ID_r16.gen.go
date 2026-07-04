// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SBAS-ID-r16 (line 4921).

var sBASIDR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sbas-id-r16"},
	},
}

const (
	SBAS_ID_r16_Sbas_Id_r16_Waas  = 0
	SBAS_ID_r16_Sbas_Id_r16_Egnos = 1
	SBAS_ID_r16_Sbas_Id_r16_Msas  = 2
	SBAS_ID_r16_Sbas_Id_r16_Gagan = 3
)

var sBASIDR16SbasIdR16Constraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{SBAS_ID_r16_Sbas_Id_r16_Waas, SBAS_ID_r16_Sbas_Id_r16_Egnos, SBAS_ID_r16_Sbas_Id_r16_Msas, SBAS_ID_r16_Sbas_Id_r16_Gagan},
}

type SBAS_ID_r16 struct {
	Sbas_Id_r16 int64
}

func (ie *SBAS_ID_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sBASIDR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. sbas-id-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sbas_Id_r16, sBASIDR16SbasIdR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SBAS_ID_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sBASIDR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. sbas-id-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(sBASIDR16SbasIdR16Constraints)
		if err != nil {
			return err
		}
		ie.Sbas_Id_r16 = v0
	}

	return nil
}
