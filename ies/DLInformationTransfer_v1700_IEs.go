package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DLInformationTransfer_v1700_IEs struct {
	DedicatedInfoF1c_r17 *DedicatedInfoF1c_r17                             `optional`
	RxTxTimeDiff_gNB_r17 *RxTxTimeDiff_r17                                 `optional`
	Ta_PDC_r17           *DLInformationTransfer_v1700_IEs_ta_PDC_r17       `optional`
	Sib9Fallback_r17     *DLInformationTransfer_v1700_IEs_sib9Fallback_r17 `optional`
	NonCriticalExtension interface{}                                       `optional`
}

func (ie *DLInformationTransfer_v1700_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.DedicatedInfoF1c_r17 != nil, ie.RxTxTimeDiff_gNB_r17 != nil, ie.Ta_PDC_r17 != nil, ie.Sib9Fallback_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DedicatedInfoF1c_r17 != nil {
		if err = ie.DedicatedInfoF1c_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DedicatedInfoF1c_r17", err)
		}
	}
	if ie.RxTxTimeDiff_gNB_r17 != nil {
		if err = ie.RxTxTimeDiff_gNB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RxTxTimeDiff_gNB_r17", err)
		}
	}
	if ie.Ta_PDC_r17 != nil {
		if err = ie.Ta_PDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ta_PDC_r17", err)
		}
	}
	if ie.Sib9Fallback_r17 != nil {
		if err = ie.Sib9Fallback_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sib9Fallback_r17", err)
		}
	}
	return nil
}

func (ie *DLInformationTransfer_v1700_IEs) Decode(r *aper.AperReader) error {
	var err error
	var DedicatedInfoF1c_r17Present bool
	if DedicatedInfoF1c_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RxTxTimeDiff_gNB_r17Present bool
	if RxTxTimeDiff_gNB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ta_PDC_r17Present bool
	if Ta_PDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sib9Fallback_r17Present bool
	if Sib9Fallback_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if DedicatedInfoF1c_r17Present {
		ie.DedicatedInfoF1c_r17 = new(DedicatedInfoF1c_r17)
		if err = ie.DedicatedInfoF1c_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DedicatedInfoF1c_r17", err)
		}
	}
	if RxTxTimeDiff_gNB_r17Present {
		ie.RxTxTimeDiff_gNB_r17 = new(RxTxTimeDiff_r17)
		if err = ie.RxTxTimeDiff_gNB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RxTxTimeDiff_gNB_r17", err)
		}
	}
	if Ta_PDC_r17Present {
		ie.Ta_PDC_r17 = new(DLInformationTransfer_v1700_IEs_ta_PDC_r17)
		if err = ie.Ta_PDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ta_PDC_r17", err)
		}
	}
	if Sib9Fallback_r17Present {
		ie.Sib9Fallback_r17 = new(DLInformationTransfer_v1700_IEs_sib9Fallback_r17)
		if err = ie.Sib9Fallback_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sib9Fallback_r17", err)
		}
	}
	return nil
}
