package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RemoteUEInformationSidelink_r17_IEs struct {
	Sl_RequestedSIB_List_r17   *SL_RequestedSIB_List_r17   `optional,setuprelease`
	Sl_PagingInfo_RemoteUE_r17 *SL_PagingInfo_RemoteUE_r17 `optional,setuprelease`
	LateNonCriticalExtension   *[]byte                     `optional`
	NonCriticalExtension       interface{}                 `optional`
}

func (ie *RemoteUEInformationSidelink_r17_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_RequestedSIB_List_r17 != nil, ie.Sl_PagingInfo_RemoteUE_r17 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_RequestedSIB_List_r17 != nil {
		tmp_Sl_RequestedSIB_List_r17 := utils.SetupRelease[*SL_RequestedSIB_List_r17]{
			Setup: ie.Sl_RequestedSIB_List_r17,
		}
		if err = tmp_Sl_RequestedSIB_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RequestedSIB_List_r17", err)
		}
	}
	if ie.Sl_PagingInfo_RemoteUE_r17 != nil {
		tmp_Sl_PagingInfo_RemoteUE_r17 := utils.SetupRelease[*SL_PagingInfo_RemoteUE_r17]{
			Setup: ie.Sl_PagingInfo_RemoteUE_r17,
		}
		if err = tmp_Sl_PagingInfo_RemoteUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PagingInfo_RemoteUE_r17", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RemoteUEInformationSidelink_r17_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Sl_RequestedSIB_List_r17Present bool
	if Sl_RequestedSIB_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PagingInfo_RemoteUE_r17Present bool
	if Sl_PagingInfo_RemoteUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_RequestedSIB_List_r17Present {
		tmp_Sl_RequestedSIB_List_r17 := utils.SetupRelease[*SL_RequestedSIB_List_r17]{}
		if err = tmp_Sl_RequestedSIB_List_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_RequestedSIB_List_r17", err)
		}
		ie.Sl_RequestedSIB_List_r17 = tmp_Sl_RequestedSIB_List_r17.Setup
	}
	if Sl_PagingInfo_RemoteUE_r17Present {
		tmp_Sl_PagingInfo_RemoteUE_r17 := utils.SetupRelease[*SL_PagingInfo_RemoteUE_r17]{}
		if err = tmp_Sl_PagingInfo_RemoteUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PagingInfo_RemoteUE_r17", err)
		}
		ie.Sl_PagingInfo_RemoteUE_r17 = tmp_Sl_PagingInfo_RemoteUE_r17.Setup
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
