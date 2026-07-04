// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultNeighFreqRSSI-r18 (line 3399).

var measResultNeighFreqRSSIR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ssbFrequency-r18", Optional: true},
		{Name: "ssbSubcarrierSpacing-r18", Optional: true},
		{Name: "refFreqCSI-RS-r18", Optional: true},
		{Name: "measResult-RSSI-r18", Optional: true},
	},
}

type MeasResultNeighFreqRSSI_r18 struct {
	SsbFrequency_r18         *ARFCN_ValueNR
	SsbSubcarrierSpacing_r18 *SubcarrierSpacing
	RefFreqCSI_RS_r18        *ARFCN_ValueNR
	MeasResult_RSSI_r18      *RSSI_Range_r16
}

func (ie *MeasResultNeighFreqRSSI_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultNeighFreqRSSIR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SsbFrequency_r18 != nil, ie.SsbSubcarrierSpacing_r18 != nil, ie.RefFreqCSI_RS_r18 != nil, ie.MeasResult_RSSI_r18 != nil}); err != nil {
		return err
	}

	// 2. ssbFrequency-r18: ref
	{
		if ie.SsbFrequency_r18 != nil {
			if err := ie.SsbFrequency_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ssbSubcarrierSpacing-r18: ref
	{
		if ie.SsbSubcarrierSpacing_r18 != nil {
			if err := ie.SsbSubcarrierSpacing_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. refFreqCSI-RS-r18: ref
	{
		if ie.RefFreqCSI_RS_r18 != nil {
			if err := ie.RefFreqCSI_RS_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. measResult-RSSI-r18: ref
	{
		if ie.MeasResult_RSSI_r18 != nil {
			if err := ie.MeasResult_RSSI_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasResultNeighFreqRSSI_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultNeighFreqRSSIR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ssbFrequency-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SsbFrequency_r18 = new(ARFCN_ValueNR)
			if err := ie.SsbFrequency_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ssbSubcarrierSpacing-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SsbSubcarrierSpacing_r18 = new(SubcarrierSpacing)
			if err := ie.SsbSubcarrierSpacing_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. refFreqCSI-RS-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.RefFreqCSI_RS_r18 = new(ARFCN_ValueNR)
			if err := ie.RefFreqCSI_RS_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. measResult-RSSI-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.MeasResult_RSSI_r18 = new(RSSI_Range_r16)
			if err := ie.MeasResult_RSSI_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
