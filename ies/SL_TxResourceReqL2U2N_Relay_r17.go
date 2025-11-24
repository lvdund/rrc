package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_TxResourceReqL2U2N_Relay_r17 struct {
	Sl_DestinationIdentityL2U2N_r17      *SL_DestinationIdentity_r16                             `optional`
	Sl_TxInterestedFreqListL2U2N_r17     SL_TxInterestedFreqList_r16                             `madatory`
	Sl_TypeTxSyncListL2U2N_r17           []SL_TypeTxSync_r16                                     `lb:1,ub:maxNrofFreqSL_r16,madatory`
	Sl_LocalID_Request_r17               *SL_TxResourceReqL2U2N_Relay_r17_sl_LocalID_Request_r17 `optional`
	Sl_PagingIdentityRemoteUE_r17        *SL_PagingIdentityRemoteUE_r17                          `optional`
	Sl_CapabilityInformationSidelink_r17 *[]byte                                                 `optional`
}

func (ie *SL_TxResourceReqL2U2N_Relay_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_DestinationIdentityL2U2N_r17 != nil, ie.Sl_LocalID_Request_r17 != nil, ie.Sl_PagingIdentityRemoteUE_r17 != nil, ie.Sl_CapabilityInformationSidelink_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_DestinationIdentityL2U2N_r17 != nil {
		if err = ie.Sl_DestinationIdentityL2U2N_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DestinationIdentityL2U2N_r17", err)
		}
	}
	if err = ie.Sl_TxInterestedFreqListL2U2N_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_TxInterestedFreqListL2U2N_r17", err)
	}
	tmp_Sl_TypeTxSyncListL2U2N_r17 := utils.NewSequence[*SL_TypeTxSync_r16]([]*SL_TypeTxSync_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
	for _, i := range ie.Sl_TypeTxSyncListL2U2N_r17 {
		tmp_Sl_TypeTxSyncListL2U2N_r17.Value = append(tmp_Sl_TypeTxSyncListL2U2N_r17.Value, &i)
	}
	if err = tmp_Sl_TypeTxSyncListL2U2N_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_TypeTxSyncListL2U2N_r17", err)
	}
	if ie.Sl_LocalID_Request_r17 != nil {
		if err = ie.Sl_LocalID_Request_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_LocalID_Request_r17", err)
		}
	}
	if ie.Sl_PagingIdentityRemoteUE_r17 != nil {
		if err = ie.Sl_PagingIdentityRemoteUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PagingIdentityRemoteUE_r17", err)
		}
	}
	if ie.Sl_CapabilityInformationSidelink_r17 != nil {
		if err = w.WriteOctetString(*ie.Sl_CapabilityInformationSidelink_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode Sl_CapabilityInformationSidelink_r17", err)
		}
	}
	return nil
}

func (ie *SL_TxResourceReqL2U2N_Relay_r17) Decode(r *uper.UperReader) error {
	var err error
	var Sl_DestinationIdentityL2U2N_r17Present bool
	if Sl_DestinationIdentityL2U2N_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_LocalID_Request_r17Present bool
	if Sl_LocalID_Request_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PagingIdentityRemoteUE_r17Present bool
	if Sl_PagingIdentityRemoteUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CapabilityInformationSidelink_r17Present bool
	if Sl_CapabilityInformationSidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_DestinationIdentityL2U2N_r17Present {
		ie.Sl_DestinationIdentityL2U2N_r17 = new(SL_DestinationIdentity_r16)
		if err = ie.Sl_DestinationIdentityL2U2N_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_DestinationIdentityL2U2N_r17", err)
		}
	}
	if err = ie.Sl_TxInterestedFreqListL2U2N_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_TxInterestedFreqListL2U2N_r17", err)
	}
	tmp_Sl_TypeTxSyncListL2U2N_r17 := utils.NewSequence[*SL_TypeTxSync_r16]([]*SL_TypeTxSync_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
	fn_Sl_TypeTxSyncListL2U2N_r17 := func() *SL_TypeTxSync_r16 {
		return new(SL_TypeTxSync_r16)
	}
	if err = tmp_Sl_TypeTxSyncListL2U2N_r17.Decode(r, fn_Sl_TypeTxSyncListL2U2N_r17); err != nil {
		return utils.WrapError("Decode Sl_TypeTxSyncListL2U2N_r17", err)
	}
	ie.Sl_TypeTxSyncListL2U2N_r17 = []SL_TypeTxSync_r16{}
	for _, i := range tmp_Sl_TypeTxSyncListL2U2N_r17.Value {
		ie.Sl_TypeTxSyncListL2U2N_r17 = append(ie.Sl_TypeTxSyncListL2U2N_r17, *i)
	}
	if Sl_LocalID_Request_r17Present {
		ie.Sl_LocalID_Request_r17 = new(SL_TxResourceReqL2U2N_Relay_r17_sl_LocalID_Request_r17)
		if err = ie.Sl_LocalID_Request_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_LocalID_Request_r17", err)
		}
	}
	if Sl_PagingIdentityRemoteUE_r17Present {
		ie.Sl_PagingIdentityRemoteUE_r17 = new(SL_PagingIdentityRemoteUE_r17)
		if err = ie.Sl_PagingIdentityRemoteUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PagingIdentityRemoteUE_r17", err)
		}
	}
	if Sl_CapabilityInformationSidelink_r17Present {
		var tmp_os_Sl_CapabilityInformationSidelink_r17 []byte
		if tmp_os_Sl_CapabilityInformationSidelink_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Sl_CapabilityInformationSidelink_r17", err)
		}
		ie.Sl_CapabilityInformationSidelink_r17 = &tmp_os_Sl_CapabilityInformationSidelink_r17
	}
	return nil
}
