package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_TxResourceReq_r16 struct {
	Sl_DestinationIdentity_r16           SL_DestinationIdentity_r16           `madatory`
	Sl_CastType_r16                      SL_TxResourceReq_r16_sl_CastType_r16 `madatory`
	Sl_RLC_ModeIndicationList_r16        []SL_RLC_ModeIndication_r16          `lb:1,ub:maxNrofSLRB_r16,optional`
	Sl_QoS_InfoList_r16                  []SL_QoS_Info_r16                    `lb:1,ub:maxNrofSL_QFIsPerDest_r16,optional`
	Sl_TypeTxSyncList_r16                []SL_TypeTxSync_r16                  `lb:1,ub:maxNrofFreqSL_r16,optional`
	Sl_TxInterestedFreqList_r16          *SL_TxInterestedFreqList_r16         `optional`
	Sl_CapabilityInformationSidelink_r16 *[]byte                              `optional`
}

func (ie *SL_TxResourceReq_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Sl_RLC_ModeIndicationList_r16) > 0, len(ie.Sl_QoS_InfoList_r16) > 0, len(ie.Sl_TypeTxSyncList_r16) > 0, ie.Sl_TxInterestedFreqList_r16 != nil, ie.Sl_CapabilityInformationSidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sl_DestinationIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_DestinationIdentity_r16", err)
	}
	if err = ie.Sl_CastType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_CastType_r16", err)
	}
	if len(ie.Sl_RLC_ModeIndicationList_r16) > 0 {
		tmp_Sl_RLC_ModeIndicationList_r16 := utils.NewSequence[*SL_RLC_ModeIndication_r16]([]*SL_RLC_ModeIndication_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		for _, i := range ie.Sl_RLC_ModeIndicationList_r16 {
			tmp_Sl_RLC_ModeIndicationList_r16.Value = append(tmp_Sl_RLC_ModeIndicationList_r16.Value, &i)
		}
		if err = tmp_Sl_RLC_ModeIndicationList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RLC_ModeIndicationList_r16", err)
		}
	}
	if len(ie.Sl_QoS_InfoList_r16) > 0 {
		tmp_Sl_QoS_InfoList_r16 := utils.NewSequence[*SL_QoS_Info_r16]([]*SL_QoS_Info_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_QFIsPerDest_r16}, false)
		for _, i := range ie.Sl_QoS_InfoList_r16 {
			tmp_Sl_QoS_InfoList_r16.Value = append(tmp_Sl_QoS_InfoList_r16.Value, &i)
		}
		if err = tmp_Sl_QoS_InfoList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_QoS_InfoList_r16", err)
		}
	}
	if len(ie.Sl_TypeTxSyncList_r16) > 0 {
		tmp_Sl_TypeTxSyncList_r16 := utils.NewSequence[*SL_TypeTxSync_r16]([]*SL_TypeTxSync_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		for _, i := range ie.Sl_TypeTxSyncList_r16 {
			tmp_Sl_TypeTxSyncList_r16.Value = append(tmp_Sl_TypeTxSyncList_r16.Value, &i)
		}
		if err = tmp_Sl_TypeTxSyncList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TypeTxSyncList_r16", err)
		}
	}
	if ie.Sl_TxInterestedFreqList_r16 != nil {
		if err = ie.Sl_TxInterestedFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TxInterestedFreqList_r16", err)
		}
	}
	if ie.Sl_CapabilityInformationSidelink_r16 != nil {
		if err = w.WriteOctetString(*ie.Sl_CapabilityInformationSidelink_r16, nil, false); err != nil {
			return utils.WrapError("Encode Sl_CapabilityInformationSidelink_r16", err)
		}
	}
	return nil
}

