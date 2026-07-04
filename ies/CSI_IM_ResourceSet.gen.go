// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-IM-ResourceSet (line 6965).

var cSIIMResourceSetConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-IM-ResourceSetId"},
		{Name: "csi-IM-Resources"},
	},
}

var cSIIMResourceSetCsiIMResourcesConstraints = per.SizeRange(1, common.MaxNrofCSI_IM_ResourcesPerSet)

type CSI_IM_ResourceSet struct {
	Csi_IM_ResourceSetId CSI_IM_ResourceSetId
	Csi_IM_Resources     []CSI_IM_ResourceId
}

func (ie *CSI_IM_ResourceSet) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIIMResourceSetConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. csi-IM-ResourceSetId: ref
	{
		if err := ie.Csi_IM_ResourceSetId.Encode(e); err != nil {
			return err
		}
	}

	// 3. csi-IM-Resources: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(cSIIMResourceSetCsiIMResourcesConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Csi_IM_Resources))); err != nil {
			return err
		}
		for i := range ie.Csi_IM_Resources {
			if err := ie.Csi_IM_Resources[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CSI_IM_ResourceSet) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIIMResourceSetConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. csi-IM-ResourceSetId: ref
	{
		if err := ie.Csi_IM_ResourceSetId.Decode(d); err != nil {
			return err
		}
	}

	// 3. csi-IM-Resources: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(cSIIMResourceSetCsiIMResourcesConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Csi_IM_Resources = make([]CSI_IM_ResourceId, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Csi_IM_Resources[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
