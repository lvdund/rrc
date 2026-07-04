// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: GNSS-ID-r16 (line 4916).

var gNSSIDR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "gnss-id-r16"},
	},
}

const (
	GNSS_ID_r16_Gnss_Id_r16_Gps     = 0
	GNSS_ID_r16_Gnss_Id_r16_Sbas    = 1
	GNSS_ID_r16_Gnss_Id_r16_Qzss    = 2
	GNSS_ID_r16_Gnss_Id_r16_Galileo = 3
	GNSS_ID_r16_Gnss_Id_r16_Glonass = 4
	GNSS_ID_r16_Gnss_Id_r16_Bds     = 5
)

var gNSSIDR16GnssIdR16Constraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{GNSS_ID_r16_Gnss_Id_r16_Gps, GNSS_ID_r16_Gnss_Id_r16_Sbas, GNSS_ID_r16_Gnss_Id_r16_Qzss, GNSS_ID_r16_Gnss_Id_r16_Galileo, GNSS_ID_r16_Gnss_Id_r16_Glonass, GNSS_ID_r16_Gnss_Id_r16_Bds},
	ExtValues:  []int64{6},
}

type GNSS_ID_r16 struct {
	Gnss_Id_r16 int64
}

func (ie *GNSS_ID_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(gNSSIDR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. gnss-id-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Gnss_Id_r16, gNSSIDR16GnssIdR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *GNSS_ID_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(gNSSIDR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. gnss-id-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(gNSSIDR16GnssIdR16Constraints)
		if err != nil {
			return err
		}
		ie.Gnss_Id_r16 = v0
	}

	return nil
}
