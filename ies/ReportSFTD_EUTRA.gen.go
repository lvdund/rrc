// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReportSFTD-EUTRA (line 13393).

var reportSFTDEUTRAConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "reportSFTD-Meas"},
		{Name: "reportRSRP"},
	},
}

type ReportSFTD_EUTRA struct {
	ReportSFTD_Meas bool
	ReportRSRP      bool
}

func (ie *ReportSFTD_EUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(reportSFTDEUTRAConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. reportSFTD-Meas: boolean
	{
		if err := e.EncodeBoolean(ie.ReportSFTD_Meas); err != nil {
			return err
		}
	}

	// 3. reportRSRP: boolean
	{
		if err := e.EncodeBoolean(ie.ReportRSRP); err != nil {
			return err
		}
	}

	return nil
}

func (ie *ReportSFTD_EUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(reportSFTDEUTRAConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. reportSFTD-Meas: boolean
	{
		v0, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.ReportSFTD_Meas = v0
	}

	// 3. reportRSRP: boolean
	{
		v1, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.ReportRSRP = v1
	}

	return nil
}
