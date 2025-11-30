package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB1_uac_BarringInfo struct {
	Uac_BarringForCommon                        *UAC_BarringPerCatList                                            `optional`
	Uac_BarringPerPLMN_List                     *UAC_BarringPerPLMN_List                                          `optional`
	Uac_BarringInfoSetList                      UAC_BarringInfoSetList                                            `madatory`
	Uac_AccessCategory1_SelectionAssistanceInfo *SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo `lb:2,ub:maxPLMN,optional`
}

func (ie *SIB1_uac_BarringInfo) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Uac_BarringForCommon != nil, ie.Uac_BarringPerPLMN_List != nil, ie.Uac_AccessCategory1_SelectionAssistanceInfo != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Uac_BarringForCommon != nil {
		if err = ie.Uac_BarringForCommon.Encode(w); err != nil {
			return utils.WrapError("Encode Uac_BarringForCommon", err)
		}
	}
	if ie.Uac_BarringPerPLMN_List != nil {
		if err = ie.Uac_BarringPerPLMN_List.Encode(w); err != nil {
			return utils.WrapError("Encode Uac_BarringPerPLMN_List", err)
		}
	}
	if err = ie.Uac_BarringInfoSetList.Encode(w); err != nil {
		return utils.WrapError("Encode Uac_BarringInfoSetList", err)
	}
	if ie.Uac_AccessCategory1_SelectionAssistanceInfo != nil {
		if err = ie.Uac_AccessCategory1_SelectionAssistanceInfo.Encode(w); err != nil {
			return utils.WrapError("Encode Uac_AccessCategory1_SelectionAssistanceInfo", err)
		}
	}
	return nil
}

func (ie *SIB1_uac_BarringInfo) Decode(r *aper.AperReader) error {
	var err error
	var Uac_BarringForCommonPresent bool
	if Uac_BarringForCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Uac_BarringPerPLMN_ListPresent bool
	if Uac_BarringPerPLMN_ListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Uac_AccessCategory1_SelectionAssistanceInfoPresent bool
	if Uac_AccessCategory1_SelectionAssistanceInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Uac_BarringForCommonPresent {
		ie.Uac_BarringForCommon = new(UAC_BarringPerCatList)
		if err = ie.Uac_BarringForCommon.Decode(r); err != nil {
			return utils.WrapError("Decode Uac_BarringForCommon", err)
		}
	}
	if Uac_BarringPerPLMN_ListPresent {
		ie.Uac_BarringPerPLMN_List = new(UAC_BarringPerPLMN_List)
		if err = ie.Uac_BarringPerPLMN_List.Decode(r); err != nil {
			return utils.WrapError("Decode Uac_BarringPerPLMN_List", err)
		}
	}
	if err = ie.Uac_BarringInfoSetList.Decode(r); err != nil {
		return utils.WrapError("Decode Uac_BarringInfoSetList", err)
	}
	if Uac_AccessCategory1_SelectionAssistanceInfoPresent {
		ie.Uac_AccessCategory1_SelectionAssistanceInfo = new(SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo)
		if err = ie.Uac_AccessCategory1_SelectionAssistanceInfo.Decode(r); err != nil {
			return utils.WrapError("Decode Uac_AccessCategory1_SelectionAssistanceInfo", err)
		}
	}
	return nil
}
