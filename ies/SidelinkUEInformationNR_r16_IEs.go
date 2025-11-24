package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SidelinkUEInformationNR_r16_IEs struct {
	Sl_RxInterestedFreqList_r16 *SL_InterestedFreqList_r16         `optional`
	Sl_TxResourceReqList_r16    *SL_TxResourceReqList_r16          `optional`
	Sl_FailureList_r16          *SL_FailureList_r16                `optional`
	LateNonCriticalExtension    *[]byte                            `optional`
	NonCriticalExtension        *SidelinkUEInformationNR_v1700_IEs `optional`
}

func (ie *SidelinkUEInformationNR_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_RxInterestedFreqList_r16 != nil, ie.Sl_TxResourceReqList_r16 != nil, ie.Sl_FailureList_r16 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_RxInterestedFreqList_r16 != nil {
		if err = ie.Sl_RxInterestedFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RxInterestedFreqList_r16", err)
		}
	}
	if ie.Sl_TxResourceReqList_r16 != nil {
		if err = ie.Sl_TxResourceReqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TxResourceReqList_r16", err)
		}
	}
	if ie.Sl_FailureList_r16 != nil {
		if err = ie.Sl_FailureList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_FailureList_r16", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SidelinkUEInformationNR_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Sl_RxInterestedFreqList_r16Present bool
	if Sl_RxInterestedFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TxResourceReqList_r16Present bool
	if Sl_TxResourceReqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_FailureList_r16Present bool
	if Sl_FailureList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_RxInterestedFreqList_r16Present {
		ie.Sl_RxInterestedFreqList_r16 = new(SL_InterestedFreqList_r16)
		if err = ie.Sl_RxInterestedFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_RxInterestedFreqList_r16", err)
		}
	}
	if Sl_TxResourceReqList_r16Present {
		ie.Sl_TxResourceReqList_r16 = new(SL_TxResourceReqList_r16)
		if err = ie.Sl_TxResourceReqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TxResourceReqList_r16", err)
		}
	}
	if Sl_FailureList_r16Present {
		ie.Sl_FailureList_r16 = new(SL_FailureList_r16)
		if err = ie.Sl_FailureList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_FailureList_r16", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(SidelinkUEInformationNR_v1700_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
