package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_TxResourceReqDisc_r17 struct {
	Sl_DestinationIdentityDisc_r17  SL_DestinationIdentity_r16                    `madatory`
	Sl_SourceIdentityRelayUE_r17    *SL_SourceIdentity_r17                        `optional`
	Sl_CastTypeDisc_r17             SL_TxResourceReqDisc_r17_sl_CastTypeDisc_r17  `madatory`
	Sl_TxInterestedFreqListDisc_r17 SL_TxInterestedFreqList_r16                   `madatory`
	Sl_TypeTxSyncListDisc_r17       []SL_TypeTxSync_r16                           `lb:1,ub:maxNrofFreqSL_r16,madatory`
	Sl_DiscoveryType_r17            SL_TxResourceReqDisc_r17_sl_DiscoveryType_r17 `madatory`
}

func (ie *SL_TxResourceReqDisc_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_SourceIdentityRelayUE_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sl_DestinationIdentityDisc_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_DestinationIdentityDisc_r17", err)
	}
	if ie.Sl_SourceIdentityRelayUE_r17 != nil {
		if err = ie.Sl_SourceIdentityRelayUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SourceIdentityRelayUE_r17", err)
		}
	}
	if err = ie.Sl_CastTypeDisc_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_CastTypeDisc_r17", err)
	}
	if err = ie.Sl_TxInterestedFreqListDisc_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_TxInterestedFreqListDisc_r17", err)
	}
	tmp_Sl_TypeTxSyncListDisc_r17 := utils.NewSequence[*SL_TypeTxSync_r16]([]*SL_TypeTxSync_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
	for _, i := range ie.Sl_TypeTxSyncListDisc_r17 {
		tmp_Sl_TypeTxSyncListDisc_r17.Value = append(tmp_Sl_TypeTxSyncListDisc_r17.Value, &i)
	}
	if err = tmp_Sl_TypeTxSyncListDisc_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_TypeTxSyncListDisc_r17", err)
	}
	if err = ie.Sl_DiscoveryType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_DiscoveryType_r17", err)
	}
	return nil
}

func (ie *SL_TxResourceReqDisc_r17) Decode(r *aper.AperReader) error {
	var err error
	var Sl_SourceIdentityRelayUE_r17Present bool
	if Sl_SourceIdentityRelayUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sl_DestinationIdentityDisc_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_DestinationIdentityDisc_r17", err)
	}
	if Sl_SourceIdentityRelayUE_r17Present {
		ie.Sl_SourceIdentityRelayUE_r17 = new(SL_SourceIdentity_r17)
		if err = ie.Sl_SourceIdentityRelayUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SourceIdentityRelayUE_r17", err)
		}
	}
	if err = ie.Sl_CastTypeDisc_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_CastTypeDisc_r17", err)
	}
	if err = ie.Sl_TxInterestedFreqListDisc_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_TxInterestedFreqListDisc_r17", err)
	}
	tmp_Sl_TypeTxSyncListDisc_r17 := utils.NewSequence[*SL_TypeTxSync_r16]([]*SL_TypeTxSync_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
	fn_Sl_TypeTxSyncListDisc_r17 := func() *SL_TypeTxSync_r16 {
		return new(SL_TypeTxSync_r16)
	}
	if err = tmp_Sl_TypeTxSyncListDisc_r17.Decode(r, fn_Sl_TypeTxSyncListDisc_r17); err != nil {
		return utils.WrapError("Decode Sl_TypeTxSyncListDisc_r17", err)
	}
	ie.Sl_TypeTxSyncListDisc_r17 = []SL_TypeTxSync_r16{}
	for _, i := range tmp_Sl_TypeTxSyncListDisc_r17.Value {
		ie.Sl_TypeTxSyncListDisc_r17 = append(ie.Sl_TypeTxSyncListDisc_r17, *i)
	}
	if err = ie.Sl_DiscoveryType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_DiscoveryType_r17", err)
	}
	return nil
}
