package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo_Choice_nothing uint64 = iota
	SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo_Choice_PlmnCommon
	SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo_Choice_IndividualPLMNList
)

type SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo struct {
	Choice             uint64
	PlmnCommon         *UAC_AccessCategory1_SelectionAssistanceInfo
	IndividualPLMNList []UAC_AccessCategory1_SelectionAssistanceInfo `lb:2,ub:maxPLMN,madatory`
}

func (ie *SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo_Choice_PlmnCommon:
		if err = ie.PlmnCommon.Encode(w); err != nil {
			err = utils.WrapError("Encode PlmnCommon", err)
		}
	case SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo_Choice_IndividualPLMNList:
		tmp := utils.NewSequence[*UAC_AccessCategory1_SelectionAssistanceInfo]([]*UAC_AccessCategory1_SelectionAssistanceInfo{}, aper.Constraint{Lb: 2, Ub: maxPLMN}, false)
		for _, i := range ie.IndividualPLMNList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode IndividualPLMNList", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo_Choice_PlmnCommon:
		ie.PlmnCommon = new(UAC_AccessCategory1_SelectionAssistanceInfo)
		if err = ie.PlmnCommon.Decode(r); err != nil {
			return utils.WrapError("Decode PlmnCommon", err)
		}
	case SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo_Choice_IndividualPLMNList:
		tmp := utils.NewSequence[*UAC_AccessCategory1_SelectionAssistanceInfo]([]*UAC_AccessCategory1_SelectionAssistanceInfo{}, aper.Constraint{Lb: 2, Ub: maxPLMN}, false)
		fn := func() *UAC_AccessCategory1_SelectionAssistanceInfo {
			return new(UAC_AccessCategory1_SelectionAssistanceInfo)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode IndividualPLMNList", err)
		}
		ie.IndividualPLMNList = []UAC_AccessCategory1_SelectionAssistanceInfo{}
		for _, i := range tmp.Value {
			ie.IndividualPLMNList = append(ie.IndividualPLMNList, *i)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
