// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PerRACSI-RSInfo-r16 (line 3182).

var perRACSIRSInfoR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-RS-Index-r16"},
		{Name: "numberOfPreamblesSentOnCSI-RS-r16"},
	},
}

var perRACSIRSInfoR16NumberOfPreamblesSentOnCSIRSR16Constraints = per.Constrained(1, 200)

type PerRACSI_RSInfo_r16 struct {
	Csi_RS_Index_r16                  CSI_RS_Index
	NumberOfPreamblesSentOnCSI_RS_r16 int64
}

func (ie *PerRACSI_RSInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(perRACSIRSInfoR16Constraints)
	_ = seq

	// 1. csi-RS-Index-r16: ref
	{
		if err := ie.Csi_RS_Index_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. numberOfPreamblesSentOnCSI-RS-r16: integer
	{
		if err := e.EncodeInteger(ie.NumberOfPreamblesSentOnCSI_RS_r16, perRACSIRSInfoR16NumberOfPreamblesSentOnCSIRSR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PerRACSI_RSInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(perRACSIRSInfoR16Constraints)
	_ = seq

	// 1. csi-RS-Index-r16: ref
	{
		if err := ie.Csi_RS_Index_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. numberOfPreamblesSentOnCSI-RS-r16: integer
	{
		v1, err := d.DecodeInteger(perRACSIRSInfoR16NumberOfPreamblesSentOnCSIRSR16Constraints)
		if err != nil {
			return err
		}
		ie.NumberOfPreamblesSentOnCSI_RS_r16 = v1
	}

	return nil
}
