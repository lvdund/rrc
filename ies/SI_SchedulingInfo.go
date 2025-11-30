package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SI_SchedulingInfo struct {
	SchedulingInfoList      []SchedulingInfo                  `lb:1,ub:maxSI_Message,madatory`
	Si_WindowLength         SI_SchedulingInfo_si_WindowLength `madatory`
	Si_RequestConfig        *SI_RequestConfig                 `optional`
	Si_RequestConfigSUL     *SI_RequestConfig                 `optional`
	SystemInformationAreaID *aper.BitString                   `lb:24,ub:24,optional`
}

func (ie *SI_SchedulingInfo) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Si_RequestConfig != nil, ie.Si_RequestConfigSUL != nil, ie.SystemInformationAreaID != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_SchedulingInfoList := utils.NewSequence[*SchedulingInfo]([]*SchedulingInfo{}, aper.Constraint{Lb: 1, Ub: maxSI_Message}, false)
	for _, i := range ie.SchedulingInfoList {
		tmp_SchedulingInfoList.Value = append(tmp_SchedulingInfoList.Value, &i)
	}
	if err = tmp_SchedulingInfoList.Encode(w); err != nil {
		return utils.WrapError("Encode SchedulingInfoList", err)
	}
	if err = ie.Si_WindowLength.Encode(w); err != nil {
		return utils.WrapError("Encode Si_WindowLength", err)
	}
	if ie.Si_RequestConfig != nil {
		if err = ie.Si_RequestConfig.Encode(w); err != nil {
			return utils.WrapError("Encode Si_RequestConfig", err)
		}
	}
	if ie.Si_RequestConfigSUL != nil {
		if err = ie.Si_RequestConfigSUL.Encode(w); err != nil {
			return utils.WrapError("Encode Si_RequestConfigSUL", err)
		}
	}
	if ie.SystemInformationAreaID != nil {
		if err = w.WriteBitString(ie.SystemInformationAreaID.Bytes, uint(ie.SystemInformationAreaID.NumBits), &aper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
			return utils.WrapError("Encode SystemInformationAreaID", err)
		}
	}
	return nil
}

func (ie *SI_SchedulingInfo) Decode(r *aper.AperReader) error {
	var err error
	var Si_RequestConfigPresent bool
	if Si_RequestConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Si_RequestConfigSULPresent bool
	if Si_RequestConfigSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SystemInformationAreaIDPresent bool
	if SystemInformationAreaIDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_SchedulingInfoList := utils.NewSequence[*SchedulingInfo]([]*SchedulingInfo{}, aper.Constraint{Lb: 1, Ub: maxSI_Message}, false)
	fn_SchedulingInfoList := func() *SchedulingInfo {
		return new(SchedulingInfo)
	}
	if err = tmp_SchedulingInfoList.Decode(r, fn_SchedulingInfoList); err != nil {
		return utils.WrapError("Decode SchedulingInfoList", err)
	}
	ie.SchedulingInfoList = []SchedulingInfo{}
	for _, i := range tmp_SchedulingInfoList.Value {
		ie.SchedulingInfoList = append(ie.SchedulingInfoList, *i)
	}
	if err = ie.Si_WindowLength.Decode(r); err != nil {
		return utils.WrapError("Decode Si_WindowLength", err)
	}
	if Si_RequestConfigPresent {
		ie.Si_RequestConfig = new(SI_RequestConfig)
		if err = ie.Si_RequestConfig.Decode(r); err != nil {
			return utils.WrapError("Decode Si_RequestConfig", err)
		}
	}
	if Si_RequestConfigSULPresent {
		ie.Si_RequestConfigSUL = new(SI_RequestConfig)
		if err = ie.Si_RequestConfigSUL.Decode(r); err != nil {
			return utils.WrapError("Decode Si_RequestConfigSUL", err)
		}
	}
	if SystemInformationAreaIDPresent {
		var tmp_bs_SystemInformationAreaID []byte
		var n_SystemInformationAreaID uint
		if tmp_bs_SystemInformationAreaID, n_SystemInformationAreaID, err = r.ReadBitString(&aper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
			return utils.WrapError("Decode SystemInformationAreaID", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_SystemInformationAreaID,
			NumBits: uint64(n_SystemInformationAreaID),
		}
		ie.SystemInformationAreaID = &tmp_bitstring
	}
	return nil
}
