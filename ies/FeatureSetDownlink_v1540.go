package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1540 struct {
	OneFL_DMRS_TwoAdditionalDMRS_DL         *FeatureSetDownlink_v1540_oneFL_DMRS_TwoAdditionalDMRS_DL         `optional`
	AdditionalDMRS_DL_Alt                   *FeatureSetDownlink_v1540_additionalDMRS_DL_Alt                   `optional`
	TwoFL_DMRS_TwoAdditionalDMRS_DL         *FeatureSetDownlink_v1540_twoFL_DMRS_TwoAdditionalDMRS_DL         `optional`
	OneFL_DMRS_ThreeAdditionalDMRS_DL       *FeatureSetDownlink_v1540_oneFL_DMRS_ThreeAdditionalDMRS_DL       `optional`
	Pdcch_MonitoringAnyOccasionsWithSpanGap *FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap `optional`
	Pdsch_SeparationWithGap                 *FeatureSetDownlink_v1540_pdsch_SeparationWithGap                 `optional`
	Pdsch_ProcessingType2                   *FeatureSetDownlink_v1540_pdsch_ProcessingType2                   `optional`
	Pdsch_ProcessingType2_Limited           *FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited           `optional`
	Dl_MCS_TableAlt_DynamicIndication       *FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication       `optional`
}

