// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultCLI-r16 (line 9853).

var measResultCLIR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measResultListSRS-RSRP-r16", Optional: true},
		{Name: "measResultListCLI-RSSI-r16", Optional: true},
	},
}

type MeasResultCLI_r16 struct {
	MeasResultListSRS_RSRP_r16 *MeasResultListSRS_RSRP_r16
	MeasResultListCLI_RSSI_r16 *MeasResultListCLI_RSSI_r16
}

func (ie *MeasResultCLI_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultCLIR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasResultListSRS_RSRP_r16 != nil, ie.MeasResultListCLI_RSSI_r16 != nil}); err != nil {
		return err
	}

	// 2. measResultListSRS-RSRP-r16: ref
	{
		if ie.MeasResultListSRS_RSRP_r16 != nil {
			if err := ie.MeasResultListSRS_RSRP_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. measResultListCLI-RSSI-r16: ref
	{
		if ie.MeasResultListCLI_RSSI_r16 != nil {
			if err := ie.MeasResultListCLI_RSSI_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasResultCLI_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultCLIR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measResultListSRS-RSRP-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasResultListSRS_RSRP_r16 = new(MeasResultListSRS_RSRP_r16)
			if err := ie.MeasResultListSRS_RSRP_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. measResultListCLI-RSSI-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MeasResultListCLI_RSSI_r16 = new(MeasResultListCLI_RSSI_r16)
			if err := ie.MeasResultListCLI_RSSI_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
