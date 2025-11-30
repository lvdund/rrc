package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UEAssistanceInformation_v1700_IEs struct {
	Ul_GapFR2_Preference_r17             *UL_GapFR2_Preference_r17                                         `optional`
	Musim_Assistance_r17                 *MUSIM_Assistance_r17                                             `optional`
	OverheatingAssistance_r17            *OverheatingAssistance_r17                                        `optional`
	MaxBW_PreferenceFR2_2_r17            *MaxBW_PreferenceFR2_2_r17                                        `optional`
	MaxMIMO_LayerPreferenceFR2_2_r17     *MaxMIMO_LayerPreferenceFR2_2_r17                                 `optional`
	MinSchedulingOffsetPreferenceExt_r17 *MinSchedulingOffsetPreferenceExt_r17                             `optional`
	Rlm_MeasRelaxationState_r17          *bool                                                             `optional`
	Bfd_MeasRelaxationState_r17          *aper.BitString                                                   `lb:1,ub:maxNrofServingCells,optional`
	NonSDT_DataIndication_r17            *UEAssistanceInformation_v1700_IEs_nonSDT_DataIndication_r17      `optional`
	Scg_DeactivationPreference_r17       *UEAssistanceInformation_v1700_IEs_scg_DeactivationPreference_r17 `optional`
	UplinkData_r17                       *UEAssistanceInformation_v1700_IEs_uplinkData_r17                 `optional`
	Rrm_MeasRelaxationFulfilment_r17     *bool                                                             `optional`
	PropagationDelayDifference_r17       *PropagationDelayDifference_r17                                   `optional`
	NonCriticalExtension                 interface{}                                                       `optional`
}

func (ie *UEAssistanceInformation_v1700_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ul_GapFR2_Preference_r17 != nil, ie.Musim_Assistance_r17 != nil, ie.OverheatingAssistance_r17 != nil, ie.MaxBW_PreferenceFR2_2_r17 != nil, ie.MaxMIMO_LayerPreferenceFR2_2_r17 != nil, ie.MinSchedulingOffsetPreferenceExt_r17 != nil, ie.Rlm_MeasRelaxationState_r17 != nil, ie.Bfd_MeasRelaxationState_r17 != nil, ie.NonSDT_DataIndication_r17 != nil, ie.Scg_DeactivationPreference_r17 != nil, ie.UplinkData_r17 != nil, ie.Rrm_MeasRelaxationFulfilment_r17 != nil, ie.PropagationDelayDifference_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ul_GapFR2_Preference_r17 != nil {
		if err = ie.Ul_GapFR2_Preference_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_GapFR2_Preference_r17", err)
		}
	}
	if ie.Musim_Assistance_r17 != nil {
		if err = ie.Musim_Assistance_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Musim_Assistance_r17", err)
		}
	}
	if ie.OverheatingAssistance_r17 != nil {
		if err = ie.OverheatingAssistance_r17.Encode(w); err != nil {
			return utils.WrapError("Encode OverheatingAssistance_r17", err)
		}
	}
	if ie.MaxBW_PreferenceFR2_2_r17 != nil {
		if err = ie.MaxBW_PreferenceFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxBW_PreferenceFR2_2_r17", err)
		}
	}
	if ie.MaxMIMO_LayerPreferenceFR2_2_r17 != nil {
		if err = ie.MaxMIMO_LayerPreferenceFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxMIMO_LayerPreferenceFR2_2_r17", err)
		}
	}
	if ie.MinSchedulingOffsetPreferenceExt_r17 != nil {
		if err = ie.MinSchedulingOffsetPreferenceExt_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MinSchedulingOffsetPreferenceExt_r17", err)
		}
	}
	if ie.Rlm_MeasRelaxationState_r17 != nil {
		if err = w.WriteBoolean(*ie.Rlm_MeasRelaxationState_r17); err != nil {
			return utils.WrapError("Encode Rlm_MeasRelaxationState_r17", err)
		}
	}
	if ie.Bfd_MeasRelaxationState_r17 != nil {
		if err = w.WriteBitString(ie.Bfd_MeasRelaxationState_r17.Bytes, uint(ie.Bfd_MeasRelaxationState_r17.NumBits), &aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false); err != nil {
			return utils.WrapError("Encode Bfd_MeasRelaxationState_r17", err)
		}
	}
	if ie.NonSDT_DataIndication_r17 != nil {
		if err = ie.NonSDT_DataIndication_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NonSDT_DataIndication_r17", err)
		}
	}
	if ie.Scg_DeactivationPreference_r17 != nil {
		if err = ie.Scg_DeactivationPreference_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scg_DeactivationPreference_r17", err)
		}
	}
	if ie.UplinkData_r17 != nil {
		if err = ie.UplinkData_r17.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkData_r17", err)
		}
	}
	if ie.Rrm_MeasRelaxationFulfilment_r17 != nil {
		if err = w.WriteBoolean(*ie.Rrm_MeasRelaxationFulfilment_r17); err != nil {
			return utils.WrapError("Encode Rrm_MeasRelaxationFulfilment_r17", err)
		}
	}
	if ie.PropagationDelayDifference_r17 != nil {
		if err = ie.PropagationDelayDifference_r17.Encode(w); err != nil {
			return utils.WrapError("Encode PropagationDelayDifference_r17", err)
		}
	}
	return nil
}

