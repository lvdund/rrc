// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RLC-Config-r16 (line 28172).

var sLRLCConfigR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl-AM-RLC-r16"},
		{Name: "sl-UM-RLC-r16"},
	},
}

const (
	SL_RLC_Config_r16_Sl_AM_RLC_r16 = 0
	SL_RLC_Config_r16_Sl_UM_RLC_r16 = 1
)

var sLRLCConfigR16SlAMRLCR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-SN-FieldLengthAM-r16", Optional: true},
		{Name: "sl-T-PollRetransmit-r16"},
		{Name: "sl-PollPDU-r16"},
		{Name: "sl-PollByte-r16"},
		{Name: "sl-MaxRetxThreshold-r16"},
	},
}

const (
	SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T1  = 0
	SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T2  = 1
	SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T3  = 2
	SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T4  = 3
	SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T6  = 4
	SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T8  = 5
	SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T16 = 6
	SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T32 = 7
)

var sLRLCConfigR16SlAMRLCR16SlMaxRetxThresholdR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T1, SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T2, SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T3, SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T4, SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T6, SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T8, SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T16, SL_RLC_Config_r16_Sl_AM_RLC_r16_Sl_MaxRetxThreshold_r16_T32},
}

var sLRLCConfigR16SlUMRLCR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-SN-FieldLengthUM-r16", Optional: true},
	},
}

type SL_RLC_Config_r16 struct {
	Choice        int
	Sl_AM_RLC_r16 *struct {
		Sl_SN_FieldLengthAM_r16 *SN_FieldLengthAM
		Sl_T_PollRetransmit_r16 T_PollRetransmit
		Sl_PollPDU_r16          PollPDU
		Sl_PollByte_r16         PollByte
		Sl_MaxRetxThreshold_r16 int64
	}
	Sl_UM_RLC_r16 *struct{ Sl_SN_FieldLengthUM_r16 *SN_FieldLengthUM }
}

func (ie *SL_RLC_Config_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sLRLCConfigR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_RLC_Config_r16_Sl_AM_RLC_r16:
		c := (*ie.Sl_AM_RLC_r16)
		sLRLCConfigR16SlAMRLCR16Seq := e.NewSequenceEncoder(sLRLCConfigR16SlAMRLCR16Constraints)
		if err := sLRLCConfigR16SlAMRLCR16Seq.EncodeExtensionBit(false); err != nil {
			return err
		}
		if err := sLRLCConfigR16SlAMRLCR16Seq.EncodePreamble([]bool{c.Sl_SN_FieldLengthAM_r16 != nil}); err != nil {
			return err
		}
		if c.Sl_SN_FieldLengthAM_r16 != nil {
			if err := c.Sl_SN_FieldLengthAM_r16.Encode(e); err != nil {
				return err
			}
		}
		if err := c.Sl_T_PollRetransmit_r16.Encode(e); err != nil {
			return err
		}
		if err := c.Sl_PollPDU_r16.Encode(e); err != nil {
			return err
		}
		if err := c.Sl_PollByte_r16.Encode(e); err != nil {
			return err
		}
		if err := e.EncodeEnumerated(c.Sl_MaxRetxThreshold_r16, sLRLCConfigR16SlAMRLCR16SlMaxRetxThresholdR16Constraints); err != nil {
			return err
		}
	case SL_RLC_Config_r16_Sl_UM_RLC_r16:
		c := (*ie.Sl_UM_RLC_r16)
		sLRLCConfigR16SlUMRLCR16Seq := e.NewSequenceEncoder(sLRLCConfigR16SlUMRLCR16Constraints)
		if err := sLRLCConfigR16SlUMRLCR16Seq.EncodeExtensionBit(false); err != nil {
			return err
		}
		if err := sLRLCConfigR16SlUMRLCR16Seq.EncodePreamble([]bool{c.Sl_SN_FieldLengthUM_r16 != nil}); err != nil {
			return err
		}
		if c.Sl_SN_FieldLengthUM_r16 != nil {
			if err := c.Sl_SN_FieldLengthUM_r16.Encode(e); err != nil {
				return err
			}
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "SL-RLC-Config-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SL_RLC_Config_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sLRLCConfigR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SL-RLC-Config-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SL_RLC_Config_r16_Sl_AM_RLC_r16:
		ie.Sl_AM_RLC_r16 = &struct {
			Sl_SN_FieldLengthAM_r16 *SN_FieldLengthAM
			Sl_T_PollRetransmit_r16 T_PollRetransmit
			Sl_PollPDU_r16          PollPDU
			Sl_PollByte_r16         PollByte
			Sl_MaxRetxThreshold_r16 int64
		}{}
		c := (*ie.Sl_AM_RLC_r16)
		sLRLCConfigR16SlAMRLCR16Seq := d.NewSequenceDecoder(sLRLCConfigR16SlAMRLCR16Constraints)
		if err := sLRLCConfigR16SlAMRLCR16Seq.DecodeExtensionBit(); err != nil {
			return err
		}
		if err := sLRLCConfigR16SlAMRLCR16Seq.DecodePreamble(); err != nil {
			return err
		}
		if sLRLCConfigR16SlAMRLCR16Seq.IsComponentPresent(0) {
			c.Sl_SN_FieldLengthAM_r16 = new(SN_FieldLengthAM)
			if err := (*c.Sl_SN_FieldLengthAM_r16).Decode(d); err != nil {
				return err
			}
		}
		{
			if err := c.Sl_T_PollRetransmit_r16.Decode(d); err != nil {
				return err
			}
		}
		{
			if err := c.Sl_PollPDU_r16.Decode(d); err != nil {
				return err
			}
		}
		{
			if err := c.Sl_PollByte_r16.Decode(d); err != nil {
				return err
			}
		}
		{
			v, err := d.DecodeEnumerated(sLRLCConfigR16SlAMRLCR16SlMaxRetxThresholdR16Constraints)
			if err != nil {
				return err
			}
			c.Sl_MaxRetxThreshold_r16 = v
		}
	case SL_RLC_Config_r16_Sl_UM_RLC_r16:
		ie.Sl_UM_RLC_r16 = &struct{ Sl_SN_FieldLengthUM_r16 *SN_FieldLengthUM }{}
		c := (*ie.Sl_UM_RLC_r16)
		sLRLCConfigR16SlUMRLCR16Seq := d.NewSequenceDecoder(sLRLCConfigR16SlUMRLCR16Constraints)
		if err := sLRLCConfigR16SlUMRLCR16Seq.DecodeExtensionBit(); err != nil {
			return err
		}
		if err := sLRLCConfigR16SlUMRLCR16Seq.DecodePreamble(); err != nil {
			return err
		}
		if sLRLCConfigR16SlUMRLCR16Seq.IsComponentPresent(0) {
			c.Sl_SN_FieldLengthUM_r16 = new(SN_FieldLengthUM)
			if err := (*c.Sl_SN_FieldLengthUM_r16).Decode(d); err != nil {
				return err
			}
		}
	default:
		return &per.DecodeError{TypeName: "SL-RLC-Config-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
