package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PBPS_CPS_Config_r17 struct {
	Sl_AllowedResourceSelectionConfig_r17 *SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17 `optional`
	Sl_MinNumCandidateSlotsPeriodic_r17   *int64                                                        `lb:1,ub:32,optional`
	Sl_PBPS_OccasionReservePeriodList_r17 []int64                                                       `lb:1,ub:16,e_lb:1,e_ub:16,optional`
	Sl_Additional_PBPS_Occasion_r17       *SL_PBPS_CPS_Config_r17_sl_Additional_PBPS_Occasion_r17       `optional`
	Sl_CPS_WindowPeriodic_r17             *int64                                                        `lb:5,ub:30,optional`
	Sl_MinNumCandidateSlotsAperiodic_r17  *int64                                                        `lb:1,ub:32,optional`
	Sl_MinNumRssiMeasurementSlots_r17     *int64                                                        `lb:1,ub:800,optional`
	Sl_DefaultCBR_RandomSelection_r17     *int64                                                        `lb:0,ub:100,optional`
	Sl_DefaultCBR_PartialSensing_r17      *int64                                                        `lb:0,ub:100,optional`
	Sl_CPS_WindowAperiodic_r17            *int64                                                        `lb:0,ub:30,optional`
	Sl_PartialSensingInactiveTime_r17     *SL_PBPS_CPS_Config_r17_sl_PartialSensingInactiveTime_r17     `optional`
}

func (ie *SL_PBPS_CPS_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_AllowedResourceSelectionConfig_r17 != nil, ie.Sl_MinNumCandidateSlotsPeriodic_r17 != nil, len(ie.Sl_PBPS_OccasionReservePeriodList_r17) > 0, ie.Sl_Additional_PBPS_Occasion_r17 != nil, ie.Sl_CPS_WindowPeriodic_r17 != nil, ie.Sl_MinNumCandidateSlotsAperiodic_r17 != nil, ie.Sl_MinNumRssiMeasurementSlots_r17 != nil, ie.Sl_DefaultCBR_RandomSelection_r17 != nil, ie.Sl_DefaultCBR_PartialSensing_r17 != nil, ie.Sl_CPS_WindowAperiodic_r17 != nil, ie.Sl_PartialSensingInactiveTime_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_AllowedResourceSelectionConfig_r17 != nil {
		if err = ie.Sl_AllowedResourceSelectionConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_AllowedResourceSelectionConfig_r17", err)
		}
	}
	if ie.Sl_MinNumCandidateSlotsPeriodic_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_MinNumCandidateSlotsPeriodic_r17, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode Sl_MinNumCandidateSlotsPeriodic_r17", err)
		}
	}
	if len(ie.Sl_PBPS_OccasionReservePeriodList_r17) > 0 {
		tmp_Sl_PBPS_OccasionReservePeriodList_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 16}, false)
		for _, i := range ie.Sl_PBPS_OccasionReservePeriodList_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 1, Ub: 16}, false)
			tmp_Sl_PBPS_OccasionReservePeriodList_r17.Value = append(tmp_Sl_PBPS_OccasionReservePeriodList_r17.Value, &tmp_ie)
		}
		if err = tmp_Sl_PBPS_OccasionReservePeriodList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PBPS_OccasionReservePeriodList_r17", err)
		}
	}
	if ie.Sl_Additional_PBPS_Occasion_r17 != nil {
		if err = ie.Sl_Additional_PBPS_Occasion_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_Additional_PBPS_Occasion_r17", err)
		}
	}
	if ie.Sl_CPS_WindowPeriodic_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_CPS_WindowPeriodic_r17, &uper.Constraint{Lb: 5, Ub: 30}, false); err != nil {
			return utils.WrapError("Encode Sl_CPS_WindowPeriodic_r17", err)
		}
	}
	if ie.Sl_MinNumCandidateSlotsAperiodic_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_MinNumCandidateSlotsAperiodic_r17, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode Sl_MinNumCandidateSlotsAperiodic_r17", err)
		}
	}
	if ie.Sl_MinNumRssiMeasurementSlots_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_MinNumRssiMeasurementSlots_r17, &uper.Constraint{Lb: 1, Ub: 800}, false); err != nil {
			return utils.WrapError("Encode Sl_MinNumRssiMeasurementSlots_r17", err)
		}
	}
	if ie.Sl_DefaultCBR_RandomSelection_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_DefaultCBR_RandomSelection_r17, &uper.Constraint{Lb: 0, Ub: 100}, false); err != nil {
			return utils.WrapError("Encode Sl_DefaultCBR_RandomSelection_r17", err)
		}
	}
	if ie.Sl_DefaultCBR_PartialSensing_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_DefaultCBR_PartialSensing_r17, &uper.Constraint{Lb: 0, Ub: 100}, false); err != nil {
			return utils.WrapError("Encode Sl_DefaultCBR_PartialSensing_r17", err)
		}
	}
	if ie.Sl_CPS_WindowAperiodic_r17 != nil {
		if err = w.WriteInteger(*ie.Sl_CPS_WindowAperiodic_r17, &uper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
			return utils.WrapError("Encode Sl_CPS_WindowAperiodic_r17", err)
		}
	}
	if ie.Sl_PartialSensingInactiveTime_r17 != nil {
		if err = ie.Sl_PartialSensingInactiveTime_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PartialSensingInactiveTime_r17", err)
		}
	}
	return nil
}

