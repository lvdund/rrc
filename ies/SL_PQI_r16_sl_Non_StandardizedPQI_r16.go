package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_PQI_r16_sl_Non_StandardizedPQI_r16 struct {
	Sl_ResourceType_r16       *SL_PQI_r16_sl_Non_StandardizedPQI_r16_sl_ResourceType_r16 `optional`
	Sl_PriorityLevel_r16      *int64                                                     `lb:1,ub:8,optional`
	Sl_PacketDelayBudget_r16  *int64                                                     `lb:0,ub:1023,optional`
	Sl_PacketErrorRate_r16    *int64                                                     `lb:0,ub:9,optional`
	Sl_AveragingWindow_r16    *int64                                                     `lb:0,ub:4095,optional`
	Sl_MaxDataBurstVolume_r16 *int64                                                     `lb:0,ub:4095,optional`
}

func (ie *SL_PQI_r16_sl_Non_StandardizedPQI_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_ResourceType_r16 != nil, ie.Sl_PriorityLevel_r16 != nil, ie.Sl_PacketDelayBudget_r16 != nil, ie.Sl_PacketErrorRate_r16 != nil, ie.Sl_AveragingWindow_r16 != nil, ie.Sl_MaxDataBurstVolume_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_ResourceType_r16 != nil {
		if err = ie.Sl_ResourceType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ResourceType_r16", err)
		}
	}
	if ie.Sl_PriorityLevel_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_PriorityLevel_r16, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Sl_PriorityLevel_r16", err)
		}
	}
	if ie.Sl_PacketDelayBudget_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_PacketDelayBudget_r16, &aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode Sl_PacketDelayBudget_r16", err)
		}
	}
	if ie.Sl_PacketErrorRate_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_PacketErrorRate_r16, &aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Encode Sl_PacketErrorRate_r16", err)
		}
	}
	if ie.Sl_AveragingWindow_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_AveragingWindow_r16, &aper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
			return utils.WrapError("Encode Sl_AveragingWindow_r16", err)
		}
	}
	if ie.Sl_MaxDataBurstVolume_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_MaxDataBurstVolume_r16, &aper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
			return utils.WrapError("Encode Sl_MaxDataBurstVolume_r16", err)
		}
	}
	return nil
}

func (ie *SL_PQI_r16_sl_Non_StandardizedPQI_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_ResourceType_r16Present bool
	if Sl_ResourceType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PriorityLevel_r16Present bool
	if Sl_PriorityLevel_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PacketDelayBudget_r16Present bool
	if Sl_PacketDelayBudget_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PacketErrorRate_r16Present bool
	if Sl_PacketErrorRate_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_AveragingWindow_r16Present bool
	if Sl_AveragingWindow_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MaxDataBurstVolume_r16Present bool
	if Sl_MaxDataBurstVolume_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_ResourceType_r16Present {
		ie.Sl_ResourceType_r16 = new(SL_PQI_r16_sl_Non_StandardizedPQI_r16_sl_ResourceType_r16)
		if err = ie.Sl_ResourceType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ResourceType_r16", err)
		}
	}
	if Sl_PriorityLevel_r16Present {
		var tmp_int_Sl_PriorityLevel_r16 int64
		if tmp_int_Sl_PriorityLevel_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Sl_PriorityLevel_r16", err)
		}
		ie.Sl_PriorityLevel_r16 = &tmp_int_Sl_PriorityLevel_r16
	}
	if Sl_PacketDelayBudget_r16Present {
		var tmp_int_Sl_PacketDelayBudget_r16 int64
		if tmp_int_Sl_PacketDelayBudget_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode Sl_PacketDelayBudget_r16", err)
		}
		ie.Sl_PacketDelayBudget_r16 = &tmp_int_Sl_PacketDelayBudget_r16
	}
	if Sl_PacketErrorRate_r16Present {
		var tmp_int_Sl_PacketErrorRate_r16 int64
		if tmp_int_Sl_PacketErrorRate_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Sl_PacketErrorRate_r16", err)
		}
		ie.Sl_PacketErrorRate_r16 = &tmp_int_Sl_PacketErrorRate_r16
	}
	if Sl_AveragingWindow_r16Present {
		var tmp_int_Sl_AveragingWindow_r16 int64
		if tmp_int_Sl_AveragingWindow_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
			return utils.WrapError("Decode Sl_AveragingWindow_r16", err)
		}
		ie.Sl_AveragingWindow_r16 = &tmp_int_Sl_AveragingWindow_r16
	}
	if Sl_MaxDataBurstVolume_r16Present {
		var tmp_int_Sl_MaxDataBurstVolume_r16 int64
		if tmp_int_Sl_MaxDataBurstVolume_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
			return utils.WrapError("Decode Sl_MaxDataBurstVolume_r16", err)
		}
		ie.Sl_MaxDataBurstVolume_r16 = &tmp_int_Sl_MaxDataBurstVolume_r16
	}
	return nil
}
