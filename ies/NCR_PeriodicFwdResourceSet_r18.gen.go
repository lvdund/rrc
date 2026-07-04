// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NCR-PeriodicFwdResourceSet-r18 (line 10382).

var nCRPeriodicFwdResourceSetR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicFwdRsrcSetId-r18"},
		{Name: "periodicFwdRsrcToAddModList-r18", Optional: true},
		{Name: "periodicFwdRsrcToReleaseList-r18", Optional: true},
		{Name: "referenceSCS-r18", Optional: true},
		{Name: "priorityFlag-r18", Optional: true},
	},
}

var nCRPeriodicFwdResourceSetR18PeriodicFwdRsrcToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofPeriodicFwdResource_r18)

var nCRPeriodicFwdResourceSetR18PeriodicFwdRsrcToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofPeriodicFwdResource_r18)

const (
	NCR_PeriodicFwdResourceSet_r18_PriorityFlag_r18_True = 0
)

var nCRPeriodicFwdResourceSetR18PriorityFlagR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NCR_PeriodicFwdResourceSet_r18_PriorityFlag_r18_True},
}

type NCR_PeriodicFwdResourceSet_r18 struct {
	PeriodicFwdRsrcSetId_r18         NCR_PeriodicFwdResourceSetId_r18
	PeriodicFwdRsrcToAddModList_r18  []NCR_PeriodicFwdResource_r18
	PeriodicFwdRsrcToReleaseList_r18 []NCR_PeriodicFwdResourceId_r18
	ReferenceSCS_r18                 *SubcarrierSpacing
	PriorityFlag_r18                 *int64
}

func (ie *NCR_PeriodicFwdResourceSet_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nCRPeriodicFwdResourceSetR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PeriodicFwdRsrcToAddModList_r18 != nil, ie.PeriodicFwdRsrcToReleaseList_r18 != nil, ie.ReferenceSCS_r18 != nil, ie.PriorityFlag_r18 != nil}); err != nil {
		return err
	}

	// 3. periodicFwdRsrcSetId-r18: ref
	{
		if err := ie.PeriodicFwdRsrcSetId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. periodicFwdRsrcToAddModList-r18: sequence-of
	{
		if ie.PeriodicFwdRsrcToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(nCRPeriodicFwdResourceSetR18PeriodicFwdRsrcToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.PeriodicFwdRsrcToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.PeriodicFwdRsrcToAddModList_r18 {
				if err := ie.PeriodicFwdRsrcToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. periodicFwdRsrcToReleaseList-r18: sequence-of
	{
		if ie.PeriodicFwdRsrcToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(nCRPeriodicFwdResourceSetR18PeriodicFwdRsrcToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.PeriodicFwdRsrcToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.PeriodicFwdRsrcToReleaseList_r18 {
				if err := ie.PeriodicFwdRsrcToReleaseList_r18[i].Encode(e); err != nil {
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
			if err := e.EncodeEnumerated(*ie.PriorityFlag_r18, nCRPeriodicFwdResourceSetR18PriorityFlagR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NCR_PeriodicFwdResourceSet_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nCRPeriodicFwdResourceSetR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. periodicFwdRsrcSetId-r18: ref
	{
		if err := ie.PeriodicFwdRsrcSetId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. periodicFwdRsrcToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(nCRPeriodicFwdResourceSetR18PeriodicFwdRsrcToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PeriodicFwdRsrcToAddModList_r18 = make([]NCR_PeriodicFwdResource_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PeriodicFwdRsrcToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. periodicFwdRsrcToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(nCRPeriodicFwdResourceSetR18PeriodicFwdRsrcToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PeriodicFwdRsrcToReleaseList_r18 = make([]NCR_PeriodicFwdResourceId_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PeriodicFwdRsrcToReleaseList_r18[i].Decode(d); err != nil {
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
			idx, err := d.DecodeEnumerated(nCRPeriodicFwdResourceSetR18PriorityFlagR18Constraints)
			if err != nil {
				return err
			}
			ie.PriorityFlag_r18 = &idx
		}
	}

	return nil
}
