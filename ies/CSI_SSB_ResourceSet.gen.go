// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-SSB-ResourceSet (line 7614).

var cSISSBResourceSetConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-SSB-ResourceSetId"},
		{Name: "csi-SSB-ResourceList"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var cSISSBResourceSetCsiSSBResourceListConstraints = per.SizeRange(1, common.MaxNrofCSI_SSB_ResourcePerSet)

var cSISSBResourceSetExtServingAdditionalPCIListR17Constraints = per.SizeRange(1, common.MaxNrofCSI_SSB_ResourcePerSet)

type CSI_SSB_ResourceSet struct {
	Csi_SSB_ResourceSetId        CSI_SSB_ResourceSetId
	Csi_SSB_ResourceList         []SSB_Index
	ServingAdditionalPCIList_r17 []ServingAdditionalPCIIndex_r17
}

func (ie *CSI_SSB_ResourceSet) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSISSBResourceSetConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ServingAdditionalPCIList_r17 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. csi-SSB-ResourceSetId: ref
	{
		if err := ie.Csi_SSB_ResourceSetId.Encode(e); err != nil {
			return err
		}
	}

	// 3. csi-SSB-ResourceList: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(cSISSBResourceSetCsiSSBResourceListConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Csi_SSB_ResourceList))); err != nil {
			return err
		}
		for i := range ie.Csi_SSB_ResourceList {
			if err := ie.Csi_SSB_ResourceList[i].Encode(e); err != nil {
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
					{Name: "servingAdditionalPCIList-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ServingAdditionalPCIList_r17 != nil}); err != nil {
				return err
			}

			if ie.ServingAdditionalPCIList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSISSBResourceSetExtServingAdditionalPCIListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.ServingAdditionalPCIList_r17))); err != nil {
					return err
				}
				for i := range ie.ServingAdditionalPCIList_r17 {
					if err := ie.ServingAdditionalPCIList_r17[i].Encode(ex); err != nil {
						return err
					}
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

func (ie *CSI_SSB_ResourceSet) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSISSBResourceSetConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. csi-SSB-ResourceSetId: ref
	{
		if err := ie.Csi_SSB_ResourceSetId.Decode(d); err != nil {
			return err
		}
	}

	// 3. csi-SSB-ResourceList: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(cSISSBResourceSetCsiSSBResourceListConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Csi_SSB_ResourceList = make([]SSB_Index, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Csi_SSB_ResourceList[i].Decode(d); err != nil {
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
				{Name: "servingAdditionalPCIList-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(cSISSBResourceSetExtServingAdditionalPCIListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ServingAdditionalPCIList_r17 = make([]ServingAdditionalPCIIndex_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ServingAdditionalPCIList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