func (ie *SL_TxResourceReq_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_RLC_ModeIndicationList_r16Present bool
	if Sl_RLC_ModeIndicationList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_QoS_InfoList_r16Present bool
	if Sl_QoS_InfoList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TypeTxSyncList_r16Present bool
	if Sl_TypeTxSyncList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TxInterestedFreqList_r16Present bool
	if Sl_TxInterestedFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CapabilityInformationSidelink_r16Present bool
	if Sl_CapabilityInformationSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sl_DestinationIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_DestinationIdentity_r16", err)
	}
	if err = ie.Sl_CastType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_CastType_r16", err)
	}
	if Sl_RLC_ModeIndicationList_r16Present {
		tmp_Sl_RLC_ModeIndicationList_r16 := utils.NewSequence[*SL_RLC_ModeIndication_r16]([]*SL_RLC_ModeIndication_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		fn_Sl_RLC_ModeIndicationList_r16 := func() *SL_RLC_ModeIndication_r16 {
			return new(SL_RLC_ModeIndication_r16)
		}
		if err = tmp_Sl_RLC_ModeIndicationList_r16.Decode(r, fn_Sl_RLC_ModeIndicationList_r16); err != nil {
			return utils.WrapError("Decode Sl_RLC_ModeIndicationList_r16", err)
		}
		ie.Sl_RLC_ModeIndicationList_r16 = []SL_RLC_ModeIndication_r16{}
		for _, i := range tmp_Sl_RLC_ModeIndicationList_r16.Value {
			ie.Sl_RLC_ModeIndicationList_r16 = append(ie.Sl_RLC_ModeIndicationList_r16, *i)
		}
	}
	if Sl_QoS_InfoList_r16Present {
		tmp_Sl_QoS_InfoList_r16 := utils.NewSequence[*SL_QoS_Info_r16]([]*SL_QoS_Info_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_QFIsPerDest_r16}, false)
		fn_Sl_QoS_InfoList_r16 := func() *SL_QoS_Info_r16 {
			return new(SL_QoS_Info_r16)
		}
		if err = tmp_Sl_QoS_InfoList_r16.Decode(r, fn_Sl_QoS_InfoList_r16); err != nil {
			return utils.WrapError("Decode Sl_QoS_InfoList_r16", err)
		}
		ie.Sl_QoS_InfoList_r16 = []SL_QoS_Info_r16{}
		for _, i := range tmp_Sl_QoS_InfoList_r16.Value {
			ie.Sl_QoS_InfoList_r16 = append(ie.Sl_QoS_InfoList_r16, *i)
		}
	}
	if Sl_TypeTxSyncList_r16Present {
		tmp_Sl_TypeTxSyncList_r16 := utils.NewSequence[*SL_TypeTxSync_r16]([]*SL_TypeTxSync_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		fn_Sl_TypeTxSyncList_r16 := func() *SL_TypeTxSync_r16 {
			return new(SL_TypeTxSync_r16)
		}
		if err = tmp_Sl_TypeTxSyncList_r16.Decode(r, fn_Sl_TypeTxSyncList_r16); err != nil {
			return utils.WrapError("Decode Sl_TypeTxSyncList_r16", err)
		}
		ie.Sl_TypeTxSyncList_r16 = []SL_TypeTxSync_r16{}
		for _, i := range tmp_Sl_TypeTxSyncList_r16.Value {
			ie.Sl_TypeTxSyncList_r16 = append(ie.Sl_TypeTxSyncList_r16, *i)
		}
	}
	if Sl_TxInterestedFreqList_r16Present {
		ie.Sl_TxInterestedFreqList_r16 = new(SL_TxInterestedFreqList_r16)
		if err = ie.Sl_TxInterestedFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TxInterestedFreqList_r16", err)
		}
	}
	if Sl_CapabilityInformationSidelink_r16Present {
		var tmp_os_Sl_CapabilityInformationSidelink_r16 []byte
		if tmp_os_Sl_CapabilityInformationSidelink_r16, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Sl_CapabilityInformationSidelink_r16", err)
		}
		ie.Sl_CapabilityInformationSidelink_r16 = &tmp_os_Sl_CapabilityInformationSidelink_r16
	}
	return nil
}
