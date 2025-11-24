package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BAP_Config_r16 struct {
	Bap_Address_r16              *uper.BitString                             `lb:10,ub:10,optional`
	DefaultUL_BAP_RoutingID_r16  *BAP_RoutingID_r16                          `optional`
	DefaultUL_BH_RLC_Channel_r16 *BH_RLC_ChannelID_r16                       `optional`
	FlowControlFeedbackType_r16  *BAP_Config_r16_flowControlFeedbackType_r16 `optional`
}

func (ie *BAP_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Bap_Address_r16 != nil, ie.DefaultUL_BAP_RoutingID_r16 != nil, ie.DefaultUL_BH_RLC_Channel_r16 != nil, ie.FlowControlFeedbackType_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Bap_Address_r16 != nil {
		if err = w.WriteBitString(ie.Bap_Address_r16.Bytes, uint(ie.Bap_Address_r16.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Encode Bap_Address_r16", err)
		}
	}
	if ie.DefaultUL_BAP_RoutingID_r16 != nil {
		if err = ie.DefaultUL_BAP_RoutingID_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DefaultUL_BAP_RoutingID_r16", err)
		}
	}
	if ie.DefaultUL_BH_RLC_Channel_r16 != nil {
		if err = ie.DefaultUL_BH_RLC_Channel_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DefaultUL_BH_RLC_Channel_r16", err)
		}
	}
	if ie.FlowControlFeedbackType_r16 != nil {
		if err = ie.FlowControlFeedbackType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode FlowControlFeedbackType_r16", err)
		}
	}
	return nil
}

func (ie *BAP_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var Bap_Address_r16Present bool
	if Bap_Address_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DefaultUL_BAP_RoutingID_r16Present bool
	if DefaultUL_BAP_RoutingID_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DefaultUL_BH_RLC_Channel_r16Present bool
	if DefaultUL_BH_RLC_Channel_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var FlowControlFeedbackType_r16Present bool
	if FlowControlFeedbackType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Bap_Address_r16Present {
		var tmp_bs_Bap_Address_r16 []byte
		var n_Bap_Address_r16 uint
		if tmp_bs_Bap_Address_r16, n_Bap_Address_r16, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode Bap_Address_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_Bap_Address_r16,
			NumBits: uint64(n_Bap_Address_r16),
		}
		ie.Bap_Address_r16 = &tmp_bitstring
	}
	if DefaultUL_BAP_RoutingID_r16Present {
		ie.DefaultUL_BAP_RoutingID_r16 = new(BAP_RoutingID_r16)
		if err = ie.DefaultUL_BAP_RoutingID_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DefaultUL_BAP_RoutingID_r16", err)
		}
	}
	if DefaultUL_BH_RLC_Channel_r16Present {
		ie.DefaultUL_BH_RLC_Channel_r16 = new(BH_RLC_ChannelID_r16)
		if err = ie.DefaultUL_BH_RLC_Channel_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DefaultUL_BH_RLC_Channel_r16", err)
		}
	}
	if FlowControlFeedbackType_r16Present {
		ie.FlowControlFeedbackType_r16 = new(BAP_Config_r16_flowControlFeedbackType_r16)
		if err = ie.FlowControlFeedbackType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode FlowControlFeedbackType_r16", err)
		}
	}
	return nil
}
