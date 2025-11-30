package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UEAssistanceInformationSidelink_r17_IEs struct {
	Sl_PreferredDRX_ConfigList_r17 []SL_DRX_ConfigUC_SemiStatic_r17 `lb:1,ub:maxNrofSL_RxInfoSet_r17,optional`
	LateNonCriticalExtension       *[]byte                          `optional`
	NonCriticalExtension           interface{}                      `optional`
}

func (ie *UEAssistanceInformationSidelink_r17_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Sl_PreferredDRX_ConfigList_r17) > 0, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Sl_PreferredDRX_ConfigList_r17) > 0 {
		tmp_Sl_PreferredDRX_ConfigList_r17 := utils.NewSequence[*SL_DRX_ConfigUC_SemiStatic_r17]([]*SL_DRX_ConfigUC_SemiStatic_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_RxInfoSet_r17}, false)
		for _, i := range ie.Sl_PreferredDRX_ConfigList_r17 {
			tmp_Sl_PreferredDRX_ConfigList_r17.Value = append(tmp_Sl_PreferredDRX_ConfigList_r17.Value, &i)
		}
		if err = tmp_Sl_PreferredDRX_ConfigList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PreferredDRX_ConfigList_r17", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UEAssistanceInformationSidelink_r17_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Sl_PreferredDRX_ConfigList_r17Present bool
	if Sl_PreferredDRX_ConfigList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_PreferredDRX_ConfigList_r17Present {
		tmp_Sl_PreferredDRX_ConfigList_r17 := utils.NewSequence[*SL_DRX_ConfigUC_SemiStatic_r17]([]*SL_DRX_ConfigUC_SemiStatic_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_RxInfoSet_r17}, false)
		fn_Sl_PreferredDRX_ConfigList_r17 := func() *SL_DRX_ConfigUC_SemiStatic_r17 {
			return new(SL_DRX_ConfigUC_SemiStatic_r17)
		}
		if err = tmp_Sl_PreferredDRX_ConfigList_r17.Decode(r, fn_Sl_PreferredDRX_ConfigList_r17); err != nil {
			return utils.WrapError("Decode Sl_PreferredDRX_ConfigList_r17", err)
		}
		ie.Sl_PreferredDRX_ConfigList_r17 = []SL_DRX_ConfigUC_SemiStatic_r17{}
		for _, i := range tmp_Sl_PreferredDRX_ConfigList_r17.Value {
			ie.Sl_PreferredDRX_ConfigList_r17 = append(ie.Sl_PreferredDRX_ConfigList_r17, *i)
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
