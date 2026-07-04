// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NCR-SemiPersistentFwdResourceSet-r18 (line 10414).

var nCRSemiPersistentFwdResourceSetR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "semiPersistentFwdRsrcSetId-r18"},
		{Name: "semiPersistentFwdRsrcToAddModList-r18", Optional: true},
		{Name: "semiPersistentFwdRsrcToReleaseList-r18", Optional: true},
		{Name: "referenceSCS-r18", Optional: true},
		{Name: "priorityFlag-r18", Optional: true},
	},
}

var nCRSemiPersistentFwdResourceSetR18SemiPersistentFwdRsrcToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofSemiPersistentFwdResource_r18)

var nCRSemiPersistentFwdResourceSetR18SemiPersistentFwdRsrcToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofSemiPersistentFwdResource_r18)

const (
	NCR_SemiPersistentFwdResourceSet_r18_PriorityFlag_r18_True = 0
)

var nCRSemiPersistentFwdResourceSetR18PriorityFlagR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NCR_SemiPersistentFwdResourceSet_r18_PriorityFlag_r18_True},
}

type NCR_SemiPersistentFwdResourceSet_r18 struct {
	SemiPersistentFwdRsrcSetId_r18         NCR_SemiPersistentFwdResourceSetId_r18
	SemiPersistentFwdRsrcToAddModList_r18  []NCR_SemiPersistentFwdResource_r18
	SemiPersistentFwdRsrcToReleaseList_r18 []NCR_SemiPersistentFwdResourceId_r18
	ReferenceSCS_r18                       *SubcarrierSpacing
	PriorityFlag_r18                       *int64
}

func (ie *NCR_SemiPersistentFwdResourceSet_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nCRSemiPersistentFwdResourceSetR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SemiPersistentFwdRsrcToAddModList_r18 != nil, ie.SemiPersistentFwdRsrcToReleaseList_r18 != nil, ie.ReferenceSCS_r18 != nil, ie.PriorityFlag_r18 != nil}); err != nil {
		return err
	}

	// 3. semiPersistentFwdRsrcSetId-r18: ref
	{
		if err := ie.SemiPersistentFwdRsrcSetId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. semiPersistentFwdRsrcToAddModList-r18: sequence-of
	{
		if ie.SemiPersistentFwdRsrcToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(nCRSemiPersistentFwdResourceSetR18SemiPersistentFwdRsrcToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SemiPersistentFwdRsrcToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.SemiPersistentFwdRsrcToAddModList_r18 {
				if err := ie.SemiPersistentFwdRsrcToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. semiPersistentFwdRsrcToReleaseList-r18: sequence-of
	{
		if ie.SemiPersistentFwdRsrcToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(nCRSemiPersistentFwdResourceSetR18SemiPersistentFwdRsrcToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SemiPersistentFwdRsrcToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.SemiPersistentFwdRsrcToReleaseList_r18 {
				if err := ie.SemiPersistentFwdRsrcToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. referenceSCS-r18: ref
	{
		if ie.ReferenceSCS_r18 != nil {
			if err := ie.ReferenceSCS_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. priorityFlag-r18: enumerated
	{
		if ie.PriorityFlag_r18 != nil {
			if err := e.EncodeEnumerated(*ie.PriorityFlag_r18, nCRSemiPersistentFwdResourceSetR18PriorityFlagR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NCR_SemiPersistentFwdResourceSet_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nCRSemiPersistentFwdResourceSetR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. semiPersistentFwdRsrcSetId-r18: ref
	{
		if err := ie.SemiPersistentFwdRsrcSetId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. semiPersistentFwdRsrcToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(nCRSemiPersistentFwdResourceSetR18SemiPersistentFwdRsrcToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SemiPersistentFwdRsrcToAddModList_r18 = make([]NCR_SemiPersistentFwdResource_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SemiPersistentFwdRsrcToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. semiPersistentFwdRsrcToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(nCRSemiPersistentFwdResourceSetR18SemiPersistentFwdRsrcToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SemiPersistentFwdRsrcToReleaseList_r18 = make([]NCR_SemiPersistentFwdResourceId_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SemiPersistentFwdRsrcToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. referenceSCS-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.ReferenceSCS_r18 = new(SubcarrierSpacing)
			if err := ie.ReferenceSCS_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. priorityFlag-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(nCRSemiPersistentFwdResourceSetR18PriorityFlagR18Constraints)
			if err != nil {
				return err
			}
			ie.PriorityFlag_r18 = &idx
		}
	}

	return nil
}
