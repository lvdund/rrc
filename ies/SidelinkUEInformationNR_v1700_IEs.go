package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SidelinkUEInformationNR_v1700_IEs struct {
	Sl_TxResourceReqList_v1700        *SL_TxResourceReqList_v1700                    `optional`
	Sl_RxDRX_ReportList_v1700         *SL_RxDRX_ReportList_v1700                     `optional`
	Sl_RxInterestedGC_BC_DestList_r17 *SL_RxInterestedGC_BC_DestList_r17             `optional`
	Sl_RxInterestedFreqListDisc_r17   *SL_InterestedFreqList_r16                     `optional`
	Sl_TxResourceReqListDisc_r17      *SL_TxResourceReqListDisc_r17                  `optional`
	Sl_TxResourceReqListCommRelay_r17 *SL_TxResourceReqListCommRelay_r17             `optional`
	Ue_Type_r17                       *SidelinkUEInformationNR_v1700_IEs_ue_Type_r17 `optional`
	Sl_SourceIdentityRemoteUE_r17     *SL_SourceIdentity_r17                         `optional`
	NonCriticalExtension              interface{}                                    `optional`
}

func (ie *SidelinkUEInformationNR_v1700_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_TxResourceReqList_v1700 != nil, ie.Sl_RxDRX_ReportList_v1700 != nil, ie.Sl_RxInterestedGC_BC_DestList_r17 != nil, ie.Sl_RxInterestedFreqListDisc_r17 != nil, ie.Sl_TxResourceReqListDisc_r17 != nil, ie.Sl_TxResourceReqListCommRelay_r17 != nil, ie.Ue_Type_r17 != nil, ie.Sl_SourceIdentityRemoteUE_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_TxResourceReqList_v1700 != nil {
		if err = ie.Sl_TxResourceReqList_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TxResourceReqList_v1700", err)
		}
	}
	if ie.Sl_RxDRX_ReportList_v1700 != nil {
		if err = ie.Sl_RxDRX_ReportList_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RxDRX_ReportList_v1700", err)
		}
	}
	if ie.Sl_RxInterestedGC_BC_DestList_r17 != nil {
		if err = ie.Sl_RxInterestedGC_BC_DestList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RxInterestedGC_BC_DestList_r17", err)
		}
	}
	if ie.Sl_RxInterestedFreqListDisc_r17 != nil {
		if err = ie.Sl_RxInterestedFreqListDisc_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RxInterestedFreqListDisc_r17", err)
		}
	}
	if ie.Sl_TxResourceReqListDisc_r17 != nil {
		if err = ie.Sl_TxResourceReqListDisc_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TxResourceReqListDisc_r17", err)
		}
	}
	if ie.Sl_TxResourceReqListCommRelay_r17 != nil {
		if err = ie.Sl_TxResourceReqListCommRelay_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TxResourceReqListCommRelay_r17", err)
		}
	}
	if ie.Ue_Type_r17 != nil {
		if err = ie.Ue_Type_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_Type_r17", err)
		}
	}
	if ie.Sl_SourceIdentityRemoteUE_r17 != nil {
		if err = ie.Sl_SourceIdentityRemoteUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SourceIdentityRemoteUE_r17", err)
		}
	}
	return nil
}

func (ie *SidelinkUEInformationNR_v1700_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Sl_TxResourceReqList_v1700Present bool
	if Sl_TxResourceReqList_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RxDRX_ReportList_v1700Present bool
	if Sl_RxDRX_ReportList_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RxInterestedGC_BC_DestList_r17Present bool
	if Sl_RxInterestedGC_BC_DestList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RxInterestedFreqListDisc_r17Present bool
	if Sl_RxInterestedFreqListDisc_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TxResourceReqListDisc_r17Present bool
	if Sl_TxResourceReqListDisc_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TxResourceReqListCommRelay_r17Present bool
	if Sl_TxResourceReqListCommRelay_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ue_Type_r17Present bool
	if Ue_Type_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SourceIdentityRemoteUE_r17Present bool
	if Sl_SourceIdentityRemoteUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_TxResourceReqList_v1700Present {
		ie.Sl_TxResourceReqList_v1700 = new(SL_TxResourceReqList_v1700)
		if err = ie.Sl_TxResourceReqList_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TxResourceReqList_v1700", err)
		}
	}
	if Sl_RxDRX_ReportList_v1700Present {
		ie.Sl_RxDRX_ReportList_v1700 = new(SL_RxDRX_ReportList_v1700)
		if err = ie.Sl_RxDRX_ReportList_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_RxDRX_ReportList_v1700", err)
		}
	}
	if Sl_RxInterestedGC_BC_DestList_r17Present {
		ie.Sl_RxInterestedGC_BC_DestList_r17 = new(SL_RxInterestedGC_BC_DestList_r17)
		if err = ie.Sl_RxInterestedGC_BC_DestList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_RxInterestedGC_BC_DestList_r17", err)
		}
	}
	if Sl_RxInterestedFreqListDisc_r17Present {
		ie.Sl_RxInterestedFreqListDisc_r17 = new(SL_InterestedFreqList_r16)
		if err = ie.Sl_RxInterestedFreqListDisc_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_RxInterestedFreqListDisc_r17", err)
		}
	}
	if Sl_TxResourceReqListDisc_r17Present {
		ie.Sl_TxResourceReqListDisc_r17 = new(SL_TxResourceReqListDisc_r17)
		if err = ie.Sl_TxResourceReqListDisc_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TxResourceReqListDisc_r17", err)
		}
	}
	if Sl_TxResourceReqListCommRelay_r17Present {
		ie.Sl_TxResourceReqListCommRelay_r17 = new(SL_TxResourceReqListCommRelay_r17)
		if err = ie.Sl_TxResourceReqListCommRelay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TxResourceReqListCommRelay_r17", err)
		}
	}
	if Ue_Type_r17Present {
		ie.Ue_Type_r17 = new(SidelinkUEInformationNR_v1700_IEs_ue_Type_r17)
		if err = ie.Ue_Type_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_Type_r17", err)
		}
	}
	if Sl_SourceIdentityRemoteUE_r17Present {
		ie.Sl_SourceIdentityRemoteUE_r17 = new(SL_SourceIdentity_r17)
		if err = ie.Sl_SourceIdentityRemoteUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SourceIdentityRemoteUE_r17", err)
		}
	}
	return nil
}
