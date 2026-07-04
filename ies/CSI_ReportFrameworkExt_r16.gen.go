// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-ReportFrameworkExt-r16 (line 22494).

var cSIReportFrameworkExtR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberAperiodicCSI-PerBWP-ForCSI-ReportExt-r16"},
	},
}

var cSIReportFrameworkExtR16MaxNumberAperiodicCSIPerBWPForCSIReportExtR16Constraints = per.Constrained(5, 8)

type CSI_ReportFrameworkExt_r16 struct {
	MaxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16 int64
}

func (ie *CSI_ReportFrameworkExt_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIReportFrameworkExtR16Constraints)
	_ = seq

	// 1. maxNumberAperiodicCSI-PerBWP-ForCSI-ReportExt-r16: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16, cSIReportFrameworkExtR16MaxNumberAperiodicCSIPerBWPForCSIReportExtR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CSI_ReportFrameworkExt_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIReportFrameworkExtR16Constraints)
	_ = seq

	// 1. maxNumberAperiodicCSI-PerBWP-ForCSI-ReportExt-r16: integer
	{
		v0, err := d.DecodeInteger(cSIReportFrameworkExtR16MaxNumberAperiodicCSIPerBWPForCSIReportExtR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16 = v0
	}

	return nil
}
