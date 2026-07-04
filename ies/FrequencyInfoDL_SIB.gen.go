// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FrequencyInfoDL-SIB (line 8412).

var frequencyInfoDLSIBConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyBandList"},
		{Name: "offsetToPointA"},
		{Name: "scs-SpecificCarrierList"},
	},
}

var frequencyInfoDLSIBOffsetToPointAConstraints = per.Constrained(0, 2199)

var frequencyInfoDLSIBScsSpecificCarrierListConstraints = per.SizeRange(1, common.MaxSCSs)

type FrequencyInfoDL_SIB struct {
	FrequencyBandList       MultiFrequencyBandListNR_SIB
	OffsetToPointA          int64
	Scs_SpecificCarrierList []SCS_SpecificCarrier
}

func (ie *FrequencyInfoDL_SIB) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(frequencyInfoDLSIBConstraints)
	_ = seq

	// 1. frequencyBandList: ref
	{
		if err := ie.FrequencyBandList.Encode(e); err != nil {
			return err
		}
	}

	// 2. offsetToPointA: integer
	{
		if err := e.EncodeInteger(ie.OffsetToPointA, frequencyInfoDLSIBOffsetToPointAConstraints); err != nil {
			return err
		}
	}

	// 3. scs-SpecificCarrierList: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(frequencyInfoDLSIBScsSpecificCarrierListConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Scs_SpecificCarrierList))); err != nil {
			return err
		}
		for i := range ie.Scs_SpecificCarrierList {
			if err := ie.Scs_SpecificCarrierList[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FrequencyInfoDL_SIB) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(frequencyInfoDLSIBConstraints)
	_ = seq

	// 1. frequencyBandList: ref
	{
		if err := ie.FrequencyBandList.Decode(d); err != nil {
			return err
		}
	}

	// 2. offsetToPointA: integer
	{
		v1, err := d.DecodeInteger(frequencyInfoDLSIBOffsetToPointAConstraints)
		if err != nil {
			return err
		}
		ie.OffsetToPointA = v1
	}

	// 3. scs-SpecificCarrierList: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(frequencyInfoDLSIBScsSpecificCarrierListConstraints)
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

	return nil
}
