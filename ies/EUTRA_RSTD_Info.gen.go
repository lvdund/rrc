// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-RSTD-Info (line 8569).

var eUTRARSTDInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq"},
		{Name: "measPRS-Offset"},
	},
}

var eUTRARSTDInfoMeasPRSOffsetConstraints = per.Constrained(0, 39)

type EUTRA_RSTD_Info struct {
	CarrierFreq    ARFCN_ValueEUTRA
	MeasPRS_Offset int64
}

func (ie *EUTRA_RSTD_Info) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eUTRARSTDInfoConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. carrierFreq: ref
	{
		if err := ie.CarrierFreq.Encode(e); err != nil {
			return err
		}
	}

	// 3. measPRS-Offset: integer
	{
		if err := e.EncodeInteger(ie.MeasPRS_Offset, eUTRARSTDInfoMeasPRSOffsetConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *EUTRA_RSTD_Info) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eUTRARSTDInfoConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. carrierFreq: ref
	{
		if err := ie.CarrierFreq.Decode(d); err != nil {
			return err
		}
	}

	// 3. measPRS-Offset: integer
	{
		v1, err := d.DecodeInteger(eUTRARSTDInfoMeasPRSOffsetConstraints)
		if err != nil {
			return err
		}
		ie.MeasPRS_Offset = v1
	}

	return nil
}