func (ie *FeatureSetDownlink_v1540) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.OneFL_DMRS_TwoAdditionalDMRS_DL != nil, ie.AdditionalDMRS_DL_Alt != nil, ie.TwoFL_DMRS_TwoAdditionalDMRS_DL != nil, ie.OneFL_DMRS_ThreeAdditionalDMRS_DL != nil, ie.Pdcch_MonitoringAnyOccasionsWithSpanGap != nil, ie.Pdsch_SeparationWithGap != nil, ie.Pdsch_ProcessingType2 != nil, ie.Pdsch_ProcessingType2_Limited != nil, ie.Dl_MCS_TableAlt_DynamicIndication != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.OneFL_DMRS_TwoAdditionalDMRS_DL != nil {
		if err = ie.OneFL_DMRS_TwoAdditionalDMRS_DL.Encode(w); err != nil {
			return utils.WrapError("Encode OneFL_DMRS_TwoAdditionalDMRS_DL", err)
		}
	}
	if ie.AdditionalDMRS_DL_Alt != nil {
		if err = ie.AdditionalDMRS_DL_Alt.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalDMRS_DL_Alt", err)
		}
	}
	if ie.TwoFL_DMRS_TwoAdditionalDMRS_DL != nil {
		if err = ie.TwoFL_DMRS_TwoAdditionalDMRS_DL.Encode(w); err != nil {
			return utils.WrapError("Encode TwoFL_DMRS_TwoAdditionalDMRS_DL", err)
		}
	}
	if ie.OneFL_DMRS_ThreeAdditionalDMRS_DL != nil {
		if err = ie.OneFL_DMRS_ThreeAdditionalDMRS_DL.Encode(w); err != nil {
			return utils.WrapError("Encode OneFL_DMRS_ThreeAdditionalDMRS_DL", err)
		}
	}
	if ie.Pdcch_MonitoringAnyOccasionsWithSpanGap != nil {
		if err = ie.Pdcch_MonitoringAnyOccasionsWithSpanGap.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_MonitoringAnyOccasionsWithSpanGap", err)
		}
	}
	if ie.Pdsch_SeparationWithGap != nil {
		if err = ie.Pdsch_SeparationWithGap.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_SeparationWithGap", err)
		}
	}
	if ie.Pdsch_ProcessingType2 != nil {
		if err = ie.Pdsch_ProcessingType2.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_ProcessingType2", err)
		}
	}
	if ie.Pdsch_ProcessingType2_Limited != nil {
		if err = ie.Pdsch_ProcessingType2_Limited.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_ProcessingType2_Limited", err)
		}
	}
	if ie.Dl_MCS_TableAlt_DynamicIndication != nil {
		if err = ie.Dl_MCS_TableAlt_DynamicIndication.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_MCS_TableAlt_DynamicIndication", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1540) Decode(r *uper.UperReader) error {
	var err error
	var OneFL_DMRS_TwoAdditionalDMRS_DLPresent bool
	if OneFL_DMRS_TwoAdditionalDMRS_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AdditionalDMRS_DL_AltPresent bool
	if AdditionalDMRS_DL_AltPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoFL_DMRS_TwoAdditionalDMRS_DLPresent bool
	if TwoFL_DMRS_TwoAdditionalDMRS_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var OneFL_DMRS_ThreeAdditionalDMRS_DLPresent bool
	if OneFL_DMRS_ThreeAdditionalDMRS_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_MonitoringAnyOccasionsWithSpanGapPresent bool
	if Pdcch_MonitoringAnyOccasionsWithSpanGapPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_SeparationWithGapPresent bool
	if Pdsch_SeparationWithGapPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_ProcessingType2Present bool
	if Pdsch_ProcessingType2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_ProcessingType2_LimitedPresent bool
	if Pdsch_ProcessingType2_LimitedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dl_MCS_TableAlt_DynamicIndicationPresent bool
	if Dl_MCS_TableAlt_DynamicIndicationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if OneFL_DMRS_TwoAdditionalDMRS_DLPresent {
		ie.OneFL_DMRS_TwoAdditionalDMRS_DL = new(FeatureSetDownlink_v1540_oneFL_DMRS_TwoAdditionalDMRS_DL)
		if err = ie.OneFL_DMRS_TwoAdditionalDMRS_DL.Decode(r); err != nil {
			return utils.WrapError("Decode OneFL_DMRS_TwoAdditionalDMRS_DL", err)
		}
	}
	if AdditionalDMRS_DL_AltPresent {
		ie.AdditionalDMRS_DL_Alt = new(FeatureSetDownlink_v1540_additionalDMRS_DL_Alt)
		if err = ie.AdditionalDMRS_DL_Alt.Decode(r); err != nil {
			return utils.WrapError("Decode AdditionalDMRS_DL_Alt", err)
		}
	}
	if TwoFL_DMRS_TwoAdditionalDMRS_DLPresent {
		ie.TwoFL_DMRS_TwoAdditionalDMRS_DL = new(FeatureSetDownlink_v1540_twoFL_DMRS_TwoAdditionalDMRS_DL)
		if err = ie.TwoFL_DMRS_TwoAdditionalDMRS_DL.Decode(r); err != nil {
			return utils.WrapError("Decode TwoFL_DMRS_TwoAdditionalDMRS_DL", err)
		}
	}
	if OneFL_DMRS_ThreeAdditionalDMRS_DLPresent {
		ie.OneFL_DMRS_ThreeAdditionalDMRS_DL = new(FeatureSetDownlink_v1540_oneFL_DMRS_ThreeAdditionalDMRS_DL)
		if err = ie.OneFL_DMRS_ThreeAdditionalDMRS_DL.Decode(r); err != nil {
			return utils.WrapError("Decode OneFL_DMRS_ThreeAdditionalDMRS_DL", err)
		}
	}
	if Pdcch_MonitoringAnyOccasionsWithSpanGapPresent {
		ie.Pdcch_MonitoringAnyOccasionsWithSpanGap = new(FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap)
		if err = ie.Pdcch_MonitoringAnyOccasionsWithSpanGap.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_MonitoringAnyOccasionsWithSpanGap", err)
		}
	}
	if Pdsch_SeparationWithGapPresent {
		ie.Pdsch_SeparationWithGap = new(FeatureSetDownlink_v1540_pdsch_SeparationWithGap)
		if err = ie.Pdsch_SeparationWithGap.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_SeparationWithGap", err)
		}
	}
	if Pdsch_ProcessingType2Present {
		ie.Pdsch_ProcessingType2 = new(FeatureSetDownlink_v1540_pdsch_ProcessingType2)
		if err = ie.Pdsch_ProcessingType2.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_ProcessingType2", err)
		}
	}
	if Pdsch_ProcessingType2_LimitedPresent {
		ie.Pdsch_ProcessingType2_Limited = new(FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited)
		if err = ie.Pdsch_ProcessingType2_Limited.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_ProcessingType2_Limited", err)
		}
	}
	if Dl_MCS_TableAlt_DynamicIndicationPresent {
		ie.Dl_MCS_TableAlt_DynamicIndication = new(FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication)
		if err = ie.Dl_MCS_TableAlt_DynamicIndication.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_MCS_TableAlt_DynamicIndication", err)
		}
	}
	return nil
}