func (ie *UEAssistanceInformation_v1700_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Ul_GapFR2_Preference_r17Present bool
	if Ul_GapFR2_Preference_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Musim_Assistance_r17Present bool
	if Musim_Assistance_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OverheatingAssistance_r17Present bool
	if OverheatingAssistance_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxBW_PreferenceFR2_2_r17Present bool
	if MaxBW_PreferenceFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxMIMO_LayerPreferenceFR2_2_r17Present bool
	if MaxMIMO_LayerPreferenceFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MinSchedulingOffsetPreferenceExt_r17Present bool
	if MinSchedulingOffsetPreferenceExt_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rlm_MeasRelaxationState_r17Present bool
	if Rlm_MeasRelaxationState_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Bfd_MeasRelaxationState_r17Present bool
	if Bfd_MeasRelaxationState_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonSDT_DataIndication_r17Present bool
	if NonSDT_DataIndication_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scg_DeactivationPreference_r17Present bool
	if Scg_DeactivationPreference_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var UplinkData_r17Present bool
	if UplinkData_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rrm_MeasRelaxationFulfilment_r17Present bool
	if Rrm_MeasRelaxationFulfilment_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PropagationDelayDifference_r17Present bool
	if PropagationDelayDifference_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ul_GapFR2_Preference_r17Present {
		ie.Ul_GapFR2_Preference_r17 = new(UL_GapFR2_Preference_r17)
		if err = ie.Ul_GapFR2_Preference_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_GapFR2_Preference_r17", err)
		}
	}
	if Musim_Assistance_r17Present {
		ie.Musim_Assistance_r17 = new(MUSIM_Assistance_r17)
		if err = ie.Musim_Assistance_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Musim_Assistance_r17", err)
		}
	}
	if OverheatingAssistance_r17Present {
		ie.OverheatingAssistance_r17 = new(OverheatingAssistance_r17)
		if err = ie.OverheatingAssistance_r17.Decode(r); err != nil {
			return utils.WrapError("Decode OverheatingAssistance_r17", err)
		}
	}
	if MaxBW_PreferenceFR2_2_r17Present {
		ie.MaxBW_PreferenceFR2_2_r17 = new(MaxBW_PreferenceFR2_2_r17)
		if err = ie.MaxBW_PreferenceFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxBW_PreferenceFR2_2_r17", err)
		}
	}
	if MaxMIMO_LayerPreferenceFR2_2_r17Present {
		ie.MaxMIMO_LayerPreferenceFR2_2_r17 = new(MaxMIMO_LayerPreferenceFR2_2_r17)
		if err = ie.MaxMIMO_LayerPreferenceFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxMIMO_LayerPreferenceFR2_2_r17", err)
		}
	}
	if MinSchedulingOffsetPreferenceExt_r17Present {
		ie.MinSchedulingOffsetPreferenceExt_r17 = new(MinSchedulingOffsetPreferenceExt_r17)
		if err = ie.MinSchedulingOffsetPreferenceExt_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MinSchedulingOffsetPreferenceExt_r17", err)
		}
	}
	if Rlm_MeasRelaxationState_r17Present {
		var tmp_bool_Rlm_MeasRelaxationState_r17 bool
		if tmp_bool_Rlm_MeasRelaxationState_r17, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode Rlm_MeasRelaxationState_r17", err)
		}
		ie.Rlm_MeasRelaxationState_r17 = &tmp_bool_Rlm_MeasRelaxationState_r17
	}
	if Bfd_MeasRelaxationState_r17Present {
		var tmp_bs_Bfd_MeasRelaxationState_r17 []byte
		var n_Bfd_MeasRelaxationState_r17 uint
		if tmp_bs_Bfd_MeasRelaxationState_r17, n_Bfd_MeasRelaxationState_r17, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false); err != nil {
			return utils.WrapError("Decode Bfd_MeasRelaxationState_r17", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_Bfd_MeasRelaxationState_r17,
			NumBits: uint64(n_Bfd_MeasRelaxationState_r17),
		}
		ie.Bfd_MeasRelaxationState_r17 = &tmp_bitstring
	}
	if NonSDT_DataIndication_r17Present {
		ie.NonSDT_DataIndication_r17 = new(UEAssistanceInformation_v1700_IEs_nonSDT_DataIndication_r17)
		if err = ie.NonSDT_DataIndication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NonSDT_DataIndication_r17", err)
		}
	}
	if Scg_DeactivationPreference_r17Present {
		ie.Scg_DeactivationPreference_r17 = new(UEAssistanceInformation_v1700_IEs_scg_DeactivationPreference_r17)
		if err = ie.Scg_DeactivationPreference_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scg_DeactivationPreference_r17", err)
		}
	}
	if UplinkData_r17Present {
		ie.UplinkData_r17 = new(UEAssistanceInformation_v1700_IEs_uplinkData_r17)
		if err = ie.UplinkData_r17.Decode(r); err != nil {
			return utils.WrapError("Decode UplinkData_r17", err)
		}
	}
	if Rrm_MeasRelaxationFulfilment_r17Present {
		var tmp_bool_Rrm_MeasRelaxationFulfilment_r17 bool
		if tmp_bool_Rrm_MeasRelaxationFulfilment_r17, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode Rrm_MeasRelaxationFulfilment_r17", err)
		}
		ie.Rrm_MeasRelaxationFulfilment_r17 = &tmp_bool_Rrm_MeasRelaxationFulfilment_r17
	}
	if PropagationDelayDifference_r17Present {
		ie.PropagationDelayDifference_r17 = new(PropagationDelayDifference_r17)
		if err = ie.PropagationDelayDifference_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PropagationDelayDifference_r17", err)
		}
	}
	return nil
}
