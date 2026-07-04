// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LoggedPeriodicalReportConfig-r16 (line 578).

var loggedPeriodicalReportConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "loggingInterval-r16"},
	},
}

type LoggedPeriodicalReportConfig_r16 struct {
	LoggingInterval_r16 LoggingInterval_r16
}

func (ie *LoggedPeriodicalReportConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(loggedPeriodicalReportConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. loggingInterval-r16: ref
	{
		if err := ie.LoggingInterval_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *LoggedPeriodicalReportConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(loggedPeriodicalReportConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. loggingInterval-r16: ref
	{
		if err := ie.LoggingInterval_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
