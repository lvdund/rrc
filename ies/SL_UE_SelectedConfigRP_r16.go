package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_UE_SelectedConfigRP_r16 struct {
	Sl_CBR_PriorityTxConfigList_r16   *SL_CBR_PriorityTxConfigList_r16                        `optional`
	Sl_Thres_RSRP_List_r16            *SL_Thres_RSRP_List_r16                                 `optional`
	Sl_MultiReserveResource_r16       *SL_UE_SelectedConfigRP_r16_sl_MultiReserveResource_r16 `optional`
	Sl_MaxNumPerReserve_r16           *SL_UE_SelectedConfigRP_r16_sl_MaxNumPerReserve_r16     `optional`
	Sl_SensingWindow_r16              *SL_UE_SelectedConfigRP_r16_sl_SensingWindow_r16        `optional`
	Sl_SelectionWindowList_r16        *SL_SelectionWindowList_r16                             `optional`
	Sl_ResourceReservePeriodList_r16  []SL_ResourceReservePeriod_r16                          `lb:1,ub:16,optional`
	Sl_RS_ForSensing_r16              SL_UE_SelectedConfigRP_r16_sl_RS_ForSensing_r16         `madatory`
	Sl_CBR_PriorityTxConfigList_v1650 *SL_CBR_PriorityTxConfigList_v1650                      `optional,ext-1`
}

