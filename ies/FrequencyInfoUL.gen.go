// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FrequencyInfoUL (line 8429).

var frequencyInfoULConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyBandList", Optional: true},
		{Name: "absoluteFrequencyPointA", Optional: true},
		{Name: "scs-SpecificCarrierList"},
		{Name: "additionalSpectrumEmission", Optional: true},
		{Name: "p-Max", Optional: true},
		{Name: "frequencyShift7p5khz", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var frequencyInfoULScsSpecificCarrierListConstraints = per.SizeRange(1, common.MaxSCSs)

const (
	FrequencyInfoUL_FrequencyShift7p5khz_True = 0
)

var frequencyInfoULFrequencyShift7p5khzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FrequencyInfoUL_FrequencyShift7p5khz_True},
}

type FrequencyInfoUL struct {
	FrequencyBandList                    *MultiFrequencyBandListNR
	AbsoluteFrequencyPointA              *ARFCN_ValueNR
	Scs_SpecificCarrierList              []SCS_SpecificCarrier
	AdditionalSpectrumEmission           *AdditionalSpectrumEmission
	P_Max                                *P_Max
	FrequencyShift7p5khz                 *int64
	AdditionalSpectrumEmission_v1760     *AdditionalSpectrumEmission_v1760
	AdditionalSpectrumEmissionAerial_r18 *AdditionalSpectrumEmission_r18
}

func (ie *FrequencyInfoUL) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(frequencyInfoULConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.AdditionalSpectrumEmission_v1760 != nil
	hasExtGroup1 := ie.AdditionalSpectrumEmissionAerial_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FrequencyBandList != nil, ie.AbsoluteFrequencyPointA != nil, ie.AdditionalSpectrumEmission != nil, ie.P_Max != nil, ie.FrequencyShift7p5khz != nil}); err != nil {
		return err
	}

	// 3. frequencyBandList: ref
	{
		if ie.FrequencyBandList != nil {
			if err := ie.FrequencyBandList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. absoluteFrequencyPointA: ref
	{
		if ie.AbsoluteFrequencyPointA != nil {
			if err := ie.AbsoluteFrequencyPointA.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. scs-SpecificCarrierList: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(frequencyInfoULScsSpecificCarrierListConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Scs_SpecificCarrierList))); err != nil {
			return err
		}
		for i := range ie.Scs_SpecificCarrierList {
			if err := ie.Scs_SpecificCarrierList[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. additionalSpectrumEmission: ref
	{
		if ie.AdditionalSpectrumEmission != nil {
			if err := ie.AdditionalSpectrumEmission.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. p-Max: ref
	{
		if ie.P_Max != nil {
			if err := ie.P_Max.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. frequencyShift7p5khz: enumerated
	{
		if ie.FrequencyShift7p5khz != nil {
			if err := e.EncodeEnumerated(*ie.FrequencyShift7p5khz, frequencyInfoULFrequencyShift7p5khzConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "additionalSpectrumEmission-v1760", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AdditionalSpectrumEmission_v1760 != nil}); err != nil {
				return err
			}

			if ie.AdditionalSpectrumEmission_v1760 != nil {
				if err := ie.AdditionalSpectrumEmission_v1760.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "additionalSpectrumEmissionAerial-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AdditionalSpectrumEmissionAerial_r18 != nil}); err != nil {
				return err
			}

			if ie.AdditionalSpectrumEmissionAerial_r18 != nil {
				if err := ie.AdditionalSpectrumEmissionAerial_r18.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *FrequencyInfoUL) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(frequencyInfoULConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. frequencyBandList: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FrequencyBandList = new(MultiFrequencyBandListNR)
			if err := ie.FrequencyBandList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. absoluteFrequencyPointA: ref
	{
		if seq.IsComponentPresent(1) {
			ie.AbsoluteFrequencyPointA = new(ARFCN_ValueNR)
			if err := ie.AbsoluteFrequencyPointA.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. scs-SpecificCarrierList: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(frequencyInfoULScsSpecificCarrierListConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Scs_SpecificCarrierList = make([]SCS_SpecificCarrier, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Scs_SpecificCarrierList[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. additionalSpectrumEmission: ref
	{
		if seq.IsComponentPresent(3) {
			ie.AdditionalSpectrumEmission = new(AdditionalSpectrumEmission)
			if err := ie.AdditionalSpectrumEmission.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. p-Max: ref
	{
		if seq.IsComponentPresent(4) {
			ie.P_Max = new(P_Max)
			if err := ie.P_Max.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. frequencyShift7p5khz: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(frequencyInfoULFrequencyShift7p5khzConstraints)
			if err != nil {
				return err
			}
			ie.FrequencyShift7p5khz = &idx
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "additionalSpectrumEmission-v1760", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.AdditionalSpectrumEmission_v1760 = new(AdditionalSpectrumEmission_v1760)
			if err := ie.AdditionalSpectrumEmission_v1760.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "additionalSpectrumEmissionAerial-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.AdditionalSpectrumEmissionAerial_r18 = new(AdditionalSpectrumEmission_r18)
			if err := ie.AdditionalSpectrumEmissionAerial_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
