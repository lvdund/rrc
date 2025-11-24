package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PDCP_Config_r16 struct {
	Sl_DiscardTimer_r16   *SL_PDCP_Config_r16_sl_DiscardTimer_r16   `optional`
	Sl_PDCP_SN_Size_r16   *SL_PDCP_Config_r16_sl_PDCP_SN_Size_r16   `optional`
	Sl_OutOfOrderDelivery *SL_PDCP_Config_r16_sl_OutOfOrderDelivery `optional`
}

func (ie *SL_PDCP_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_DiscardTimer_r16 != nil, ie.Sl_PDCP_SN_Size_r16 != nil, ie.Sl_OutOfOrderDelivery != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_DiscardTimer_r16 != nil {
		if err = ie.Sl_DiscardTimer_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DiscardTimer_r16", err)
		}
	}
	if ie.Sl_PDCP_SN_Size_r16 != nil {
		if err = ie.Sl_PDCP_SN_Size_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PDCP_SN_Size_r16", err)
		}
	}
	if ie.Sl_OutOfOrderDelivery != nil {
		if err = ie.Sl_OutOfOrderDelivery.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_OutOfOrderDelivery", err)
		}
	}
	return nil
}

func (ie *SL_PDCP_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_DiscardTimer_r16Present bool
	if Sl_DiscardTimer_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PDCP_SN_Size_r16Present bool
	if Sl_PDCP_SN_Size_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_OutOfOrderDeliveryPresent bool
	if Sl_OutOfOrderDeliveryPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_DiscardTimer_r16Present {
		ie.Sl_DiscardTimer_r16 = new(SL_PDCP_Config_r16_sl_DiscardTimer_r16)
		if err = ie.Sl_DiscardTimer_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_DiscardTimer_r16", err)
		}
	}
	if Sl_PDCP_SN_Size_r16Present {
		ie.Sl_PDCP_SN_Size_r16 = new(SL_PDCP_Config_r16_sl_PDCP_SN_Size_r16)
		if err = ie.Sl_PDCP_SN_Size_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PDCP_SN_Size_r16", err)
		}
	}
	if Sl_OutOfOrderDeliveryPresent {
		ie.Sl_OutOfOrderDelivery = new(SL_PDCP_Config_r16_sl_OutOfOrderDelivery)
		if err = ie.Sl_OutOfOrderDelivery.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_OutOfOrderDelivery", err)
		}
	}
	return nil
}
