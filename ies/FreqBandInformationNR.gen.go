// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FreqBandInformationNR (line 20743).

var freqBandInformationNRConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandNR"},
		{Name: "maxBandwidthRequestedDL", Optional: true},
		{Name: "maxBandwidthRequestedUL", Optional: true},
		{Name: "maxCarriersRequestedDL", Optional: true},
		{Name: "maxCarriersRequestedUL", Optional: true},
	},
}

var freqBandInformationNRMaxCarriersRequestedDLConstraints = per.Constrained(1, common.MaxNrofServingCells)

var freqBandInformationNRMaxCarriersRequestedULConstraints = per.Constrained(1, common.MaxNrofServingCells)

type FreqBandInformationNR struct {
	BandNR                  FreqBandIndicatorNR
	MaxBandwidthRequestedDL *AggregatedBandwidth
	MaxBandwidthRequestedUL *AggregatedBandwidth
	MaxCarriersRequestedDL  *int64
	MaxCarriersRequestedUL  *int64
}

func (ie *FreqBandInformationNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(freqBandInformationNRConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxBandwidthRequestedDL != nil, ie.MaxBandwidthRequestedUL != nil, ie.MaxCarriersRequestedDL != nil, ie.MaxCarriersRequestedUL != nil}); err != nil {
		return err
	}

	// 2. bandNR: ref
	{
		if err := ie.BandNR.Encode(e); err != nil {
			return err
		}
	}

	// 3. maxBandwidthRequestedDL: ref
	{
		if ie.MaxBandwidthRequestedDL != nil {
			if err := ie.MaxBandwidthRequestedDL.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. maxBandwidthRequestedUL: ref
	{
		if ie.MaxBandwidthRequestedUL != nil {
			if err := ie.MaxBandwidthRequestedUL.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. maxCarriersRequestedDL: integer
	{
		if ie.MaxCarriersRequestedDL != nil {
			if err := e.EncodeInteger(*ie.MaxCarriersRequestedDL, freqBandInformationNRMaxCarriersRequestedDLConstraints); err != nil {
				return err
			}
		}
	}

	// 6. maxCarriersRequestedUL: integer
	{
		if ie.MaxCarriersRequestedUL != nil {
			if err := e.EncodeInteger(*ie.MaxCarriersRequestedUL, freqBandInformationNRMaxCarriersRequestedULConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FreqBandInformationNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(freqBandInformationNRConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandNR: ref
	{
		if err := ie.BandNR.Decode(d); err != nil {
			return err
		}
	}

	// 3. maxBandwidthRequestedDL: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MaxBandwidthRequestedDL = new(AggregatedBandwidth)
			if err := ie.MaxBandwidthRequestedDL.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. maxBandwidthRequestedUL: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MaxBandwidthRequestedUL = new(AggregatedBandwidth)
			if err := ie.MaxBandwidthRequestedUL.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. maxCarriersRequestedDL: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(freqBandInformationNRMaxCarriersRequestedDLConstraints)
			if err != nil {
				return err
			}
			ie.MaxCarriersRequestedDL = &v
		}
	}

	// 6. maxCarriersRequestedUL: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(freqBandInformationNRMaxCarriersRequestedULConstraints)
			if err != nil {
				return err
			}
			ie.MaxCarriersRequestedUL = &v
		}
	}

	return nil
}
