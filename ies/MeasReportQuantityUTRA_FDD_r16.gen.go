// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasReportQuantityUTRA-FDD-r16 (line 13514).

var measReportQuantityUTRAFDDR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cpich-RSCP"},
		{Name: "cpich-EcN0"},
	},
}

type MeasReportQuantityUTRA_FDD_r16 struct {
	Cpich_RSCP bool
	Cpich_EcN0 bool
}

func (ie *MeasReportQuantityUTRA_FDD_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measReportQuantityUTRAFDDR16Constraints)
	_ = seq

	// 1. cpich-RSCP: boolean
	{
		if err := e.EncodeBoolean(ie.Cpich_RSCP); err != nil {
			return err
		}
	}

	// 2. cpich-EcN0: boolean
	{
		if err := e.EncodeBoolean(ie.Cpich_EcN0); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasReportQuantityUTRA_FDD_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measReportQuantityUTRAFDDR16Constraints)
	_ = seq

	// 1. cpich-RSCP: boolean
	{
		v0, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.Cpich_RSCP = v0
	}

	// 2. cpich-EcN0: boolean
	{
		v1, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.Cpich_EcN0 = v1
	}

	return nil
}
