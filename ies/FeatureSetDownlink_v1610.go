package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1610 struct {
	CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16 *FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16 `optional`
	CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16 *FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16 `optional`
	IntraFreqDAPS_r16                                *FeatureSetDownlink_v1610_intraFreqDAPS_r16                                `optional`
	IntraBandFreqSeparationDL_v1620                  *FreqSeparationClassDL_v1620                                               `optional`
	IntraBandFreqSeparationDL_Only_r16               *FreqSeparationClassDL_Only_r16                                            `optional`
	Pdcch_Monitoring_r16                             *FeatureSetDownlink_v1610_pdcch_Monitoring_r16                             `optional`
	Pdcch_MonitoringMixed_r16                        *FeatureSetDownlink_v1610_pdcch_MonitoringMixed_r16                        `optional`
	CrossCarrierSchedulingProcessing_DiffSCS_r16     *CrossCarrierSchedulingProcessing_DiffSCS_r16                              `optional`
	SingleDCI_SDM_scheme_r16                         *FeatureSetDownlink_v1610_singleDCI_SDM_scheme_r16                         `optional`
}

func (ie *FeatureSetDownlink_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16 != nil, ie.CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16 != nil, ie.IntraFreqDAPS_r16 != nil, ie.IntraBandFreqSeparationDL_v1620 != nil, ie.IntraBandFreqSeparationDL_Only_r16 != nil, ie.Pdcch_Monitoring_r16 != nil, ie.Pdcch_MonitoringMixed_r16 != nil, ie.CrossCarrierSchedulingProcessing_DiffSCS_r16 != nil, ie.SingleDCI_SDM_scheme_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16 != nil {
		if err = ie.CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16", err)
		}
	}
	if ie.CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16 != nil {
		if err = ie.CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16", err)
		}
	}
	if ie.IntraFreqDAPS_r16 != nil {
		if err = ie.IntraFreqDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IntraFreqDAPS_r16", err)
		}
	}
	if ie.IntraBandFreqSeparationDL_v1620 != nil {
		if err = ie.IntraBandFreqSeparationDL_v1620.Encode(w); err != nil {
			return utils.WrapError("Encode IntraBandFreqSeparationDL_v1620", err)
		}
	}
	if ie.IntraBandFreqSeparationDL_Only_r16 != nil {
		if err = ie.IntraBandFreqSeparationDL_Only_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IntraBandFreqSeparationDL_Only_r16", err)
		}
	}
	if ie.Pdcch_Monitoring_r16 != nil {
		if err = ie.Pdcch_Monitoring_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_Monitoring_r16", err)
		}
	}
	if ie.Pdcch_MonitoringMixed_r16 != nil {
		if err = ie.Pdcch_MonitoringMixed_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_MonitoringMixed_r16", err)
		}
	}
	if ie.CrossCarrierSchedulingProcessing_DiffSCS_r16 != nil {
		if err = ie.CrossCarrierSchedulingProcessing_DiffSCS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CrossCarrierSchedulingProcessing_DiffSCS_r16", err)
		}
	}
	if ie.SingleDCI_SDM_scheme_r16 != nil {
		if err = ie.SingleDCI_SDM_scheme_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SingleDCI_SDM_scheme_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1610) Decode(r *uper.UperReader) error {
	var err error
	var CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16Present bool
	if CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16Present bool
	if CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraFreqDAPS_r16Present bool
	if IntraFreqDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraBandFreqSeparationDL_v1620Present bool
	if IntraBandFreqSeparationDL_v1620Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraBandFreqSeparationDL_Only_r16Present bool
	if IntraBandFreqSeparationDL_Only_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_Monitoring_r16Present bool
	if Pdcch_Monitoring_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_MonitoringMixed_r16Present bool
	if Pdcch_MonitoringMixed_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CrossCarrierSchedulingProcessing_DiffSCS_r16Present bool
	if CrossCarrierSchedulingProcessing_DiffSCS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SingleDCI_SDM_scheme_r16Present bool
	if SingleDCI_SDM_scheme_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16Present {
		ie.CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16 = new(FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16)
		if err = ie.CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16", err)
		}
	}
	if CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16Present {
		ie.CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16 = new(FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16)
		if err = ie.CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16", err)
		}
	}
	if IntraFreqDAPS_r16Present {
		ie.IntraFreqDAPS_r16 = new(FeatureSetDownlink_v1610_intraFreqDAPS_r16)
		if err = ie.IntraFreqDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IntraFreqDAPS_r16", err)
		}
	}
	if IntraBandFreqSeparationDL_v1620Present {
		ie.IntraBandFreqSeparationDL_v1620 = new(FreqSeparationClassDL_v1620)
		if err = ie.IntraBandFreqSeparationDL_v1620.Decode(r); err != nil {
			return utils.WrapError("Decode IntraBandFreqSeparationDL_v1620", err)
		}
	}
	if IntraBandFreqSeparationDL_Only_r16Present {
		ie.IntraBandFreqSeparationDL_Only_r16 = new(FreqSeparationClassDL_Only_r16)
		if err = ie.IntraBandFreqSeparationDL_Only_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IntraBandFreqSeparationDL_Only_r16", err)
		}
	}
	if Pdcch_Monitoring_r16Present {
		ie.Pdcch_Monitoring_r16 = new(FeatureSetDownlink_v1610_pdcch_Monitoring_r16)
		if err = ie.Pdcch_Monitoring_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_Monitoring_r16", err)
		}
	}
	if Pdcch_MonitoringMixed_r16Present {
		ie.Pdcch_MonitoringMixed_r16 = new(FeatureSetDownlink_v1610_pdcch_MonitoringMixed_r16)
		if err = ie.Pdcch_MonitoringMixed_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_MonitoringMixed_r16", err)
		}
	}
	if CrossCarrierSchedulingProcessing_DiffSCS_r16Present {
		ie.CrossCarrierSchedulingProcessing_DiffSCS_r16 = new(CrossCarrierSchedulingProcessing_DiffSCS_r16)
		if err = ie.CrossCarrierSchedulingProcessing_DiffSCS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CrossCarrierSchedulingProcessing_DiffSCS_r16", err)
		}
	}
	if SingleDCI_SDM_scheme_r16Present {
		ie.SingleDCI_SDM_scheme_r16 = new(FeatureSetDownlink_v1610_singleDCI_SDM_scheme_r16)
		if err = ie.SingleDCI_SDM_scheme_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SingleDCI_SDM_scheme_r16", err)
		}
	}
	return nil
}
