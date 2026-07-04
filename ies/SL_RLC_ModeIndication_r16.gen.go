// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-RLC-ModeIndication-r16 (line 2290).

var sLRLCModeIndicationR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-Mode-r16"},
		{Name: "sl-QoS-InfoList-r16"},
	},
}

var sL_RLC_ModeIndication_r16SlModeR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl-AM-Mode-r16"},
		{Name: "sl-UM-Mode-r16"},
	},
}

const (
	SL_RLC_ModeIndication_r16_Sl_Mode_r16_Sl_AM_Mode_r16 = 0
	SL_RLC_ModeIndication_r16_Sl_Mode_r16_Sl_UM_Mode_r16 = 1
)

var sLRLCModeIndicationR16SlQoSInfoListR16Constraints = per.SizeRange(1, common.MaxNrofSL_QFIsPerDest_r16)

type SL_RLC_ModeIndication_r16 struct {
	Sl_Mode_r16 struct {
		Choice         int
		Sl_AM_Mode_r16 *struct{}
		Sl_UM_Mode_r16 *struct{}
	}
	Sl_QoS_InfoList_r16 []SL_QoS_Info_r16
}

func (ie *SL_RLC_ModeIndication_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLRLCModeIndicationR16Constraints)
	_ = seq

	// 1. sl-Mode-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(sL_RLC_ModeIndication_r16SlModeR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Sl_Mode_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Sl_Mode_r16.Choice {
		case SL_RLC_ModeIndication_r16_Sl_Mode_r16_Sl_AM_Mode_r16:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case SL_RLC_ModeIndication_r16_Sl_Mode_r16_Sl_UM_Mode_r16:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Sl_Mode_r16.Choice), Constraint: "index out of range"}
		}
	}

	// 2. sl-QoS-InfoList-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sLRLCModeIndicationR16SlQoSInfoListR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sl_QoS_InfoList_r16))); err != nil {
			return err
		}
		for i := range ie.Sl_QoS_InfoList_r16 {
			if err := ie.Sl_QoS_InfoList_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_RLC_ModeIndication_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLRLCModeIndicationR16Constraints)
	_ = seq

	// 1. sl-Mode-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(sL_RLC_ModeIndication_r16SlModeR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Sl_Mode_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SL_RLC_ModeIndication_r16_Sl_Mode_r16_Sl_AM_Mode_r16:
			ie.Sl_Mode_r16.Sl_AM_Mode_r16 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case SL_RLC_ModeIndication_r16_Sl_Mode_r16_Sl_UM_Mode_r16:
			ie.Sl_Mode_r16.Sl_UM_Mode_r16 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		}
	}

	// 2. sl-QoS-InfoList-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sLRLCModeIndicationR16SlQoSInfoListR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sl_QoS_InfoList_r16 = make([]SL_QoS_Info_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Sl_QoS_InfoList_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
