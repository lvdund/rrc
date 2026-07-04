// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasReportQuantityCLI-r16 (line 13912).
// MeasReportQuantityCLI-r16 ::=               ENUMERATED {srs-rsrp, cli-rssi}

const (
	MeasReportQuantityCLI_r16_Srs_Rsrp = 0
	MeasReportQuantityCLI_r16_Cli_Rssi = 1
)

var measReportQuantityCLIR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasReportQuantityCLI_r16_Srs_Rsrp, MeasReportQuantityCLI_r16_Cli_Rssi},
}

type MeasReportQuantityCLI_r16 struct {
	Value int64
}

func (v *MeasReportQuantityCLI_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, measReportQuantityCLIR16Constraints)
}

func (v *MeasReportQuantityCLI_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(measReportQuantityCLIR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
