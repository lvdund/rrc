// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ExcessDelay-DRB-IdentityInfo-r17 (line 16255).

var excessDelayDRBIdentityInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-IdentityList"},
		{Name: "delayThreshold"},
	},
}

var excessDelayDRBIdentityInfoR17DrbIdentityListConstraints = per.SizeRange(1, common.MaxDRB)

const (
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms0dot25 = 0
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms0dot5  = 1
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms1      = 2
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms2      = 3
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms4      = 4
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms5      = 5
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms10     = 6
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms20     = 7
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms30     = 8
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms40     = 9
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms50     = 10
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms60     = 11
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms70     = 12
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms80     = 13
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms90     = 14
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms100    = 15
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms150    = 16
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms300    = 17
	ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms500    = 18
)

var excessDelayDRBIdentityInfoR17DelayThresholdConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms0dot25, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms0dot5, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms1, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms2, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms4, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms5, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms10, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms20, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms30, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms40, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms50, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms60, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms70, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms80, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms90, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms100, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms150, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms300, ExcessDelay_DRB_IdentityInfo_r17_DelayThreshold_Ms500},
}

type ExcessDelay_DRB_IdentityInfo_r17 struct {
	Drb_IdentityList []DRB_Identity
	DelayThreshold   int64
}

func (ie *ExcessDelay_DRB_IdentityInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(excessDelayDRBIdentityInfoR17Constraints)
	_ = seq

	// 1. drb-IdentityList: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(excessDelayDRBIdentityInfoR17DrbIdentityListConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Drb_IdentityList))); err != nil {
			return err
		}
		for i := range ie.Drb_IdentityList {
			if err := ie.Drb_IdentityList[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 2. delayThreshold: enumerated
	{
		if err := e.EncodeEnumerated(ie.DelayThreshold, excessDelayDRBIdentityInfoR17DelayThresholdConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *ExcessDelay_DRB_IdentityInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(excessDelayDRBIdentityInfoR17Constraints)
	_ = seq

	// 1. drb-IdentityList: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(excessDelayDRBIdentityInfoR17DrbIdentityListConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Drb_IdentityList = make([]DRB_Identity, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Drb_IdentityList[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 2. delayThreshold: enumerated
	{
		v1, err := d.DecodeEnumerated(excessDelayDRBIdentityInfoR17DelayThresholdConstraints)
		if err != nil {
			return err
		}
		ie.DelayThreshold = v1
	}

	return nil
}
