package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1610 struct {
	ParallelTxMsgA_SRS_PUCCH_PUSCH_r16       *CA_ParametersNR_v1610_parallelTxMsgA_SRS_PUCCH_PUSCH_r16    `optional`
	MsgA_SUL_r16                             *CA_ParametersNR_v1610_msgA_SUL_r16                          `optional`
	JointSearchSpaceSwitchAcrossCells_r16    *CA_ParametersNR_v1610_jointSearchSpaceSwitchAcrossCells_r16 `optional`
	Half_DuplexTDD_CA_SameSCS_r16            *CA_ParametersNR_v1610_half_DuplexTDD_CA_SameSCS_r16         `optional`
	ScellDormancyWithinActiveTime_r16        *CA_ParametersNR_v1610_scellDormancyWithinActiveTime_r16     `optional`
	ScellDormancyOutsideActiveTime_r16       *CA_ParametersNR_v1610_scellDormancyOutsideActiveTime_r16    `optional`
	CrossCarrierA_CSI_trigDiffSCS_r16        *CA_ParametersNR_v1610_crossCarrierA_CSI_trigDiffSCS_r16     `optional`
	DefaultQCL_CrossCarrierA_CSI_Trig_r16    *CA_ParametersNR_v1610_defaultQCL_CrossCarrierA_CSI_Trig_r16 `optional`
	InterCA_NonAlignedFrame_r16              *CA_ParametersNR_v1610_interCA_NonAlignedFrame_r16           `optional`
	Simul_SRS_Trans_BC_r16                   *CA_ParametersNR_v1610_simul_SRS_Trans_BC_r16                `optional`
	InterFreqDAPS_r16                        *CA_ParametersNR_v1610_interFreqDAPS_r16                     `optional`
	CodebookParametersPerBC_r16              *CodebookParameters_v1610                                    `optional`
	BlindDetectFactor_r16                    *int64                                                       `lb:1,ub:2,optional`
	Pdcch_MonitoringCA_r16                   *CA_ParametersNR_v1610_pdcch_MonitoringCA_r16                `lb:2,ub:16,optional`
	Pdcch_BlindDetectionCA_Mixed_r16         *CA_ParametersNR_v1610_pdcch_BlindDetectionCA_Mixed_r16      `lb:1,ub:15,optional`
	Pdcch_BlindDetectionMCG_UE_r16           *int64                                                       `lb:1,ub:14,optional`
	Pdcch_BlindDetectionSCG_UE_r16           *int64                                                       `lb:1,ub:14,optional`
	Pdcch_BlindDetectionMCG_UE_Mixed_r16     *CA_ParametersNR_v1610_pdcch_BlindDetectionMCG_UE_Mixed_r16  `lb:0,ub:15,optional`
	Pdcch_BlindDetectionSCG_UE_Mixed_r16     *CA_ParametersNR_v1610_pdcch_BlindDetectionSCG_UE_Mixed_r16  `lb:0,ub:15,optional`
	CrossCarrierSchedulingDL_DiffSCS_r16     *CA_ParametersNR_v1610_crossCarrierSchedulingDL_DiffSCS_r16  `optional`
	CrossCarrierSchedulingDefaultQCL_r16     *CA_ParametersNR_v1610_crossCarrierSchedulingDefaultQCL_r16  `optional`
	CrossCarrierSchedulingUL_DiffSCS_r16     *CA_ParametersNR_v1610_crossCarrierSchedulingUL_DiffSCS_r16  `optional`
	Simul_SRS_MIMO_Trans_BC_r16              *CA_ParametersNR_v1610_simul_SRS_MIMO_Trans_BC_r16           `optional`
	CodebookParametersAdditionPerBC_r16      *CodebookParametersAdditionPerBC_r16                         `optional`
	CodebookComboParametersAdditionPerBC_r16 *CodebookComboParametersAdditionPerBC_r16                    `optional`
}

