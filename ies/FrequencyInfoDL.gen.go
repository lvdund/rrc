// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FrequencyInfoDL (line 8398).

var frequencyInfoDLConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "absoluteFrequencySSB", Optional: true},
		{Name: "frequencyBandList"},
		{Name: "absoluteFrequencyPointA"},
		{Name: "scs-SpecificCarrierList"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var frequencyInfoDLScsSpecificCarrierListConstraints = per.SizeRange(1, common.MaxSCSs)

type FrequencyInfoDL struct {
	AbsoluteFrequencySSB    *ARFCN_ValueNR
	FrequencyBandList       MultiFrequencyBandListNR
	AbsoluteFrequencyPointA ARFCN_ValueNR
	Scs_SpecificCarrierList []SCS_SpecificCarrier
	ReferenceCell_r18       *ServCellIndex
}

func (ie *FrequencyInfoDL) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(frequencyInfoDLConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ReferenceCell_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AbsoluteFrequencySSB != nil}); err != nil {
		return err
	}

	// 3. absoluteFrequencySSB: ref
	{
		if ie.AbsoluteFrequencySSB != nil {
			if err := ie.AbsoluteFrequencySSB.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. frequencyBandList: ref
	{
		if err := ie.FrequencyBandList.Encode(e); err != nil {
			return err
		}
	}

	// 5. absoluteFrequencyPointA: ref
	{
		if err := ie.AbsoluteFrequencyPointA.Encode(e); err != nil {
			return err
		}
	}

	// 6. scs-SpecificCarrierList: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(frequencyInfoDLScsSpecificCarrierListConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Scs_SpecificCarrierList))); err != nil {
			return err
		}
		for i := range ie.Scs_SpecificCarrierList {
			if err := ie.Scs_SpecificCarrierList[i].Encode(e); err != nil {
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
					{Name: "referenceCell-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ReferenceCell_r18 != nil}); err != nil {
				return err
			}

			if ie.ReferenceCell_r18 != nil {
				if err := ie.ReferenceCell_r18.Encode(ex); err != nil {
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

func (ie *FrequencyInfoDL) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(frequencyInfoDLConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. absoluteFrequencySSB: ref
	{
		if seq.IsComponentPresent(0) {
			ie.AbsoluteFrequencySSB = new(ARFCN_ValueNR)
			if err := ie.AbsoluteFrequencySSB.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. frequencyBandList: ref
	{
		if err := ie.FrequencyBandList.Decode(d); err != nil {
			return err
		}
	}

	// 5. absoluteFrequencyPointA: ref
	{
		if err := ie.AbsoluteFrequencyPointA.Decode(d); err != nil {
			return err
		}
	}

	// 6. scs-SpecificCarrierList: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(frequencyInfoDLScsSpecificCarrierListConstraints)
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
				{Name: "referenceCell-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ReferenceCell_r18 = new(ServCellIndex)
			if err := ie.ReferenceCell_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
