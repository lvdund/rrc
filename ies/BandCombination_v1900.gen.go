// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-v1900 (line 16816).

var bandCombinationV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-v1900", Optional: true},
		{Name: "ca-ParametersNRDC-v1900", Optional: true},
		{Name: "mrdc-Parameters-v1900", Optional: true},
		{Name: "bandList-v1900", Optional: true},
		{Name: "featureSetCombinationLowBandSwitching-r19", Optional: true},
		{Name: "switchingPeriodForFDD-SDL-r19", Optional: true},
		{Name: "directSCellActivation-r19", Optional: true},
	},
}

var bandCombinationV1900BandListV1900Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

const (
	BandCombination_v1900_SwitchingPeriodForFDD_SDL_r19_N35us  = 0
	BandCombination_v1900_SwitchingPeriodForFDD_SDL_r19_N70us  = 1
	BandCombination_v1900_SwitchingPeriodForFDD_SDL_r19_N140us = 2
)

var bandCombinationV1900SwitchingPeriodForFDDSDLR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandCombination_v1900_SwitchingPeriodForFDD_SDL_r19_N35us, BandCombination_v1900_SwitchingPeriodForFDD_SDL_r19_N70us, BandCombination_v1900_SwitchingPeriodForFDD_SDL_r19_N140us},
}

const (
	BandCombination_v1900_DirectSCellActivation_r19_Supported = 0
)

var bandCombinationV1900DirectSCellActivationR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandCombination_v1900_DirectSCellActivation_r19_Supported},
}

type BandCombination_v1900 struct {
	Ca_ParametersNR_v1900                     *CA_ParametersNR_v1900
	Ca_ParametersNRDC_v1900                   *CA_ParametersNRDC_v1900
	Mrdc_Parameters_v1900                     *MRDC_Parameters_v1900
	BandList_v1900                            []BandParameters_v1900
	FeatureSetCombinationLowBandSwitching_r19 *FeatureSetCombinationId
	SwitchingPeriodForFDD_SDL_r19             *int64
	DirectSCellActivation_r19                 *int64
}

func (ie *BandCombination_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_v1900 != nil, ie.Ca_ParametersNRDC_v1900 != nil, ie.Mrdc_Parameters_v1900 != nil, ie.BandList_v1900 != nil, ie.FeatureSetCombinationLowBandSwitching_r19 != nil, ie.SwitchingPeriodForFDD_SDL_r19 != nil, ie.DirectSCellActivation_r19 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1900: ref
	{
		if ie.Ca_ParametersNR_v1900 != nil {
			if err := ie.Ca_ParametersNR_v1900.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v1900: ref
	{
		if ie.Ca_ParametersNRDC_v1900 != nil {
			if err := ie.Ca_ParametersNRDC_v1900.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. mrdc-Parameters-v1900: ref
	{
		if ie.Mrdc_Parameters_v1900 != nil {
			if err := ie.Mrdc_Parameters_v1900.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. bandList-v1900: sequence-of
	{
		if ie.BandList_v1900 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationV1900BandListV1900Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.BandList_v1900))); err != nil {
				return err
			}
			for i := range ie.BandList_v1900 {
				if err := ie.BandList_v1900[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. featureSetCombinationLowBandSwitching-r19: ref
	{
		if ie.FeatureSetCombinationLowBandSwitching_r19 != nil {
			if err := ie.FeatureSetCombinationLowBandSwitching_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. switchingPeriodForFDD-SDL-r19: enumerated
	{
		if ie.SwitchingPeriodForFDD_SDL_r19 != nil {
			if err := e.EncodeEnumerated(*ie.SwitchingPeriodForFDD_SDL_r19, bandCombinationV1900SwitchingPeriodForFDDSDLR19Constraints); err != nil {
				return err
			}
		}
	}

	// 8. directSCellActivation-r19: enumerated
	{
		if ie.DirectSCellActivation_r19 != nil {
			if err := e.EncodeEnumerated(*ie.DirectSCellActivation_r19, bandCombinationV1900DirectSCellActivationR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1900: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_v1900 = new(CA_ParametersNR_v1900)
			if err := ie.Ca_ParametersNR_v1900.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v1900: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_ParametersNRDC_v1900 = new(CA_ParametersNRDC_v1900)
			if err := ie.Ca_ParametersNRDC_v1900.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. mrdc-Parameters-v1900: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Mrdc_Parameters_v1900 = new(MRDC_Parameters_v1900)
			if err := ie.Mrdc_Parameters_v1900.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. bandList-v1900: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationV1900BandListV1900Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.BandList_v1900 = make([]BandParameters_v1900, n)
			for i := int64(0); i < n; i++ {
				if err := ie.BandList_v1900[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. featureSetCombinationLowBandSwitching-r19: ref
	{
		if seq.IsComponentPresent(4) {
			ie.FeatureSetCombinationLowBandSwitching_r19 = new(FeatureSetCombinationId)
			if err := ie.FeatureSetCombinationLowBandSwitching_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. switchingPeriodForFDD-SDL-r19: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(bandCombinationV1900SwitchingPeriodForFDDSDLR19Constraints)
			if err != nil {
				return err
			}
			ie.SwitchingPeriodForFDD_SDL_r19 = &idx
		}
	}

	// 8. directSCellActivation-r19: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(bandCombinationV1900DirectSCellActivationR19Constraints)
			if err != nil {
				return err
			}
			ie.DirectSCellActivation_r19 = &idx
		}
	}

	return nil
}
