// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUCCH-ResourceSet (line 12064).

var pUCCHResourceSetConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pucch-ResourceSetId"},
		{Name: "resourceList"},
		{Name: "maxPayloadSize", Optional: true},
	},
}

var pUCCHResourceSetResourceListConstraints = per.SizeRange(1, common.MaxNrofPUCCH_ResourcesPerSet)

var pUCCHResourceSetMaxPayloadSizeConstraints = per.Constrained(4, 256)

type PUCCH_ResourceSet struct {
	Pucch_ResourceSetId PUCCH_ResourceSetId
	ResourceList        []PUCCH_ResourceId
	MaxPayloadSize      *int64
}

func (ie *PUCCH_ResourceSet) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHResourceSetConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxPayloadSize != nil}); err != nil {
		return err
	}

	// 2. pucch-ResourceSetId: ref
	{
		if err := ie.Pucch_ResourceSetId.Encode(e); err != nil {
			return err
		}
	}

	// 3. resourceList: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(pUCCHResourceSetResourceListConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.ResourceList))); err != nil {
			return err
		}
		for i := range ie.ResourceList {
			if err := ie.ResourceList[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. maxPayloadSize: integer
	{
		if ie.MaxPayloadSize != nil {
			if err := e.EncodeInteger(*ie.MaxPayloadSize, pUCCHResourceSetMaxPayloadSizeConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PUCCH_ResourceSet) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHResourceSetConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pucch-ResourceSetId: ref
	{
		if err := ie.Pucch_ResourceSetId.Decode(d); err != nil {
			return err
		}
	}

	// 3. resourceList: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(pUCCHResourceSetResourceListConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.ResourceList = make([]PUCCH_ResourceId, n)
		for i := int64(0); i < n; i++ {
			if err := ie.ResourceList[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. maxPayloadSize: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(pUCCHResourceSetMaxPayloadSizeConstraints)
			if err != nil {
				return err
			}
			ie.MaxPayloadSize = &v
		}
	}

	return nil
}