func (ie *CA_ParametersNR_v1610) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_r16 != nil, ie.MsgA_SUL_r16 != nil, ie.JointSearchSpaceSwitchAcrossCells_r16 != nil, ie.Half_DuplexTDD_CA_SameSCS_r16 != nil, ie.ScellDormancyWithinActiveTime_r16 != nil, ie.ScellDormancyOutsideActiveTime_r16 != nil, ie.CrossCarrierA_CSI_trigDiffSCS_r16 != nil, ie.DefaultQCL_CrossCarrierA_CSI_Trig_r16 != nil, ie.InterCA_NonAlignedFrame_r16 != nil, ie.Simul_SRS_Trans_BC_r16 != nil, ie.InterFreqDAPS_r16 != nil, ie.CodebookParametersPerBC_r16 != nil, ie.BlindDetectFactor_r16 != nil, ie.Pdcch_MonitoringCA_r16 != nil, ie.Pdcch_BlindDetectionCA_Mixed_r16 != nil, ie.Pdcch_BlindDetectionMCG_UE_r16 != nil, ie.Pdcch_BlindDetectionSCG_UE_r16 != nil, ie.Pdcch_BlindDetectionMCG_UE_Mixed_r16 != nil, ie.Pdcch_BlindDetectionSCG_UE_Mixed_r16 != nil, ie.CrossCarrierSchedulingDL_DiffSCS_r16 != nil, ie.CrossCarrierSchedulingDefaultQCL_r16 != nil, ie.CrossCarrierSchedulingUL_DiffSCS_r16 != nil, ie.Simul_SRS_MIMO_Trans_BC_r16 != nil, ie.CodebookParametersAdditionPerBC_r16 != nil, ie.CodebookComboParametersAdditionPerBC_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_r16 != nil {
		if err = ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ParallelTxMsgA_SRS_PUCCH_PUSCH_r16", err)
		}
	}
	if ie.MsgA_SUL_r16 != nil {
		if err = ie.MsgA_SUL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_SUL_r16", err)
		}
	}
	if ie.JointSearchSpaceSwitchAcrossCells_r16 != nil {
		if err = ie.JointSearchSpaceSwitchAcrossCells_r16.Encode(w); err != nil {
			return utils.WrapError("Encode JointSearchSpaceSwitchAcrossCells_r16", err)
		}
	}
	if ie.Half_DuplexTDD_CA_SameSCS_r16 != nil {
		if err = ie.Half_DuplexTDD_CA_SameSCS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Half_DuplexTDD_CA_SameSCS_r16", err)
		}
	}
	if ie.ScellDormancyWithinActiveTime_r16 != nil {
		if err = ie.ScellDormancyWithinActiveTime_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ScellDormancyWithinActiveTime_r16", err)
		}
	}
	if ie.ScellDormancyOutsideActiveTime_r16 != nil {
		if err = ie.ScellDormancyOutsideActiveTime_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ScellDormancyOutsideActiveTime_r16", err)
		}
	}
	if ie.CrossCarrierA_CSI_trigDiffSCS_r16 != nil {
		if err = ie.CrossCarrierA_CSI_trigDiffSCS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CrossCarrierA_CSI_trigDiffSCS_r16", err)
		}
	}
	if ie.DefaultQCL_CrossCarrierA_CSI_Trig_r16 != nil {
		if err = ie.DefaultQCL_CrossCarrierA_CSI_Trig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DefaultQCL_CrossCarrierA_CSI_Trig_r16", err)
		}
	}
	if ie.InterCA_NonAlignedFrame_r16 != nil {
		if err = ie.InterCA_NonAlignedFrame_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterCA_NonAlignedFrame_r16", err)
		}
	}
	if ie.Simul_SRS_Trans_BC_r16 != nil {
		if err = ie.Simul_SRS_Trans_BC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Simul_SRS_Trans_BC_r16", err)
		}
	}
	if ie.InterFreqDAPS_r16 != nil {
		if err = ie.InterFreqDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqDAPS_r16", err)
		}
	}
	if ie.CodebookParametersPerBC_r16 != nil {
		if err = ie.CodebookParametersPerBC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CodebookParametersPerBC_r16", err)
		}
	}
	if ie.BlindDetectFactor_r16 != nil {
		if err = w.WriteInteger(*ie.BlindDetectFactor_r16, &aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode BlindDetectFactor_r16", err)
		}
	}
	if ie.Pdcch_MonitoringCA_r16 != nil {
		if err = ie.Pdcch_MonitoringCA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_MonitoringCA_r16", err)
		}
	}
	if ie.Pdcch_BlindDetectionCA_Mixed_r16 != nil {
		if err = ie.Pdcch_BlindDetectionCA_Mixed_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_BlindDetectionCA_Mixed_r16", err)
		}
	}
	if ie.Pdcch_BlindDetectionMCG_UE_r16 != nil {
		if err = w.WriteInteger(*ie.Pdcch_BlindDetectionMCG_UE_r16, &aper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
			return utils.WrapError("Encode Pdcch_BlindDetectionMCG_UE_r16", err)
		}
	}
	if ie.Pdcch_BlindDetectionSCG_UE_r16 != nil {
		if err = w.WriteInteger(*ie.Pdcch_BlindDetectionSCG_UE_r16, &aper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
			return utils.WrapError("Encode Pdcch_BlindDetectionSCG_UE_r16", err)
		}
	}
	if ie.Pdcch_BlindDetectionMCG_UE_Mixed_r16 != nil {
		if err = ie.Pdcch_BlindDetectionMCG_UE_Mixed_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_BlindDetectionMCG_UE_Mixed_r16", err)
		}
	}
	if ie.Pdcch_BlindDetectionSCG_UE_Mixed_r16 != nil {
		if err = ie.Pdcch_BlindDetectionSCG_UE_Mixed_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_BlindDetectionSCG_UE_Mixed_r16", err)
		}
	}
	if ie.CrossCarrierSchedulingDL_DiffSCS_r16 != nil {
		if err = ie.CrossCarrierSchedulingDL_DiffSCS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CrossCarrierSchedulingDL_DiffSCS_r16", err)
		}
	}
	if ie.CrossCarrierSchedulingDefaultQCL_r16 != nil {
		if err = ie.CrossCarrierSchedulingDefaultQCL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CrossCarrierSchedulingDefaultQCL_r16", err)
		}
	}
	if ie.CrossCarrierSchedulingUL_DiffSCS_r16 != nil {
		if err = ie.CrossCarrierSchedulingUL_DiffSCS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CrossCarrierSchedulingUL_DiffSCS_r16", err)
		}
	}
	if ie.Simul_SRS_MIMO_Trans_BC_r16 != nil {
		if err = ie.Simul_SRS_MIMO_Trans_BC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Simul_SRS_MIMO_Trans_BC_r16", err)
		}
	}
	if ie.CodebookParametersAdditionPerBC_r16 != nil {
		if err = ie.CodebookParametersAdditionPerBC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CodebookParametersAdditionPerBC_r16", err)
		}
	}
	if ie.CodebookComboParametersAdditionPerBC_r16 != nil {
		if err = ie.CodebookComboParametersAdditionPerBC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CodebookComboParametersAdditionPerBC_r16", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1610) Decode(r *aper.AperReader) error {
	var err error
	var ParallelTxMsgA_SRS_PUCCH_PUSCH_r16Present bool
	if ParallelTxMsgA_SRS_PUCCH_PUSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_SUL_r16Present bool
	if MsgA_SUL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var JointSearchSpaceSwitchAcrossCells_r16Present bool
	if JointSearchSpaceSwitchAcrossCells_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Half_DuplexTDD_CA_SameSCS_r16Present bool
	if Half_DuplexTDD_CA_SameSCS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ScellDormancyWithinActiveTime_r16Present bool
	if ScellDormancyWithinActiveTime_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ScellDormancyOutsideActiveTime_r16Present bool
	if ScellDormancyOutsideActiveTime_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CrossCarrierA_CSI_trigDiffSCS_r16Present bool
	if CrossCarrierA_CSI_trigDiffSCS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DefaultQCL_CrossCarrierA_CSI_Trig_r16Present bool
	if DefaultQCL_CrossCarrierA_CSI_Trig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterCA_NonAlignedFrame_r16Present bool
	if InterCA_NonAlignedFrame_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Simul_SRS_Trans_BC_r16Present bool
	if Simul_SRS_Trans_BC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterFreqDAPS_r16Present bool
	if InterFreqDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CodebookParametersPerBC_r16Present bool
	if CodebookParametersPerBC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BlindDetectFactor_r16Present bool
	if BlindDetectFactor_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_MonitoringCA_r16Present bool
	if Pdcch_MonitoringCA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_BlindDetectionCA_Mixed_r16Present bool
	if Pdcch_BlindDetectionCA_Mixed_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_BlindDetectionMCG_UE_r16Present bool
	if Pdcch_BlindDetectionMCG_UE_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_BlindDetectionSCG_UE_r16Present bool
	if Pdcch_BlindDetectionSCG_UE_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_BlindDetectionMCG_UE_Mixed_r16Present bool
	if Pdcch_BlindDetectionMCG_UE_Mixed_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_BlindDetectionSCG_UE_Mixed_r16Present bool
	if Pdcch_BlindDetectionSCG_UE_Mixed_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CrossCarrierSchedulingDL_DiffSCS_r16Present bool
	if CrossCarrierSchedulingDL_DiffSCS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CrossCarrierSchedulingDefaultQCL_r16Present bool
	if CrossCarrierSchedulingDefaultQCL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CrossCarrierSchedulingUL_DiffSCS_r16Present bool
	if CrossCarrierSchedulingUL_DiffSCS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Simul_SRS_MIMO_Trans_BC_r16Present bool
	if Simul_SRS_MIMO_Trans_BC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CodebookParametersAdditionPerBC_r16Present bool
	if CodebookParametersAdditionPerBC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CodebookComboParametersAdditionPerBC_r16Present bool
	if CodebookComboParametersAdditionPerBC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ParallelTxMsgA_SRS_PUCCH_PUSCH_r16Present {
		ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_r16 = new(CA_ParametersNR_v1610_parallelTxMsgA_SRS_PUCCH_PUSCH_r16)
		if err = ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ParallelTxMsgA_SRS_PUCCH_PUSCH_r16", err)
		}
	}
	if MsgA_SUL_r16Present {
		ie.MsgA_SUL_r16 = new(CA_ParametersNR_v1610_msgA_SUL_r16)
		if err = ie.MsgA_SUL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_SUL_r16", err)
		}
	}
	if JointSearchSpaceSwitchAcrossCells_r16Present {
		ie.JointSearchSpaceSwitchAcrossCells_r16 = new(CA_ParametersNR_v1610_jointSearchSpaceSwitchAcrossCells_r16)
		if err = ie.JointSearchSpaceSwitchAcrossCells_r16.Decode(r); err != nil {
			return utils.WrapError("Decode JointSearchSpaceSwitchAcrossCells_r16", err)
		}
	}
	if Half_DuplexTDD_CA_SameSCS_r16Present {
		ie.Half_DuplexTDD_CA_SameSCS_r16 = new(CA_ParametersNR_v1610_half_DuplexTDD_CA_SameSCS_r16)
		if err = ie.Half_DuplexTDD_CA_SameSCS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Half_DuplexTDD_CA_SameSCS_r16", err)
		}
	}
	if ScellDormancyWithinActiveTime_r16Present {
		ie.ScellDormancyWithinActiveTime_r16 = new(CA_ParametersNR_v1610_scellDormancyWithinActiveTime_r16)
		if err = ie.ScellDormancyWithinActiveTime_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ScellDormancyWithinActiveTime_r16", err)
		}
	}
	if ScellDormancyOutsideActiveTime_r16Present {
		ie.ScellDormancyOutsideActiveTime_r16 = new(CA_ParametersNR_v1610_scellDormancyOutsideActiveTime_r16)
		if err = ie.ScellDormancyOutsideActiveTime_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ScellDormancyOutsideActiveTime_r16", err)
		}
	}
	if CrossCarrierA_CSI_trigDiffSCS_r16Present {
		ie.CrossCarrierA_CSI_trigDiffSCS_r16 = new(CA_ParametersNR_v1610_crossCarrierA_CSI_trigDiffSCS_r16)
		if err = ie.CrossCarrierA_CSI_trigDiffSCS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CrossCarrierA_CSI_trigDiffSCS_r16", err)
		}
	}
	if DefaultQCL_CrossCarrierA_CSI_Trig_r16Present {
		ie.DefaultQCL_CrossCarrierA_CSI_Trig_r16 = new(CA_ParametersNR_v1610_defaultQCL_CrossCarrierA_CSI_Trig_r16)
		if err = ie.DefaultQCL_CrossCarrierA_CSI_Trig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DefaultQCL_CrossCarrierA_CSI_Trig_r16", err)
		}
	}
	if InterCA_NonAlignedFrame_r16Present {
		ie.InterCA_NonAlignedFrame_r16 = new(CA_ParametersNR_v1610_interCA_NonAlignedFrame_r16)
		if err = ie.InterCA_NonAlignedFrame_r16.Decode(r); err != nil {
			return utils.WrapError("Decode InterCA_NonAlignedFrame_r16", err)
		}
	}
	if Simul_SRS_Trans_BC_r16Present {
		ie.Simul_SRS_Trans_BC_r16 = new(CA_ParametersNR_v1610_simul_SRS_Trans_BC_r16)
		if err = ie.Simul_SRS_Trans_BC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Simul_SRS_Trans_BC_r16", err)
		}
	}
	if InterFreqDAPS_r16Present {
		ie.InterFreqDAPS_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16)
		if err = ie.InterFreqDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode InterFreqDAPS_r16", err)
		}
	}
	if CodebookParametersPerBC_r16Present {
		ie.CodebookParametersPerBC_r16 = new(CodebookParameters_v1610)
		if err = ie.CodebookParametersPerBC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CodebookParametersPerBC_r16", err)
		}
	}
	if BlindDetectFactor_r16Present {
		var tmp_int_BlindDetectFactor_r16 int64
		if tmp_int_BlindDetectFactor_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode BlindDetectFactor_r16", err)
		}
		ie.BlindDetectFactor_r16 = &tmp_int_BlindDetectFactor_r16
	}
	if Pdcch_MonitoringCA_r16Present {
		ie.Pdcch_MonitoringCA_r16 = new(CA_ParametersNR_v1610_pdcch_MonitoringCA_r16)
		if err = ie.Pdcch_MonitoringCA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_MonitoringCA_r16", err)
		}
	}
	if Pdcch_BlindDetectionCA_Mixed_r16Present {
		ie.Pdcch_BlindDetectionCA_Mixed_r16 = new(CA_ParametersNR_v1610_pdcch_BlindDetectionCA_Mixed_r16)
		if err = ie.Pdcch_BlindDetectionCA_Mixed_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_BlindDetectionCA_Mixed_r16", err)
		}
	}
	if Pdcch_BlindDetectionMCG_UE_r16Present {
		var tmp_int_Pdcch_BlindDetectionMCG_UE_r16 int64
		if tmp_int_Pdcch_BlindDetectionMCG_UE_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
			return utils.WrapError("Decode Pdcch_BlindDetectionMCG_UE_r16", err)
		}
		ie.Pdcch_BlindDetectionMCG_UE_r16 = &tmp_int_Pdcch_BlindDetectionMCG_UE_r16
	}
	if Pdcch_BlindDetectionSCG_UE_r16Present {
		var tmp_int_Pdcch_BlindDetectionSCG_UE_r16 int64
		if tmp_int_Pdcch_BlindDetectionSCG_UE_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
			return utils.WrapError("Decode Pdcch_BlindDetectionSCG_UE_r16", err)
		}
		ie.Pdcch_BlindDetectionSCG_UE_r16 = &tmp_int_Pdcch_BlindDetectionSCG_UE_r16
	}
	if Pdcch_BlindDetectionMCG_UE_Mixed_r16Present {
		ie.Pdcch_BlindDetectionMCG_UE_Mixed_r16 = new(CA_ParametersNR_v1610_pdcch_BlindDetectionMCG_UE_Mixed_r16)
		if err = ie.Pdcch_BlindDetectionMCG_UE_Mixed_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_BlindDetectionMCG_UE_Mixed_r16", err)
		}
	}
	if Pdcch_BlindDetectionSCG_UE_Mixed_r16Present {
		ie.Pdcch_BlindDetectionSCG_UE_Mixed_r16 = new(CA_ParametersNR_v1610_pdcch_BlindDetectionSCG_UE_Mixed_r16)
		if err = ie.Pdcch_BlindDetectionSCG_UE_Mixed_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_BlindDetectionSCG_UE_Mixed_r16", err)
		}
	}
	if CrossCarrierSchedulingDL_DiffSCS_r16Present {
		ie.CrossCarrierSchedulingDL_DiffSCS_r16 = new(CA_ParametersNR_v1610_crossCarrierSchedulingDL_DiffSCS_r16)
		if err = ie.CrossCarrierSchedulingDL_DiffSCS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CrossCarrierSchedulingDL_DiffSCS_r16", err)
		}
	}
	if CrossCarrierSchedulingDefaultQCL_r16Present {
		ie.CrossCarrierSchedulingDefaultQCL_r16 = new(CA_ParametersNR_v1610_crossCarrierSchedulingDefaultQCL_r16)
		if err = ie.CrossCarrierSchedulingDefaultQCL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CrossCarrierSchedulingDefaultQCL_r16", err)
		}
	}
	if CrossCarrierSchedulingUL_DiffSCS_r16Present {
		ie.CrossCarrierSchedulingUL_DiffSCS_r16 = new(CA_ParametersNR_v1610_crossCarrierSchedulingUL_DiffSCS_r16)
		if err = ie.CrossCarrierSchedulingUL_DiffSCS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CrossCarrierSchedulingUL_DiffSCS_r16", err)
		}
	}
	if Simul_SRS_MIMO_Trans_BC_r16Present {
		ie.Simul_SRS_MIMO_Trans_BC_r16 = new(CA_ParametersNR_v1610_simul_SRS_MIMO_Trans_BC_r16)
		if err = ie.Simul_SRS_MIMO_Trans_BC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Simul_SRS_MIMO_Trans_BC_r16", err)
		}
	}
	if CodebookParametersAdditionPerBC_r16Present {
		ie.CodebookParametersAdditionPerBC_r16 = new(CodebookParametersAdditionPerBC_r16)
		if err = ie.CodebookParametersAdditionPerBC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CodebookParametersAdditionPerBC_r16", err)
		}
	}
	if CodebookComboParametersAdditionPerBC_r16Present {
		ie.CodebookComboParametersAdditionPerBC_r16 = new(CodebookComboParametersAdditionPerBC_r16)
		if err = ie.CodebookComboParametersAdditionPerBC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CodebookComboParametersAdditionPerBC_r16", err)
		}
	}
	return nil
}
