package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BAP_Parameters_r16 struct {
	FlowControlBH_RLC_ChannelBased_r16 *BAP_Parameters_r16_flowControlBH_RLC_ChannelBased_r16 `optional`
	FlowControlRouting_ID_Based_r16    *BAP_Parameters_r16_flowControlRouting_ID_Based_r16    `optional`
}

func (ie *BAP_Parameters_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.FlowControlBH_RLC_ChannelBased_r16 != nil, ie.FlowControlRouting_ID_Based_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FlowControlBH_RLC_ChannelBased_r16 != nil {
		if err = ie.FlowControlBH_RLC_ChannelBased_r16.Encode(w); err != nil {
			return utils.WrapError("Encode FlowControlBH_RLC_ChannelBased_r16", err)
		}
	}
	if ie.FlowControlRouting_ID_Based_r16 != nil {
		if err = ie.FlowControlRouting_ID_Based_r16.Encode(w); err != nil {
			return utils.WrapError("Encode FlowControlRouting_ID_Based_r16", err)
		}
	}
	return nil
}

func (ie *BAP_Parameters_r16) Decode(r *aper.AperReader) error {
	var err error
	var FlowControlBH_RLC_ChannelBased_r16Present bool
	if FlowControlBH_RLC_ChannelBased_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var FlowControlRouting_ID_Based_r16Present bool
	if FlowControlRouting_ID_Based_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if FlowControlBH_RLC_ChannelBased_r16Present {
		ie.FlowControlBH_RLC_ChannelBased_r16 = new(BAP_Parameters_r16_flowControlBH_RLC_ChannelBased_r16)
		if err = ie.FlowControlBH_RLC_ChannelBased_r16.Decode(r); err != nil {
			return utils.WrapError("Decode FlowControlBH_RLC_ChannelBased_r16", err)
		}
	}
	if FlowControlRouting_ID_Based_r16Present {
		ie.FlowControlRouting_ID_Based_r16 = new(BAP_Parameters_r16_flowControlRouting_ID_Based_r16)
		if err = ie.FlowControlRouting_ID_Based_r16.Decode(r); err != nil {
			return utils.WrapError("Decode FlowControlRouting_ID_Based_r16", err)
		}
	}
	return nil
}
