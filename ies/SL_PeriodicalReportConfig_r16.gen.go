// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PeriodicalReportConfig-r16 (line 27891).

var sLPeriodicalReportConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ReportInterval-r16"},
		{Name: "sl-ReportAmount-r16"},
		{Name: "sl-ReportQuantity-r16"},
		{Name: "sl-RS-Type-r16"},
	},
}

const (
	SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_r1       = 0
	SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_r2       = 1
	SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_r4       = 2
	SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_r8       = 3
	SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_r16      = 4
	SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_r32      = 5
	SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_r64      = 6
	SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_Infinity = 7
)

var sLPeriodicalReportConfigR16SlReportAmountR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_r1, SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_r2, SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_r4, SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_r8, SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_r16, SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_r32, SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_r64, SL_PeriodicalReportConfig_r16_Sl_ReportAmount_r16_Infinity},
}

type SL_PeriodicalReportConfig_r16 struct {
	Sl_ReportInterval_r16 ReportInterval
	Sl_ReportAmount_r16   int64
	Sl_ReportQuantity_r16 SL_MeasReportQuantity_r16
	Sl_RS_Type_r16        SL_RS_Type_r16
}

func (ie *SL_PeriodicalReportConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPeriodicalReportConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. sl-ReportInterval-r16: ref
	{
		if err := ie.Sl_ReportInterval_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-ReportAmount-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_ReportAmount_r16, sLPeriodicalReportConfigR16SlReportAmountR16Constraints); err != nil {
			return err
		}
	}

	// 4. sl-ReportQuantity-r16: ref
	{
		if err := ie.Sl_ReportQuantity_r16.Encode(e); err != nil {
			return err
		}
	}

	// 5. sl-RS-Type-r16: ref
	{
		if err := ie.Sl_RS_Type_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_PeriodicalReportConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPeriodicalReportConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. sl-ReportInterval-r16: ref
	{
		if err := ie.Sl_ReportInterval_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-ReportAmount-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(sLPeriodicalReportConfigR16SlReportAmountR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_ReportAmount_r16 = v1
	}

	// 4. sl-ReportQuantity-r16: ref
	{
		if err := ie.Sl_ReportQuantity_r16.Decode(d); err != nil {
			return err
		}
	}

	// 5. sl-RS-Type-r16: ref
	{
		if err := ie.Sl_RS_Type_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
