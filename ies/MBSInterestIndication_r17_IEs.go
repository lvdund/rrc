package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MBSInterestIndication_r17_IEs struct {
	Mbs_FreqList_r17         *CarrierFreqListMBS_r17                         `optional`
	Mbs_Priority_r17         *MBSInterestIndication_r17_IEs_mbs_Priority_r17 `optional`
	Mbs_ServiceList_r17      *MBS_ServiceList_r17                            `optional`
	LateNonCriticalExtension *[]byte                                         `optional`
	NonCriticalExtension     interface{}                                     `optional`
}

func (ie *MBSInterestIndication_r17_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Mbs_FreqList_r17 != nil, ie.Mbs_Priority_r17 != nil, ie.Mbs_ServiceList_r17 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Mbs_FreqList_r17 != nil {
		if err = ie.Mbs_FreqList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mbs_FreqList_r17", err)
		}
	}
	if ie.Mbs_Priority_r17 != nil {
		if err = ie.Mbs_Priority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mbs_Priority_r17", err)
		}
	}
	if ie.Mbs_ServiceList_r17 != nil {
		if err = ie.Mbs_ServiceList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mbs_ServiceList_r17", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *MBSInterestIndication_r17_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Mbs_FreqList_r17Present bool
	if Mbs_FreqList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mbs_Priority_r17Present bool
	if Mbs_Priority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mbs_ServiceList_r17Present bool
	if Mbs_ServiceList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Mbs_FreqList_r17Present {
		ie.Mbs_FreqList_r17 = new(CarrierFreqListMBS_r17)
		if err = ie.Mbs_FreqList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mbs_FreqList_r17", err)
		}
	}
	if Mbs_Priority_r17Present {
		ie.Mbs_Priority_r17 = new(MBSInterestIndication_r17_IEs_mbs_Priority_r17)
		if err = ie.Mbs_Priority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mbs_Priority_r17", err)
		}
	}
	if Mbs_ServiceList_r17Present {
		ie.Mbs_ServiceList_r17 = new(MBS_ServiceList_r17)
		if err = ie.Mbs_ServiceList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mbs_ServiceList_r17", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
