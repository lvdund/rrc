// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LoggingDuration-r16 (line 26237).

const (
	LoggingDuration_r16_Min10  = 0
	LoggingDuration_r16_Min20  = 1
	LoggingDuration_r16_Min40  = 2
	LoggingDuration_r16_Min60  = 3
	LoggingDuration_r16_Min90  = 4
	LoggingDuration_r16_Min120 = 5
	LoggingDuration_r16_Spare2 = 6
	LoggingDuration_r16_Spare1 = 7
)

var loggingDurationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LoggingDuration_r16_Min10, LoggingDuration_r16_Min20, LoggingDuration_r16_Min40, LoggingDuration_r16_Min60, LoggingDuration_r16_Min90, LoggingDuration_r16_Min120, LoggingDuration_r16_Spare2, LoggingDuration_r16_Spare1},
}

type LoggingDuration_r16 struct {
	Value int64
}

func (v *LoggingDuration_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, loggingDurationR16Constraints)
}

func (v *LoggingDuration_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(loggingDurationR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
