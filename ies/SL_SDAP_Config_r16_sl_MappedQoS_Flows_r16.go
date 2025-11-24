package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_SDAP_Config_r16_sl_MappedQoS_Flows_r16_Choice_nothing uint64 = iota
	SL_SDAP_Config_r16_sl_MappedQoS_Flows_r16_Choice_Sl_MappedQoS_FlowsList_r16
	SL_SDAP_Config_r16_sl_MappedQoS_Flows_r16_Choice_Sl_MappedQoS_FlowsListDedicated_r16
)

type SL_SDAP_Config_r16_sl_MappedQoS_Flows_r16 struct {
	Choice                              uint64
	Sl_MappedQoS_FlowsList_r16          []SL_QoS_Profile_r16 `lb:1,ub:maxNrofSL_QFIs_r16,madatory`
	Sl_MappedQoS_FlowsListDedicated_r16 *SL_MappedQoS_FlowsListDedicated_r16
}

func (ie *SL_SDAP_Config_r16_sl_MappedQoS_Flows_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_SDAP_Config_r16_sl_MappedQoS_Flows_r16_Choice_Sl_MappedQoS_FlowsList_r16:
		tmp := utils.NewSequence[*SL_QoS_Profile_r16]([]*SL_QoS_Profile_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_QFIs_r16}, false)
		for _, i := range ie.Sl_MappedQoS_FlowsList_r16 {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode Sl_MappedQoS_FlowsList_r16", err)
		}
	case SL_SDAP_Config_r16_sl_MappedQoS_Flows_r16_Choice_Sl_MappedQoS_FlowsListDedicated_r16:
		if err = ie.Sl_MappedQoS_FlowsListDedicated_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Sl_MappedQoS_FlowsListDedicated_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_SDAP_Config_r16_sl_MappedQoS_Flows_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_SDAP_Config_r16_sl_MappedQoS_Flows_r16_Choice_Sl_MappedQoS_FlowsList_r16:
		tmp := utils.NewSequence[*SL_QoS_Profile_r16]([]*SL_QoS_Profile_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_QFIs_r16}, false)
		fn := func() *SL_QoS_Profile_r16 {
			return new(SL_QoS_Profile_r16)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode Sl_MappedQoS_FlowsList_r16", err)
		}
		ie.Sl_MappedQoS_FlowsList_r16 = []SL_QoS_Profile_r16{}
		for _, i := range tmp.Value {
			ie.Sl_MappedQoS_FlowsList_r16 = append(ie.Sl_MappedQoS_FlowsList_r16, *i)
		}
	case SL_SDAP_Config_r16_sl_MappedQoS_Flows_r16_Choice_Sl_MappedQoS_FlowsListDedicated_r16:
		ie.Sl_MappedQoS_FlowsListDedicated_r16 = new(SL_MappedQoS_FlowsListDedicated_r16)
		if err = ie.Sl_MappedQoS_FlowsListDedicated_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MappedQoS_FlowsListDedicated_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
