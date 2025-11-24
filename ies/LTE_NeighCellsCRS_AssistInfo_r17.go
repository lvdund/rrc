package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LTE_NeighCellsCRS_AssistInfo_r17 struct {
	NeighCarrierBandwidthDL_r17       *LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17 `optional`
	NeighCarrierFreqDL_r17            *int64                                                        `lb:0,ub:16383,optional`
	NeighCellId_r17                   *EUTRA_PhysCellId                                             `optional`
	NeighCRS_muting_r17               *LTE_NeighCellsCRS_AssistInfo_r17_neighCRS_muting_r17         `optional`
	NeighMBSFN_SubframeConfigList_r17 *EUTRA_MBSFN_SubframeConfigList                               `optional`
	NeighNrofCRS_Ports_r17            *LTE_NeighCellsCRS_AssistInfo_r17_neighNrofCRS_Ports_r17      `optional`
	NeighV_Shift_r17                  *LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17            `optional`
}

func (ie *LTE_NeighCellsCRS_AssistInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.NeighCarrierBandwidthDL_r17 != nil, ie.NeighCarrierFreqDL_r17 != nil, ie.NeighCellId_r17 != nil, ie.NeighCRS_muting_r17 != nil, ie.NeighMBSFN_SubframeConfigList_r17 != nil, ie.NeighNrofCRS_Ports_r17 != nil, ie.NeighV_Shift_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.NeighCarrierBandwidthDL_r17 != nil {
		if err = ie.NeighCarrierBandwidthDL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NeighCarrierBandwidthDL_r17", err)
		}
	}
	if ie.NeighCarrierFreqDL_r17 != nil {
		if err = w.WriteInteger(*ie.NeighCarrierFreqDL_r17, &uper.Constraint{Lb: 0, Ub: 16383}, false); err != nil {
			return utils.WrapError("Encode NeighCarrierFreqDL_r17", err)
		}
	}
	if ie.NeighCellId_r17 != nil {
		if err = ie.NeighCellId_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NeighCellId_r17", err)
		}
	}
	if ie.NeighCRS_muting_r17 != nil {
		if err = ie.NeighCRS_muting_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NeighCRS_muting_r17", err)
		}
	}
	if ie.NeighMBSFN_SubframeConfigList_r17 != nil {
		if err = ie.NeighMBSFN_SubframeConfigList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NeighMBSFN_SubframeConfigList_r17", err)
		}
	}
	if ie.NeighNrofCRS_Ports_r17 != nil {
		if err = ie.NeighNrofCRS_Ports_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NeighNrofCRS_Ports_r17", err)
		}
	}
	if ie.NeighV_Shift_r17 != nil {
		if err = ie.NeighV_Shift_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NeighV_Shift_r17", err)
		}
	}
	return nil
}

func (ie *LTE_NeighCellsCRS_AssistInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var NeighCarrierBandwidthDL_r17Present bool
	if NeighCarrierBandwidthDL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeighCarrierFreqDL_r17Present bool
	if NeighCarrierFreqDL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeighCellId_r17Present bool
	if NeighCellId_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeighCRS_muting_r17Present bool
	if NeighCRS_muting_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeighMBSFN_SubframeConfigList_r17Present bool
	if NeighMBSFN_SubframeConfigList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeighNrofCRS_Ports_r17Present bool
	if NeighNrofCRS_Ports_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeighV_Shift_r17Present bool
	if NeighV_Shift_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if NeighCarrierBandwidthDL_r17Present {
		ie.NeighCarrierBandwidthDL_r17 = new(LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17)
		if err = ie.NeighCarrierBandwidthDL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NeighCarrierBandwidthDL_r17", err)
		}
	}
	if NeighCarrierFreqDL_r17Present {
		var tmp_int_NeighCarrierFreqDL_r17 int64
		if tmp_int_NeighCarrierFreqDL_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 16383}, false); err != nil {
			return utils.WrapError("Decode NeighCarrierFreqDL_r17", err)
		}
		ie.NeighCarrierFreqDL_r17 = &tmp_int_NeighCarrierFreqDL_r17
	}
	if NeighCellId_r17Present {
		ie.NeighCellId_r17 = new(EUTRA_PhysCellId)
		if err = ie.NeighCellId_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NeighCellId_r17", err)
		}
	}
	if NeighCRS_muting_r17Present {
		ie.NeighCRS_muting_r17 = new(LTE_NeighCellsCRS_AssistInfo_r17_neighCRS_muting_r17)
		if err = ie.NeighCRS_muting_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NeighCRS_muting_r17", err)
		}
	}
	if NeighMBSFN_SubframeConfigList_r17Present {
		ie.NeighMBSFN_SubframeConfigList_r17 = new(EUTRA_MBSFN_SubframeConfigList)
		if err = ie.NeighMBSFN_SubframeConfigList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NeighMBSFN_SubframeConfigList_r17", err)
		}
	}
	if NeighNrofCRS_Ports_r17Present {
		ie.NeighNrofCRS_Ports_r17 = new(LTE_NeighCellsCRS_AssistInfo_r17_neighNrofCRS_Ports_r17)
		if err = ie.NeighNrofCRS_Ports_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NeighNrofCRS_Ports_r17", err)
		}
	}
	if NeighV_Shift_r17Present {
		ie.NeighV_Shift_r17 = new(LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17)
		if err = ie.NeighV_Shift_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NeighV_Shift_r17", err)
		}
	}
	return nil
}
