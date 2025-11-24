package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1710 struct {
	MTRP_PUSCH_TypeA_CB_r17               *FeatureSetUplink_v1710_mTRP_PUSCH_TypeA_CB_r17               `optional`
	MTRP_PUSCH_RepetitionTypeA_r17        *FeatureSetUplink_v1710_mTRP_PUSCH_RepetitionTypeA_r17        `optional`
	MTRP_PUCCH_IntraSlot_r17              *FeatureSetUplink_v1710_mTRP_PUCCH_IntraSlot_r17              `optional`
	Srs_AntennaSwitching2SP_1Periodic_r17 *FeatureSetUplink_v1710_srs_AntennaSwitching2SP_1Periodic_r17 `optional`
	Srs_ExtensionAperiodicSRS_r17         *FeatureSetUplink_v1710_srs_ExtensionAperiodicSRS_r17         `optional`
	Srs_OneAP_SRS_r17                     *FeatureSetUplink_v1710_srs_OneAP_SRS_r17                     `optional`
	Ue_PowerClassPerBandPerBC_r17         *FeatureSetUplink_v1710_ue_PowerClassPerBandPerBC_r17         `optional`
	Tx_Support_UL_GapFR2_r17              *FeatureSetUplink_v1710_tx_Support_UL_GapFR2_r17              `optional`
}

func (ie *FeatureSetUplink_v1710) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MTRP_PUSCH_TypeA_CB_r17 != nil, ie.MTRP_PUSCH_RepetitionTypeA_r17 != nil, ie.MTRP_PUCCH_IntraSlot_r17 != nil, ie.Srs_AntennaSwitching2SP_1Periodic_r17 != nil, ie.Srs_ExtensionAperiodicSRS_r17 != nil, ie.Srs_OneAP_SRS_r17 != nil, ie.Ue_PowerClassPerBandPerBC_r17 != nil, ie.Tx_Support_UL_GapFR2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MTRP_PUSCH_TypeA_CB_r17 != nil {
		if err = ie.MTRP_PUSCH_TypeA_CB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MTRP_PUSCH_TypeA_CB_r17", err)
		}
	}
	if ie.MTRP_PUSCH_RepetitionTypeA_r17 != nil {
		if err = ie.MTRP_PUSCH_RepetitionTypeA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MTRP_PUSCH_RepetitionTypeA_r17", err)
		}
	}
	if ie.MTRP_PUCCH_IntraSlot_r17 != nil {
		if err = ie.MTRP_PUCCH_IntraSlot_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MTRP_PUCCH_IntraSlot_r17", err)
		}
	}
	if ie.Srs_AntennaSwitching2SP_1Periodic_r17 != nil {
		if err = ie.Srs_AntennaSwitching2SP_1Periodic_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_AntennaSwitching2SP_1Periodic_r17", err)
		}
	}
	if ie.Srs_ExtensionAperiodicSRS_r17 != nil {
		if err = ie.Srs_ExtensionAperiodicSRS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_ExtensionAperiodicSRS_r17", err)
		}
	}
	if ie.Srs_OneAP_SRS_r17 != nil {
		if err = ie.Srs_OneAP_SRS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_OneAP_SRS_r17", err)
		}
	}
	if ie.Ue_PowerClassPerBandPerBC_r17 != nil {
		if err = ie.Ue_PowerClassPerBandPerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_PowerClassPerBandPerBC_r17", err)
		}
	}
	if ie.Tx_Support_UL_GapFR2_r17 != nil {
		if err = ie.Tx_Support_UL_GapFR2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Tx_Support_UL_GapFR2_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1710) Decode(r *uper.UperReader) error {
	var err error
	var MTRP_PUSCH_TypeA_CB_r17Present bool
	if MTRP_PUSCH_TypeA_CB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MTRP_PUSCH_RepetitionTypeA_r17Present bool
	if MTRP_PUSCH_RepetitionTypeA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MTRP_PUCCH_IntraSlot_r17Present bool
	if MTRP_PUCCH_IntraSlot_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_AntennaSwitching2SP_1Periodic_r17Present bool
	if Srs_AntennaSwitching2SP_1Periodic_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_ExtensionAperiodicSRS_r17Present bool
	if Srs_ExtensionAperiodicSRS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_OneAP_SRS_r17Present bool
	if Srs_OneAP_SRS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ue_PowerClassPerBandPerBC_r17Present bool
	if Ue_PowerClassPerBandPerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Tx_Support_UL_GapFR2_r17Present bool
	if Tx_Support_UL_GapFR2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MTRP_PUSCH_TypeA_CB_r17Present {
		ie.MTRP_PUSCH_TypeA_CB_r17 = new(FeatureSetUplink_v1710_mTRP_PUSCH_TypeA_CB_r17)
		if err = ie.MTRP_PUSCH_TypeA_CB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MTRP_PUSCH_TypeA_CB_r17", err)
		}
	}
	if MTRP_PUSCH_RepetitionTypeA_r17Present {
		ie.MTRP_PUSCH_RepetitionTypeA_r17 = new(FeatureSetUplink_v1710_mTRP_PUSCH_RepetitionTypeA_r17)
		if err = ie.MTRP_PUSCH_RepetitionTypeA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MTRP_PUSCH_RepetitionTypeA_r17", err)
		}
	}
	if MTRP_PUCCH_IntraSlot_r17Present {
		ie.MTRP_PUCCH_IntraSlot_r17 = new(FeatureSetUplink_v1710_mTRP_PUCCH_IntraSlot_r17)
		if err = ie.MTRP_PUCCH_IntraSlot_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MTRP_PUCCH_IntraSlot_r17", err)
		}
	}
	if Srs_AntennaSwitching2SP_1Periodic_r17Present {
		ie.Srs_AntennaSwitching2SP_1Periodic_r17 = new(FeatureSetUplink_v1710_srs_AntennaSwitching2SP_1Periodic_r17)
		if err = ie.Srs_AntennaSwitching2SP_1Periodic_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_AntennaSwitching2SP_1Periodic_r17", err)
		}
	}
	if Srs_ExtensionAperiodicSRS_r17Present {
		ie.Srs_ExtensionAperiodicSRS_r17 = new(FeatureSetUplink_v1710_srs_ExtensionAperiodicSRS_r17)
		if err = ie.Srs_ExtensionAperiodicSRS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_ExtensionAperiodicSRS_r17", err)
		}
	}
	if Srs_OneAP_SRS_r17Present {
		ie.Srs_OneAP_SRS_r17 = new(FeatureSetUplink_v1710_srs_OneAP_SRS_r17)
		if err = ie.Srs_OneAP_SRS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_OneAP_SRS_r17", err)
		}
	}
	if Ue_PowerClassPerBandPerBC_r17Present {
		ie.Ue_PowerClassPerBandPerBC_r17 = new(FeatureSetUplink_v1710_ue_PowerClassPerBandPerBC_r17)
		if err = ie.Ue_PowerClassPerBandPerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_PowerClassPerBandPerBC_r17", err)
		}
	}
	if Tx_Support_UL_GapFR2_r17Present {
		ie.Tx_Support_UL_GapFR2_r17 = new(FeatureSetUplink_v1710_tx_Support_UL_GapFR2_r17)
		if err = ie.Tx_Support_UL_GapFR2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Tx_Support_UL_GapFR2_r17", err)
		}
	}
	return nil
}
