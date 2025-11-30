package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqCarrierFreqInfo_v1730 struct {
	ChannelAccessMode2_r17 *InterFreqCarrierFreqInfo_v1730_channelAccessMode2_r17 `optional`
}

func (ie *InterFreqCarrierFreqInfo_v1730) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ChannelAccessMode2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ChannelAccessMode2_r17 != nil {
		if err = ie.ChannelAccessMode2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ChannelAccessMode2_r17", err)
		}
	}
	return nil
}

func (ie *InterFreqCarrierFreqInfo_v1730) Decode(r *aper.AperReader) error {
	var err error
	var ChannelAccessMode2_r17Present bool
	if ChannelAccessMode2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ChannelAccessMode2_r17Present {
		ie.ChannelAccessMode2_r17 = new(InterFreqCarrierFreqInfo_v1730_channelAccessMode2_r17)
		if err = ie.ChannelAccessMode2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ChannelAccessMode2_r17", err)
		}
	}
	return nil
}
