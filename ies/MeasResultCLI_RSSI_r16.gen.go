// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultCLI-RSSI-r16 (line 9867).

var measResultCLIRSSIR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rssi-ResourceId-r16"},
		{Name: "cli-RSSI-Result-r16"},
	},
}

type MeasResultCLI_RSSI_r16 struct {
	Rssi_ResourceId_r16 RSSI_ResourceId_r16
	Cli_RSSI_Result_r16 CLI_RSSI_Range_r16
}

func (ie *MeasResultCLI_RSSI_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultCLIRSSIR16Constraints)
	_ = seq

	// 1. rssi-ResourceId-r16: ref
	{
		if err := ie.Rssi_ResourceId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. cli-RSSI-Result-r16: ref
	{
		if err := ie.Cli_RSSI_Result_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasResultCLI_RSSI_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultCLIRSSIR16Constraints)
	_ = seq

	// 1. rssi-ResourceId-r16: ref
	{
		if err := ie.Rssi_ResourceId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. cli-RSSI-Result-r16: ref
	{
		if err := ie.Cli_RSSI_Result_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
