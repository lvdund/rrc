package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UuMessageTransferSidelink_r17_IEs struct {
	Sl_PagingDelivery_r17            *[]byte     `optional`
	Sl_SIB1_Delivery_r17             *[]byte     `optional`
	Sl_SystemInformationDelivery_r17 *[]byte     `optional`
	LateNonCriticalExtension         *[]byte     `optional`
	NonCriticalExtension             interface{} `optional`
}

func (ie *UuMessageTransferSidelink_r17_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_PagingDelivery_r17 != nil, ie.Sl_SIB1_Delivery_r17 != nil, ie.Sl_SystemInformationDelivery_r17 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_PagingDelivery_r17 != nil {
		if err = w.WriteOctetString(*ie.Sl_PagingDelivery_r17, nil, false); err != nil {
			return utils.WrapError("Encode Sl_PagingDelivery_r17", err)
		}
	}
	if ie.Sl_SIB1_Delivery_r17 != nil {
		if err = w.WriteOctetString(*ie.Sl_SIB1_Delivery_r17, nil, false); err != nil {
			return utils.WrapError("Encode Sl_SIB1_Delivery_r17", err)
		}
	}
	if ie.Sl_SystemInformationDelivery_r17 != nil {
		if err = w.WriteOctetString(*ie.Sl_SystemInformationDelivery_r17, nil, false); err != nil {
			return utils.WrapError("Encode Sl_SystemInformationDelivery_r17", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UuMessageTransferSidelink_r17_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Sl_PagingDelivery_r17Present bool
	if Sl_PagingDelivery_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SIB1_Delivery_r17Present bool
	if Sl_SIB1_Delivery_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SystemInformationDelivery_r17Present bool
	if Sl_SystemInformationDelivery_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_PagingDelivery_r17Present {
		var tmp_os_Sl_PagingDelivery_r17 []byte
		if tmp_os_Sl_PagingDelivery_r17, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Sl_PagingDelivery_r17", err)
		}
		ie.Sl_PagingDelivery_r17 = &tmp_os_Sl_PagingDelivery_r17
	}
	if Sl_SIB1_Delivery_r17Present {
		var tmp_os_Sl_SIB1_Delivery_r17 []byte
		if tmp_os_Sl_SIB1_Delivery_r17, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Sl_SIB1_Delivery_r17", err)
		}
		ie.Sl_SIB1_Delivery_r17 = &tmp_os_Sl_SIB1_Delivery_r17
	}
	if Sl_SystemInformationDelivery_r17Present {
		var tmp_os_Sl_SystemInformationDelivery_r17 []byte
		if tmp_os_Sl_SystemInformationDelivery_r17, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Sl_SystemInformationDelivery_r17", err)
		}
		ie.Sl_SystemInformationDelivery_r17 = &tmp_os_Sl_SystemInformationDelivery_r17
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