func (ie *SL_UE_SelectedConfigRP_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Sl_CBR_PriorityTxConfigList_v1650 != nil
	preambleBits := []bool{hasExtensions, ie.Sl_CBR_PriorityTxConfigList_r16 != nil, ie.Sl_Thres_RSRP_List_r16 != nil, ie.Sl_MultiReserveResource_r16 != nil, ie.Sl_MaxNumPerReserve_r16 != nil, ie.Sl_SensingWindow_r16 != nil, ie.Sl_SelectionWindowList_r16 != nil, len(ie.Sl_ResourceReservePeriodList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_CBR_PriorityTxConfigList_r16 != nil {
		if err = ie.Sl_CBR_PriorityTxConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_CBR_PriorityTxConfigList_r16", err)
		}
	}
	if ie.Sl_Thres_RSRP_List_r16 != nil {
		if err = ie.Sl_Thres_RSRP_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_Thres_RSRP_List_r16", err)
		}
	}
	if ie.Sl_MultiReserveResource_r16 != nil {
		if err = ie.Sl_MultiReserveResource_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MultiReserveResource_r16", err)
		}
	}
	if ie.Sl_MaxNumPerReserve_r16 != nil {
		if err = ie.Sl_MaxNumPerReserve_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MaxNumPerReserve_r16", err)
		}
	}
	if ie.Sl_SensingWindow_r16 != nil {
		if err = ie.Sl_SensingWindow_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SensingWindow_r16", err)
		}
	}
	if ie.Sl_SelectionWindowList_r16 != nil {
		if err = ie.Sl_SelectionWindowList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SelectionWindowList_r16", err)
		}
	}
	if len(ie.Sl_ResourceReservePeriodList_r16) > 0 {
		tmp_Sl_ResourceReservePeriodList_r16 := utils.NewSequence[*SL_ResourceReservePeriod_r16]([]*SL_ResourceReservePeriod_r16{}, aper.Constraint{Lb: 1, Ub: 16}, false)
		for _, i := range ie.Sl_ResourceReservePeriodList_r16 {
			tmp_Sl_ResourceReservePeriodList_r16.Value = append(tmp_Sl_ResourceReservePeriodList_r16.Value, &i)
		}
		if err = tmp_Sl_ResourceReservePeriodList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ResourceReservePeriodList_r16", err)
		}
	}
	if err = ie.Sl_RS_ForSensing_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_RS_ForSensing_r16", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Sl_CBR_PriorityTxConfigList_v1650 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_UE_SelectedConfigRP_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Sl_CBR_PriorityTxConfigList_v1650 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_CBR_PriorityTxConfigList_v1650 optional
			if ie.Sl_CBR_PriorityTxConfigList_v1650 != nil {
				if err = ie.Sl_CBR_PriorityTxConfigList_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_CBR_PriorityTxConfigList_v1650", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *SL_UE_SelectedConfigRP_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CBR_PriorityTxConfigList_r16Present bool
	if Sl_CBR_PriorityTxConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_Thres_RSRP_List_r16Present bool
	if Sl_Thres_RSRP_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MultiReserveResource_r16Present bool
	if Sl_MultiReserveResource_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MaxNumPerReserve_r16Present bool
	if Sl_MaxNumPerReserve_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SensingWindow_r16Present bool
	if Sl_SensingWindow_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SelectionWindowList_r16Present bool
	if Sl_SelectionWindowList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ResourceReservePeriodList_r16Present bool
	if Sl_ResourceReservePeriodList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_CBR_PriorityTxConfigList_r16Present {
		ie.Sl_CBR_PriorityTxConfigList_r16 = new(SL_CBR_PriorityTxConfigList_r16)
		if err = ie.Sl_CBR_PriorityTxConfigList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_CBR_PriorityTxConfigList_r16", err)
		}
	}
	if Sl_Thres_RSRP_List_r16Present {
		ie.Sl_Thres_RSRP_List_r16 = new(SL_Thres_RSRP_List_r16)
		if err = ie.Sl_Thres_RSRP_List_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_Thres_RSRP_List_r16", err)
		}
	}
	if Sl_MultiReserveResource_r16Present {
		ie.Sl_MultiReserveResource_r16 = new(SL_UE_SelectedConfigRP_r16_sl_MultiReserveResource_r16)
		if err = ie.Sl_MultiReserveResource_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MultiReserveResource_r16", err)
		}
	}
	if Sl_MaxNumPerReserve_r16Present {
		ie.Sl_MaxNumPerReserve_r16 = new(SL_UE_SelectedConfigRP_r16_sl_MaxNumPerReserve_r16)
		if err = ie.Sl_MaxNumPerReserve_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MaxNumPerReserve_r16", err)
		}
	}
	if Sl_SensingWindow_r16Present {
		ie.Sl_SensingWindow_r16 = new(SL_UE_SelectedConfigRP_r16_sl_SensingWindow_r16)
		if err = ie.Sl_SensingWindow_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SensingWindow_r16", err)
		}
	}
	if Sl_SelectionWindowList_r16Present {
		ie.Sl_SelectionWindowList_r16 = new(SL_SelectionWindowList_r16)
		if err = ie.Sl_SelectionWindowList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SelectionWindowList_r16", err)
		}
	}
	if Sl_ResourceReservePeriodList_r16Present {
		tmp_Sl_ResourceReservePeriodList_r16 := utils.NewSequence[*SL_ResourceReservePeriod_r16]([]*SL_ResourceReservePeriod_r16{}, aper.Constraint{Lb: 1, Ub: 16}, false)
		fn_Sl_ResourceReservePeriodList_r16 := func() *SL_ResourceReservePeriod_r16 {
			return new(SL_ResourceReservePeriod_r16)
		}
		if err = tmp_Sl_ResourceReservePeriodList_r16.Decode(r, fn_Sl_ResourceReservePeriodList_r16); err != nil {
			return utils.WrapError("Decode Sl_ResourceReservePeriodList_r16", err)
		}
		ie.Sl_ResourceReservePeriodList_r16 = []SL_ResourceReservePeriod_r16{}
		for _, i := range tmp_Sl_ResourceReservePeriodList_r16.Value {
			ie.Sl_ResourceReservePeriodList_r16 = append(ie.Sl_ResourceReservePeriodList_r16, *i)
		}
	}
	if err = ie.Sl_RS_ForSensing_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_RS_ForSensing_r16", err)
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Sl_CBR_PriorityTxConfigList_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_CBR_PriorityTxConfigList_v1650 optional
			if Sl_CBR_PriorityTxConfigList_v1650Present {
				ie.Sl_CBR_PriorityTxConfigList_v1650 = new(SL_CBR_PriorityTxConfigList_v1650)
				if err = ie.Sl_CBR_PriorityTxConfigList_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_CBR_PriorityTxConfigList_v1650", err)
				}
			}
		}
	}
	return nil
}
