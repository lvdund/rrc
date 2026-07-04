// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ReportConfigInfo-r16 (line 27874).

var sLReportConfigInfoR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ReportConfigId-r16"},
		{Name: "sl-ReportConfig-r16"},
	},
}

type SL_ReportConfigInfo_r16 struct {
	Sl_ReportConfigId_r16 SL_ReportConfigId_r16
	Sl_ReportConfig_r16   SL_ReportConfig_r16
}

func (ie *SL_ReportConfigInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLReportConfigInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. sl-ReportConfigId-r16: ref
	{
		if err := ie.Sl_ReportConfigId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-ReportConfig-r16: ref
	{
		if err := ie.Sl_ReportConfig_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_ReportConfigInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLReportConfigInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. sl-ReportConfigId-r16: ref
	{
		if err := ie.Sl_ReportConfigId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-ReportConfig-r16: ref
	{
		if err := ie.Sl_ReportConfig_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
