package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RLC_Config_r16_sl_AM_RLC_r16 struct {
	Sl_SN_FieldLengthAM_r16 *SN_FieldLengthAM                                       `optional`
	Sl_T_PollRetransmit_r16 T_PollRetransmit                                        `madatory`
	Sl_PollPDU_r16          PollPDU                                                 `madatory`
	Sl_PollByte_r16         PollByte                                                `madatory`
	Sl_MaxRetxThreshold_r16 SL_RLC_Config_r16_sl_AM_RLC_r16_sl_MaxRetxThreshold_r16 `madatory`
}

func (ie *SL_RLC_Config_r16_sl_AM_RLC_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_SN_FieldLengthAM_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_SN_FieldLengthAM_r16 != nil {
		if err = ie.Sl_SN_FieldLengthAM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SN_FieldLengthAM_r16", err)
		}
	}
	if err = ie.Sl_T_PollRetransmit_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_T_PollRetransmit_r16", err)
	}
	if err = ie.Sl_PollPDU_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_PollPDU_r16", err)
	}
	if err = ie.Sl_PollByte_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_PollByte_r16", err)
	}
	if err = ie.Sl_MaxRetxThreshold_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_MaxRetxThreshold_r16", err)
	}
	return nil
}

func (ie *SL_RLC_Config_r16_sl_AM_RLC_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_SN_FieldLengthAM_r16Present bool
	if Sl_SN_FieldLengthAM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_SN_FieldLengthAM_r16Present {
		ie.Sl_SN_FieldLengthAM_r16 = new(SN_FieldLengthAM)
		if err = ie.Sl_SN_FieldLengthAM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SN_FieldLengthAM_r16", err)
		}
	}
	if err = ie.Sl_T_PollRetransmit_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_T_PollRetransmit_r16", err)
	}
	if err = ie.Sl_PollPDU_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_PollPDU_r16", err)
	}
	if err = ie.Sl_PollByte_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_PollByte_r16", err)
	}
	if err = ie.Sl_MaxRetxThreshold_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_MaxRetxThreshold_r16", err)
	}
	return nil
}
