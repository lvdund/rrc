// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FrequencyInfoUL-SIB (line 8448).

var frequencyInfoULSIBConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyBandList", Optional: true},
		{Name: "absoluteFrequencyPointA", Optional: true},
		{Name: "scs-SpecificCarrierList"},
		{Name: "p-Max", Optional: true},
		{Name: "frequencyShift7p5khz", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var frequencyInfoULSIBScsSpecificCarrierListConstraints = per.SizeRange(1, common.MaxSCSs)

const (
	FrequencyInfoUL_SIB_FrequencyShift7p5khz_True = 0
)

var frequencyInfoULSIBFrequencyShift7p5khzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FrequencyInfoUL_SIB_FrequencyShift7p5khz_True},
}

type FrequencyInfoUL_SIB struct {
	FrequencyBandList           *MultiFrequencyBandListNR_SIB
	AbsoluteFrequencyPointA     *ARFCN_ValueNR
	Scs_SpecificCarrierList     []SCS_SpecificCarrier
	P_Max                       *P_Max
	FrequencyShift7p5khz        *int64
	FrequencyBandListAerial_r18 *MultiFrequencyBandListNR_Aerial_SIB_r18
}

func (ie *FrequencyInfoUL_SIB) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(frequencyInfoULSIBConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.FrequencyBandListAerial_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FrequencyBandList != nil, ie.AbsoluteFrequencyPointA != nil, ie.P_Max != nil, ie.FrequencyShift7p5khz != nil}); err != nil {
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
		seqOf := e.NewSequenceOfEncoder(frequencyInfoULSIBScsSpecificCarrierListConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Scs_SpecificCarrierList))); err != nil {
			return err
		}
		for i := range ie.Scs_SpecificCarrierList {
			if err := ie.Scs_SpecificCarrierList[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. p-Max: ref
	{
		if ie.P_Max != nil {
			if err := ie.P_Max.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. frequencyShift7p5khz: enumerated
	{
		if ie.FrequencyShift7p5khz != nil {
			if err := e.EncodeEnumerated(*ie.FrequencyShift7p5khz, frequencyInfoULSIBFrequencyShift7p5khzConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "frequencyBandListAerial-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FrequencyBandListAerial_r18 != nil}); err != nil {
				return err
			}

			if ie.FrequencyBandListAerial_r18 != nil {
				if err := ie.FrequencyBandListAerial_r18.Encode(ex); err != nil {
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

func (ie *FrequencyInfoUL_SIB) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(frequencyInfoULSIBConstraints)

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
			ie.FrequencyBandList = new(MultiFrequencyBandListNR_SIB)
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
		seqOf := d.NewSequenceOfDecoder(frequencyInfoULSIBScsSpecificCarrierListConstraints)
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

	// 6. p-Max: ref
	{
		if seq.IsComponentPresent(3) {
			ie.P_Max = new(P_Max)
			if err := ie.P_Max.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. frequencyShift7p5khz: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(frequencyInfoULSIBFrequencyShift7p5khzConstraints)
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
				{Name: "frequencyBandListAerial-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.FrequencyBandListAerial_r18 = new(MultiFrequencyBandListNR_Aerial_SIB_r18)
			if err := ie.FrequencyBandListAerial_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
