package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RAN_VisibleMeasurements_r17 struct {
	AppLayerBufferLevelList_r17     []AppLayerBufferLevel_r17 `lb:1,ub:8,optional`
	PlayoutDelayForMediaStartup_r17 *int64                    `lb:0,ub:30000,optional`
	Pdu_SessionIdList_r17           []PDU_SessionID           `lb:1,ub:maxNrofPDU_Sessions_r17,optional`
}

func (ie *RAN_VisibleMeasurements_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.AppLayerBufferLevelList_r17) > 0, ie.PlayoutDelayForMediaStartup_r17 != nil, len(ie.Pdu_SessionIdList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.AppLayerBufferLevelList_r17) > 0 {
		tmp_AppLayerBufferLevelList_r17 := utils.NewSequence[*AppLayerBufferLevel_r17]([]*AppLayerBufferLevel_r17{}, aper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.AppLayerBufferLevelList_r17 {
			tmp_AppLayerBufferLevelList_r17.Value = append(tmp_AppLayerBufferLevelList_r17.Value, &i)
		}
		if err = tmp_AppLayerBufferLevelList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AppLayerBufferLevelList_r17", err)
		}
	}
	if ie.PlayoutDelayForMediaStartup_r17 != nil {
		if err = w.WriteInteger(*ie.PlayoutDelayForMediaStartup_r17, &aper.Constraint{Lb: 0, Ub: 30000}, false); err != nil {
			return utils.WrapError("Encode PlayoutDelayForMediaStartup_r17", err)
		}
	}
	if len(ie.Pdu_SessionIdList_r17) > 0 {
		tmp_Pdu_SessionIdList_r17 := utils.NewSequence[*PDU_SessionID]([]*PDU_SessionID{}, aper.Constraint{Lb: 1, Ub: maxNrofPDU_Sessions_r17}, false)
		for _, i := range ie.Pdu_SessionIdList_r17 {
			tmp_Pdu_SessionIdList_r17.Value = append(tmp_Pdu_SessionIdList_r17.Value, &i)
		}
		if err = tmp_Pdu_SessionIdList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdu_SessionIdList_r17", err)
		}
	}
	return nil
}

func (ie *RAN_VisibleMeasurements_r17) Decode(r *aper.AperReader) error {
	var err error
	var AppLayerBufferLevelList_r17Present bool
	if AppLayerBufferLevelList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PlayoutDelayForMediaStartup_r17Present bool
	if PlayoutDelayForMediaStartup_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdu_SessionIdList_r17Present bool
	if Pdu_SessionIdList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if AppLayerBufferLevelList_r17Present {
		tmp_AppLayerBufferLevelList_r17 := utils.NewSequence[*AppLayerBufferLevel_r17]([]*AppLayerBufferLevel_r17{}, aper.Constraint{Lb: 1, Ub: 8}, false)
		fn_AppLayerBufferLevelList_r17 := func() *AppLayerBufferLevel_r17 {
			return new(AppLayerBufferLevel_r17)
		}
		if err = tmp_AppLayerBufferLevelList_r17.Decode(r, fn_AppLayerBufferLevelList_r17); err != nil {
			return utils.WrapError("Decode AppLayerBufferLevelList_r17", err)
		}
		ie.AppLayerBufferLevelList_r17 = []AppLayerBufferLevel_r17{}
		for _, i := range tmp_AppLayerBufferLevelList_r17.Value {
			ie.AppLayerBufferLevelList_r17 = append(ie.AppLayerBufferLevelList_r17, *i)
		}
	}
	if PlayoutDelayForMediaStartup_r17Present {
		var tmp_int_PlayoutDelayForMediaStartup_r17 int64
		if tmp_int_PlayoutDelayForMediaStartup_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 30000}, false); err != nil {
			return utils.WrapError("Decode PlayoutDelayForMediaStartup_r17", err)
		}
		ie.PlayoutDelayForMediaStartup_r17 = &tmp_int_PlayoutDelayForMediaStartup_r17
	}
	if Pdu_SessionIdList_r17Present {
		tmp_Pdu_SessionIdList_r17 := utils.NewSequence[*PDU_SessionID]([]*PDU_SessionID{}, aper.Constraint{Lb: 1, Ub: maxNrofPDU_Sessions_r17}, false)
		fn_Pdu_SessionIdList_r17 := func() *PDU_SessionID {
			return new(PDU_SessionID)
		}
		if err = tmp_Pdu_SessionIdList_r17.Decode(r, fn_Pdu_SessionIdList_r17); err != nil {
			return utils.WrapError("Decode Pdu_SessionIdList_r17", err)
		}
		ie.Pdu_SessionIdList_r17 = []PDU_SessionID{}
		for _, i := range tmp_Pdu_SessionIdList_r17.Value {
			ie.Pdu_SessionIdList_r17 = append(ie.Pdu_SessionIdList_r17, *i)
		}
	}
	return nil
}
