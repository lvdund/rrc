package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SDT_ConfigCommonSIB_r17 struct {
	Sdt_RSRP_Threshold_r17              *RSRP_Range                                                  `optional`
	Sdt_LogicalChannelSR_DelayTimer_r17 *SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17 `optional`
	Sdt_DataVolumeThreshold_r17         SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17          `madatory`
	T319a_r17                           SDT_ConfigCommonSIB_r17_t319a_r17                            `madatory`
}

func (ie *SDT_ConfigCommonSIB_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sdt_RSRP_Threshold_r17 != nil, ie.Sdt_LogicalChannelSR_DelayTimer_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sdt_RSRP_Threshold_r17 != nil {
		if err = ie.Sdt_RSRP_Threshold_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sdt_RSRP_Threshold_r17", err)
		}
	}
	if ie.Sdt_LogicalChannelSR_DelayTimer_r17 != nil {
		if err = ie.Sdt_LogicalChannelSR_DelayTimer_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sdt_LogicalChannelSR_DelayTimer_r17", err)
		}
	}
	if err = ie.Sdt_DataVolumeThreshold_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sdt_DataVolumeThreshold_r17", err)
	}
	if err = ie.T319a_r17.Encode(w); err != nil {
		return utils.WrapError("Encode T319a_r17", err)
	}
	return nil
}

func (ie *SDT_ConfigCommonSIB_r17) Decode(r *uper.UperReader) error {
	var err error
	var Sdt_RSRP_Threshold_r17Present bool
	if Sdt_RSRP_Threshold_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sdt_LogicalChannelSR_DelayTimer_r17Present bool
	if Sdt_LogicalChannelSR_DelayTimer_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sdt_RSRP_Threshold_r17Present {
		ie.Sdt_RSRP_Threshold_r17 = new(RSRP_Range)
		if err = ie.Sdt_RSRP_Threshold_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sdt_RSRP_Threshold_r17", err)
		}
	}
	if Sdt_LogicalChannelSR_DelayTimer_r17Present {
		ie.Sdt_LogicalChannelSR_DelayTimer_r17 = new(SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17)
		if err = ie.Sdt_LogicalChannelSR_DelayTimer_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sdt_LogicalChannelSR_DelayTimer_r17", err)
		}
	}
	if err = ie.Sdt_DataVolumeThreshold_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sdt_DataVolumeThreshold_r17", err)
	}
	if err = ie.T319a_r17.Decode(r); err != nil {
		return utils.WrapError("Decode T319a_r17", err)
	}
	return nil
}
