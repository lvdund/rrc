// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-SCCH-CarrierSetConfig-r18 (line 27030).

var sLSCCHCarrierSetConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DestinationList-r18"},
		{Name: "sl-SRB-Identity-r18"},
		{Name: "sl-AllowedCarrierFreqSet1-r18"},
		{Name: "sl-AllowedCarrierFreqSet2-r18"},
	},
}

var sLSCCHCarrierSetConfigR18SlDestinationListR18Constraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

var sLSCCHCarrierSetConfigR18SlSRBIdentityR18Constraints = per.SizeRange(1, 3)

var sLSCCHCarrierSetConfigR18SlAllowedCarrierFreqSet1R18Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

var sLSCCHCarrierSetConfigR18SlAllowedCarrierFreqSet2R18Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

type SL_SCCH_CarrierSetConfig_r18 struct {
	Sl_DestinationList_r18        []SL_DestinationIdentity_r16
	Sl_SRB_Identity_r18           []SRB_Identity
	Sl_AllowedCarrierFreqSet1_r18 []int64
	Sl_AllowedCarrierFreqSet2_r18 []int64
}

func (ie *SL_SCCH_CarrierSetConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLSCCHCarrierSetConfigR18Constraints)
	_ = seq

	// 1. sl-DestinationList-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sLSCCHCarrierSetConfigR18SlDestinationListR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sl_DestinationList_r18))); err != nil {
			return err
		}
		for i := range ie.Sl_DestinationList_r18 {
			if err := ie.Sl_DestinationList_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 2. sl-SRB-Identity-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sLSCCHCarrierSetConfigR18SlSRBIdentityR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sl_SRB_Identity_r18))); err != nil {
			return err
		}
		for i := range ie.Sl_SRB_Identity_r18 {
			if err := ie.Sl_SRB_Identity_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. sl-AllowedCarrierFreqSet1-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sLSCCHCarrierSetConfigR18SlAllowedCarrierFreqSet1R18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sl_AllowedCarrierFreqSet1_r18))); err != nil {
			return err
		}
		for i := range ie.Sl_AllowedCarrierFreqSet1_r18 {
			if err := e.EncodeInteger(ie.Sl_AllowedCarrierFreqSet1_r18[i], per.Constrained(1, common.MaxNrofFreqSL_r16)); err != nil {
				return err
			}
		}
	}

	// 4. sl-AllowedCarrierFreqSet2-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sLSCCHCarrierSetConfigR18SlAllowedCarrierFreqSet2R18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sl_AllowedCarrierFreqSet2_r18))); err != nil {
			return err
		}
		for i := range ie.Sl_AllowedCarrierFreqSet2_r18 {
			if err := e.EncodeInteger(ie.Sl_AllowedCarrierFreqSet2_r18[i], per.Constrained(1, common.MaxNrofFreqSL_r16)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_SCCH_CarrierSetConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLSCCHCarrierSetConfigR18Constraints)
	_ = seq

	// 1. sl-DestinationList-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sLSCCHCarrierSetConfigR18SlDestinationListR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sl_DestinationList_r18 = make([]SL_DestinationIdentity_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Sl_DestinationList_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 2. sl-SRB-Identity-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sLSCCHCarrierSetConfigR18SlSRBIdentityR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sl_SRB_Identity_r18 = make([]SRB_Identity, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Sl_SRB_Identity_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sl-AllowedCarrierFreqSet1-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sLSCCHCarrierSetConfigR18SlAllowedCarrierFreqSet1R18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sl_AllowedCarrierFreqSet1_r18 = make([]int64, n)
		for i := int64(0); i < n; i++ {
			v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofFreqSL_r16))
			if err != nil {
				return err
			}
			ie.Sl_AllowedCarrierFreqSet1_r18[i] = v
		}
	}

	// 4. sl-AllowedCarrierFreqSet2-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sLSCCHCarrierSetConfigR18SlAllowedCarrierFreqSet2R18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sl_AllowedCarrierFreqSet2_r18 = make([]int64, n)
		for i := int64(0); i < n; i++ {
			v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofFreqSL_r16))
			if err != nil {
				return err
			}
			ie.Sl_AllowedCarrierFreqSet2_r18[i] = v
		}
	}

	return nil
}
