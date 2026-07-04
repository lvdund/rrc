// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ZP-CSI-RS-ResourceSet (line 16457).

var zPCSIRSResourceSetConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "zp-CSI-RS-ResourceSetId"},
		{Name: "zp-CSI-RS-ResourceIdList"},
	},
}

var zPCSIRSResourceSetZpCSIRSResourceIdListConstraints = per.SizeRange(1, common.MaxNrofZP_CSI_RS_ResourcesPerSet)

type ZP_CSI_RS_ResourceSet struct {
	Zp_CSI_RS_ResourceSetId  ZP_CSI_RS_ResourceSetId
	Zp_CSI_RS_ResourceIdList []ZP_CSI_RS_ResourceId
}

func (ie *ZP_CSI_RS_ResourceSet) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(zPCSIRSResourceSetConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. zp-CSI-RS-ResourceSetId: ref
	{
		if err := ie.Zp_CSI_RS_ResourceSetId.Encode(e); err != nil {
			return err
		}
	}

	// 3. zp-CSI-RS-ResourceIdList: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(zPCSIRSResourceSetZpCSIRSResourceIdListConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Zp_CSI_RS_ResourceIdList))); err != nil {
			return err
		}
		for i := range ie.Zp_CSI_RS_ResourceIdList {
			if err := ie.Zp_CSI_RS_ResourceIdList[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ZP_CSI_RS_ResourceSet) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(zPCSIRSResourceSetConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. zp-CSI-RS-ResourceSetId: ref
	{
		if err := ie.Zp_CSI_RS_ResourceSetId.Decode(d); err != nil {
			return err
		}
	}

	// 3. zp-CSI-RS-ResourceIdList: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(zPCSIRSResourceSetZpCSIRSResourceIdListConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Zp_CSI_RS_ResourceIdList = make([]ZP_CSI_RS_ResourceId, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Zp_CSI_RS_ResourceIdList[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
