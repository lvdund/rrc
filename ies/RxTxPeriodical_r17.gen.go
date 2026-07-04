// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RxTxPeriodical-r17 (line 13899).

var rxTxPeriodicalR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rxTxReportInterval-r17", Optional: true},
		{Name: "reportAmount-r17"},
	},
}

const (
	RxTxPeriodical_r17_ReportAmount_r17_r1       = 0
	RxTxPeriodical_r17_ReportAmount_r17_Infinity = 1
	RxTxPeriodical_r17_ReportAmount_r17_Spare6   = 2
	RxTxPeriodical_r17_ReportAmount_r17_Spare5   = 3
	RxTxPeriodical_r17_ReportAmount_r17_Spare4   = 4
	RxTxPeriodical_r17_ReportAmount_r17_Spare3   = 5
	RxTxPeriodical_r17_ReportAmount_r17_Spare2   = 6
	RxTxPeriodical_r17_ReportAmount_r17_Spare1   = 7
)

var rxTxPeriodicalR17ReportAmountR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RxTxPeriodical_r17_ReportAmount_r17_r1, RxTxPeriodical_r17_ReportAmount_r17_Infinity, RxTxPeriodical_r17_ReportAmount_r17_Spare6, RxTxPeriodical_r17_ReportAmount_r17_Spare5, RxTxPeriodical_r17_ReportAmount_r17_Spare4, RxTxPeriodical_r17_ReportAmount_r17_Spare3, RxTxPeriodical_r17_ReportAmount_r17_Spare2, RxTxPeriodical_r17_ReportAmount_r17_Spare1},
}

type RxTxPeriodical_r17 struct {
	RxTxReportInterval_r17 *RxTxReportInterval_r17
	ReportAmount_r17       int64
}

func (ie *RxTxPeriodical_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rxTxPeriodicalR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RxTxReportInterval_r17 != nil}); err != nil {
		return err
	}

	// 3. rxTxReportInterval-r17: ref
	{
		if ie.RxTxReportInterval_r17 != nil {
			if err := ie.RxTxReportInterval_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. reportAmount-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.ReportAmount_r17, rxTxPeriodicalR17ReportAmountR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RxTxPeriodical_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rxTxPeriodicalR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rxTxReportInterval-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.RxTxReportInterval_r17 = new(RxTxReportInterval_r17)
			if err := ie.RxTxReportInterval_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. reportAmount-r17: enumerated
	{
		v1, err := d.DecodeEnumerated(rxTxPeriodicalR17ReportAmountR17Constraints)
		if err != nil {
			return err
		}
		ie.ReportAmount_r17 = v1
	}

	return nil
}