func (ie *SL_PBPS_CPS_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var Sl_AllowedResourceSelectionConfig_r17Present bool
	if Sl_AllowedResourceSelectionConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MinNumCandidateSlotsPeriodic_r17Present bool
	if Sl_MinNumCandidateSlotsPeriodic_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PBPS_OccasionReservePeriodList_r17Present bool
	if Sl_PBPS_OccasionReservePeriodList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_Additional_PBPS_Occasion_r17Present bool
	if Sl_Additional_PBPS_Occasion_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CPS_WindowPeriodic_r17Present bool
	if Sl_CPS_WindowPeriodic_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MinNumCandidateSlotsAperiodic_r17Present bool
	if Sl_MinNumCandidateSlotsAperiodic_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MinNumRssiMeasurementSlots_r17Present bool
	if Sl_MinNumRssiMeasurementSlots_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_DefaultCBR_RandomSelection_r17Present bool
	if Sl_DefaultCBR_RandomSelection_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_DefaultCBR_PartialSensing_r17Present bool
	if Sl_DefaultCBR_PartialSensing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CPS_WindowAperiodic_r17Present bool
	if Sl_CPS_WindowAperiodic_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PartialSensingInactiveTime_r17Present bool
	if Sl_PartialSensingInactiveTime_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_AllowedResourceSelectionConfig_r17Present {
		ie.Sl_AllowedResourceSelectionConfig_r17 = new(SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17)
		if err = ie.Sl_AllowedResourceSelectionConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_AllowedResourceSelectionConfig_r17", err)
		}
	}
	if Sl_MinNumCandidateSlotsPeriodic_r17Present {
		var tmp_int_Sl_MinNumCandidateSlotsPeriodic_r17 int64
		if tmp_int_Sl_MinNumCandidateSlotsPeriodic_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode Sl_MinNumCandidateSlotsPeriodic_r17", err)
		}
		ie.Sl_MinNumCandidateSlotsPeriodic_r17 = &tmp_int_Sl_MinNumCandidateSlotsPeriodic_r17
	}
	if Sl_PBPS_OccasionReservePeriodList_r17Present {
		tmp_Sl_PBPS_OccasionReservePeriodList_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 16}, false)
		fn_Sl_PBPS_OccasionReservePeriodList_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 1, Ub: 16}, false)
			return &ie
		}
		if err = tmp_Sl_PBPS_OccasionReservePeriodList_r17.Decode(r, fn_Sl_PBPS_OccasionReservePeriodList_r17); err != nil {
			return utils.WrapError("Decode Sl_PBPS_OccasionReservePeriodList_r17", err)
		}
		ie.Sl_PBPS_OccasionReservePeriodList_r17 = []int64{}
		for _, i := range tmp_Sl_PBPS_OccasionReservePeriodList_r17.Value {
			ie.Sl_PBPS_OccasionReservePeriodList_r17 = append(ie.Sl_PBPS_OccasionReservePeriodList_r17, int64(i.Value))
		}
	}
	if Sl_Additional_PBPS_Occasion_r17Present {
		ie.Sl_Additional_PBPS_Occasion_r17 = new(SL_PBPS_CPS_Config_r17_sl_Additional_PBPS_Occasion_r17)
		if err = ie.Sl_Additional_PBPS_Occasion_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_Additional_PBPS_Occasion_r17", err)
		}
	}
	if Sl_CPS_WindowPeriodic_r17Present {
		var tmp_int_Sl_CPS_WindowPeriodic_r17 int64
		if tmp_int_Sl_CPS_WindowPeriodic_r17, err = r.ReadInteger(&uper.Constraint{Lb: 5, Ub: 30}, false); err != nil {
			return utils.WrapError("Decode Sl_CPS_WindowPeriodic_r17", err)
		}
		ie.Sl_CPS_WindowPeriodic_r17 = &tmp_int_Sl_CPS_WindowPeriodic_r17
	}
	if Sl_MinNumCandidateSlotsAperiodic_r17Present {
		var tmp_int_Sl_MinNumCandidateSlotsAperiodic_r17 int64
		if tmp_int_Sl_MinNumCandidateSlotsAperiodic_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode Sl_MinNumCandidateSlotsAperiodic_r17", err)
		}
		ie.Sl_MinNumCandidateSlotsAperiodic_r17 = &tmp_int_Sl_MinNumCandidateSlotsAperiodic_r17
	}
	if Sl_MinNumRssiMeasurementSlots_r17Present {
		var tmp_int_Sl_MinNumRssiMeasurementSlots_r17 int64
		if tmp_int_Sl_MinNumRssiMeasurementSlots_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 800}, false); err != nil {
			return utils.WrapError("Decode Sl_MinNumRssiMeasurementSlots_r17", err)
		}
		ie.Sl_MinNumRssiMeasurementSlots_r17 = &tmp_int_Sl_MinNumRssiMeasurementSlots_r17
	}
	if Sl_DefaultCBR_RandomSelection_r17Present {
		var tmp_int_Sl_DefaultCBR_RandomSelection_r17 int64
		if tmp_int_Sl_DefaultCBR_RandomSelection_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 100}, false); err != nil {
			return utils.WrapError("Decode Sl_DefaultCBR_RandomSelection_r17", err)
		}
		ie.Sl_DefaultCBR_RandomSelection_r17 = &tmp_int_Sl_DefaultCBR_RandomSelection_r17
	}
	if Sl_DefaultCBR_PartialSensing_r17Present {
		var tmp_int_Sl_DefaultCBR_PartialSensing_r17 int64
		if tmp_int_Sl_DefaultCBR_PartialSensing_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 100}, false); err != nil {
			return utils.WrapError("Decode Sl_DefaultCBR_PartialSensing_r17", err)
		}
		ie.Sl_DefaultCBR_PartialSensing_r17 = &tmp_int_Sl_DefaultCBR_PartialSensing_r17
	}
	if Sl_CPS_WindowAperiodic_r17Present {
		var tmp_int_Sl_CPS_WindowAperiodic_r17 int64
		if tmp_int_Sl_CPS_WindowAperiodic_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
			return utils.WrapError("Decode Sl_CPS_WindowAperiodic_r17", err)
		}
		ie.Sl_CPS_WindowAperiodic_r17 = &tmp_int_Sl_CPS_WindowAperiodic_r17
	}
	if Sl_PartialSensingInactiveTime_r17Present {
		ie.Sl_PartialSensingInactiveTime_r17 = new(SL_PBPS_CPS_Config_r17_sl_PartialSensingInactiveTime_r17)
		if err = ie.Sl_PartialSensingInactiveTime_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PartialSensingInactiveTime_r17", err)
		}
	}
	return nil
}
