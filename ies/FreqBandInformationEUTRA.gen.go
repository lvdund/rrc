// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FreqBandInformationEUTRA (line 20737).

var freqBandInformationEUTRAConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandEUTRA"},
		{Name: "ca-BandwidthClassDL-EUTRA", Optional: true},
		{Name: "ca-BandwidthClassUL-EUTRA", Optional: true},
	},
}

type FreqBandInformationEUTRA struct {
	BandEUTRA                 FreqBandIndicatorEUTRA
	Ca_BandwidthClassDL_EUTRA *CA_BandwidthClassEUTRA
	Ca_BandwidthClassUL_EUTRA *CA_BandwidthClassEUTRA
}

func (ie *FreqBandInformationEUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(freqBandInformationEUTRAConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_BandwidthClassDL_EUTRA != nil, ie.Ca_BandwidthClassUL_EUTRA != nil}); err != nil {
		return err
	}

	// 2. bandEUTRA: ref
	{
		if err := ie.BandEUTRA.Encode(e); err != nil {
			return err
		}
	}

	// 3. ca-BandwidthClassDL-EUTRA: ref
	{
		if ie.Ca_BandwidthClassDL_EUTRA != nil {
			if err := ie.Ca_BandwidthClassDL_EUTRA.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ca-BandwidthClassUL-EUTRA: ref
	{
		if ie.Ca_BandwidthClassUL_EUTRA != nil {
			if err := ie.Ca_BandwidthClassUL_EUTRA.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FreqBandInformationEUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(freqBandInformationEUTRAConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandEUTRA: ref
	{
		if err := ie.BandEUTRA.Decode(d); err != nil {
			return err
		}
	}

	// 3. ca-BandwidthClassDL-EUTRA: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_BandwidthClassDL_EUTRA = new(CA_BandwidthClassEUTRA)
			if err := ie.Ca_BandwidthClassDL_EUTRA.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ca-BandwidthClassUL-EUTRA: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ca_BandwidthClassUL_EUTRA = new(CA_BandwidthClassEUTRA)
			if err := ie.Ca_BandwidthClassUL_EUTRA.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
