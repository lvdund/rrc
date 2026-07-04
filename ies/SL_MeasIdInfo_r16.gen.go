// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-MeasIdInfo-r16 (line 27540).

var sLMeasIdInfoR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-MeasId-r16"},
		{Name: "sl-MeasObjectId-r16"},
		{Name: "sl-ReportConfigId-r16"},
	},
}

type SL_MeasIdInfo_r16 struct {
	Sl_MeasId_r16         SL_MeasId_r16
	Sl_MeasObjectId_r16   SL_MeasObjectId_r16
	Sl_ReportConfigId_r16 SL_ReportConfigId_r16
}

func (ie *SL_MeasIdInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLMeasIdInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. sl-MeasId-r16: ref
	{
		if err := ie.Sl_MeasId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-MeasObjectId-r16: ref
	{
		if err := ie.Sl_MeasObjectId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-ReportConfigId-r16: ref
	{
		if err := ie.Sl_ReportConfigId_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_MeasIdInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLMeasIdInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. sl-MeasId-r16: ref
	{
		if err := ie.Sl_MeasId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-MeasObjectId-r16: ref
	{
		if err := ie.Sl_MeasObjectId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-ReportConfigId-r16: ref
	{
		if err := ie.Sl_ReportConfigId_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
