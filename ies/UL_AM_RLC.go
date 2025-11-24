package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UL_AM_RLC struct {
	Sn_FieldLength   *SN_FieldLengthAM          `optional`
	T_PollRetransmit T_PollRetransmit           `madatory`
	PollPDU          PollPDU                    `madatory`
	PollByte         PollByte                   `madatory`
	MaxRetxThreshold UL_AM_RLC_maxRetxThreshold `madatory`
}

func (ie *UL_AM_RLC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sn_FieldLength != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sn_FieldLength != nil {
		if err = ie.Sn_FieldLength.Encode(w); err != nil {
			return utils.WrapError("Encode Sn_FieldLength", err)
		}
	}
	if err = ie.T_PollRetransmit.Encode(w); err != nil {
		return utils.WrapError("Encode T_PollRetransmit", err)
	}
	if err = ie.PollPDU.Encode(w); err != nil {
		return utils.WrapError("Encode PollPDU", err)
	}
	if err = ie.PollByte.Encode(w); err != nil {
		return utils.WrapError("Encode PollByte", err)
	}
	if err = ie.MaxRetxThreshold.Encode(w); err != nil {
		return utils.WrapError("Encode MaxRetxThreshold", err)
	}
	return nil
}

func (ie *UL_AM_RLC) Decode(r *uper.UperReader) error {
	var err error
	var Sn_FieldLengthPresent bool
	if Sn_FieldLengthPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Sn_FieldLengthPresent {
		ie.Sn_FieldLength = new(SN_FieldLengthAM)
		if err = ie.Sn_FieldLength.Decode(r); err != nil {
			return utils.WrapError("Decode Sn_FieldLength", err)
		}
	}
	if err = ie.T_PollRetransmit.Decode(r); err != nil {
		return utils.WrapError("Decode T_PollRetransmit", err)
	}
	if err = ie.PollPDU.Decode(r); err != nil {
		return utils.WrapError("Decode PollPDU", err)
	}
	if err = ie.PollByte.Decode(r); err != nil {
		return utils.WrapError("Decode PollByte", err)
	}
	if err = ie.MaxRetxThreshold.Decode(r); err != nil {
		return utils.WrapError("Decode MaxRetxThreshold", err)
	}
	return nil
}
