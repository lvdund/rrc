// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-MRDC-CapabilityAddXDD-Mode (line 25612).

var uEMRDCCapabilityAddXDDModeConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersMRDC-XDD-Diff", Optional: true},
		{Name: "generalParametersMRDC-XDD-Diff", Optional: true},
	},
}

type UE_MRDC_CapabilityAddXDD_Mode struct {
	MeasAndMobParametersMRDC_XDD_Diff *MeasAndMobParametersMRDC_XDD_Diff
	GeneralParametersMRDC_XDD_Diff    *GeneralParametersMRDC_XDD_Diff
}

func (ie *UE_MRDC_CapabilityAddXDD_Mode) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEMRDCCapabilityAddXDDModeConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersMRDC_XDD_Diff != nil, ie.GeneralParametersMRDC_XDD_Diff != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-XDD-Diff: ref
	{
		if ie.MeasAndMobParametersMRDC_XDD_Diff != nil {
			if err := ie.MeasAndMobParametersMRDC_XDD_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. generalParametersMRDC-XDD-Diff: ref
	{
		if ie.GeneralParametersMRDC_XDD_Diff != nil {
			if err := ie.GeneralParametersMRDC_XDD_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_MRDC_CapabilityAddXDD_Mode) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEMRDCCapabilityAddXDDModeConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-XDD-Diff: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersMRDC_XDD_Diff = new(MeasAndMobParametersMRDC_XDD_Diff)
			if err := ie.MeasAndMobParametersMRDC_XDD_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. generalParametersMRDC-XDD-Diff: ref
	{
		if seq.IsComponentPresent(1) {
			ie.GeneralParametersMRDC_XDD_Diff = new(GeneralParametersMRDC_XDD_Diff)
			if err := ie.GeneralParametersMRDC_XDD_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
