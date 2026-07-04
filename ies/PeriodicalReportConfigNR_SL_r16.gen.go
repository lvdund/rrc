// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PeriodicalReportConfigNR-SL-r16 (line 13957).

var periodicalReportConfigNRSLR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "reportInterval-r16"},
		{Name: "reportAmount-r16"},
		{Name: "reportQuantity-r16"},
	},
}

const (
	PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_r1       = 0
	PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_r2       = 1
	PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_r4       = 2
	PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_r8       = 3
	PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_r16      = 4
	PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_r32      = 5
	PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_r64      = 6
	PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_Infinity = 7
)

var periodicalReportConfigNRSLR16ReportAmountR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_r1, PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_r2, PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_r4, PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_r8, PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_r16, PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_r32, PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_r64, PeriodicalReportConfigNR_SL_r16_ReportAmount_r16_Infinity},
}

type PeriodicalReportConfigNR_SL_r16 struct {
	ReportInterval_r16 ReportInterval
	ReportAmount_r16   int64
	ReportQuantity_r16 MeasReportQuantity_r16
}

func (ie *PeriodicalReportConfigNR_SL_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(periodicalReportConfigNRSLR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. reportInterval-r16: ref
	{
		if err := ie.ReportInterval_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. reportAmount-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.ReportAmount_r16, periodicalReportConfigNRSLR16ReportAmountR16Constraints); err != nil {
			return err
		}
	}

	// 4. reportQuantity-r16: ref
	{
		if err := ie.ReportQuantity_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PeriodicalReportConfigNR_SL_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(periodicalReportConfigNRSLR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. reportInterval-r16: ref
	{
		if err := ie.ReportInterval_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. reportAmount-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(periodicalReportConfigNRSLR16ReportAmountR16Constraints)
		if err != nil {
			return err
		}
		ie.ReportAmount_r16 = v1
	}

	// 4. reportQuantity-r16: ref
	{
		if err := ie.ReportQuantity_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
