// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CLI-PeriodicalReportConfig-r16 (line 13891).

var cLIPeriodicalReportConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "reportInterval-r16"},
		{Name: "reportAmount-r16"},
		{Name: "reportQuantityCLI-r16"},
		{Name: "maxReportCLI-r16"},
	},
}

const (
	CLI_PeriodicalReportConfig_r16_ReportAmount_r16_r1       = 0
	CLI_PeriodicalReportConfig_r16_ReportAmount_r16_r2       = 1
	CLI_PeriodicalReportConfig_r16_ReportAmount_r16_r4       = 2
	CLI_PeriodicalReportConfig_r16_ReportAmount_r16_r8       = 3
	CLI_PeriodicalReportConfig_r16_ReportAmount_r16_r16      = 4
	CLI_PeriodicalReportConfig_r16_ReportAmount_r16_r32      = 5
	CLI_PeriodicalReportConfig_r16_ReportAmount_r16_r64      = 6
	CLI_PeriodicalReportConfig_r16_ReportAmount_r16_Infinity = 7
)

var cLIPeriodicalReportConfigR16ReportAmountR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CLI_PeriodicalReportConfig_r16_ReportAmount_r16_r1, CLI_PeriodicalReportConfig_r16_ReportAmount_r16_r2, CLI_PeriodicalReportConfig_r16_ReportAmount_r16_r4, CLI_PeriodicalReportConfig_r16_ReportAmount_r16_r8, CLI_PeriodicalReportConfig_r16_ReportAmount_r16_r16, CLI_PeriodicalReportConfig_r16_ReportAmount_r16_r32, CLI_PeriodicalReportConfig_r16_ReportAmount_r16_r64, CLI_PeriodicalReportConfig_r16_ReportAmount_r16_Infinity},
}

var cLIPeriodicalReportConfigR16MaxReportCLIR16Constraints = per.Constrained(1, common.MaxCLI_Report_r16)

type CLI_PeriodicalReportConfig_r16 struct {
	ReportInterval_r16    ReportInterval
	ReportAmount_r16      int64
	ReportQuantityCLI_r16 MeasReportQuantityCLI_r16
	MaxReportCLI_r16      int64
}

func (ie *CLI_PeriodicalReportConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cLIPeriodicalReportConfigR16Constraints)

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
		if err := e.EncodeEnumerated(ie.ReportAmount_r16, cLIPeriodicalReportConfigR16ReportAmountR16Constraints); err != nil {
			return err
		}
	}

	// 4. reportQuantityCLI-r16: ref
	{
		if err := ie.ReportQuantityCLI_r16.Encode(e); err != nil {
			return err
		}
	}

	// 5. maxReportCLI-r16: integer
	{
		if err := e.EncodeInteger(ie.MaxReportCLI_r16, cLIPeriodicalReportConfigR16MaxReportCLIR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CLI_PeriodicalReportConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cLIPeriodicalReportConfigR16Constraints)

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
		v1, err := d.DecodeEnumerated(cLIPeriodicalReportConfigR16ReportAmountR16Constraints)
		if err != nil {
			return err
		}
		ie.ReportAmount_r16 = v1
	}

	// 4. reportQuantityCLI-r16: ref
	{
		if err := ie.ReportQuantityCLI_r16.Decode(d); err != nil {
			return err
		}
	}

	// 5. maxReportCLI-r16: integer
	{
		v3, err := d.DecodeInteger(cLIPeriodicalReportConfigR16MaxReportCLIR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxReportCLI_r16 = v3
	}

	return nil
}
