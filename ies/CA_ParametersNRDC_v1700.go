package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNRDC_v1700 struct {
	SimultaneousRxTx_IAB_MultipleParents_r17 *CA_ParametersNRDC_v1700_simultaneousRxTx_IAB_MultipleParents_r17 `optional`
	CondPSCellAdditionNRDC_r17               *CA_ParametersNRDC_v1700_condPSCellAdditionNRDC_r17               `optional`
	Scg_ActivationDeactivationNRDC_r17       *CA_ParametersNRDC_v1700_scg_ActivationDeactivationNRDC_r17       `optional`
	Scg_ActivationDeactivationResumeNRDC_r17 *CA_ParametersNRDC_v1700_scg_ActivationDeactivationResumeNRDC_r17 `optional`
	BeamManagementType_CBM_r17               *CA_ParametersNRDC_v1700_beamManagementType_CBM_r17               `optional`
}

func (ie *CA_ParametersNRDC_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SimultaneousRxTx_IAB_MultipleParents_r17 != nil, ie.CondPSCellAdditionNRDC_r17 != nil, ie.Scg_ActivationDeactivationNRDC_r17 != nil, ie.Scg_ActivationDeactivationResumeNRDC_r17 != nil, ie.BeamManagementType_CBM_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SimultaneousRxTx_IAB_MultipleParents_r17 != nil {
		if err = ie.SimultaneousRxTx_IAB_MultipleParents_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SimultaneousRxTx_IAB_MultipleParents_r17", err)
		}
	}
	if ie.CondPSCellAdditionNRDC_r17 != nil {
		if err = ie.CondPSCellAdditionNRDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CondPSCellAdditionNRDC_r17", err)
		}
	}
	if ie.Scg_ActivationDeactivationNRDC_r17 != nil {
		if err = ie.Scg_ActivationDeactivationNRDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scg_ActivationDeactivationNRDC_r17", err)
		}
	}
	if ie.Scg_ActivationDeactivationResumeNRDC_r17 != nil {
		if err = ie.Scg_ActivationDeactivationResumeNRDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scg_ActivationDeactivationResumeNRDC_r17", err)
		}
	}
	if ie.BeamManagementType_CBM_r17 != nil {
		if err = ie.BeamManagementType_CBM_r17.Encode(w); err != nil {
			return utils.WrapError("Encode BeamManagementType_CBM_r17", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNRDC_v1700) Decode(r *uper.UperReader) error {
	var err error
	var SimultaneousRxTx_IAB_MultipleParents_r17Present bool
	if SimultaneousRxTx_IAB_MultipleParents_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CondPSCellAdditionNRDC_r17Present bool
	if CondPSCellAdditionNRDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scg_ActivationDeactivationNRDC_r17Present bool
	if Scg_ActivationDeactivationNRDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scg_ActivationDeactivationResumeNRDC_r17Present bool
	if Scg_ActivationDeactivationResumeNRDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BeamManagementType_CBM_r17Present bool
	if BeamManagementType_CBM_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SimultaneousRxTx_IAB_MultipleParents_r17Present {
		ie.SimultaneousRxTx_IAB_MultipleParents_r17 = new(CA_ParametersNRDC_v1700_simultaneousRxTx_IAB_MultipleParents_r17)
		if err = ie.SimultaneousRxTx_IAB_MultipleParents_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SimultaneousRxTx_IAB_MultipleParents_r17", err)
		}
	}
	if CondPSCellAdditionNRDC_r17Present {
		ie.CondPSCellAdditionNRDC_r17 = new(CA_ParametersNRDC_v1700_condPSCellAdditionNRDC_r17)
		if err = ie.CondPSCellAdditionNRDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CondPSCellAdditionNRDC_r17", err)
		}
	}
	if Scg_ActivationDeactivationNRDC_r17Present {
		ie.Scg_ActivationDeactivationNRDC_r17 = new(CA_ParametersNRDC_v1700_scg_ActivationDeactivationNRDC_r17)
		if err = ie.Scg_ActivationDeactivationNRDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scg_ActivationDeactivationNRDC_r17", err)
		}
	}
	if Scg_ActivationDeactivationResumeNRDC_r17Present {
		ie.Scg_ActivationDeactivationResumeNRDC_r17 = new(CA_ParametersNRDC_v1700_scg_ActivationDeactivationResumeNRDC_r17)
		if err = ie.Scg_ActivationDeactivationResumeNRDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scg_ActivationDeactivationResumeNRDC_r17", err)
		}
	}
	if BeamManagementType_CBM_r17Present {
		ie.BeamManagementType_CBM_r17 = new(CA_ParametersNRDC_v1700_beamManagementType_CBM_r17)
		if err = ie.BeamManagementType_CBM_r17.Decode(r); err != nil {
			return utils.WrapError("Decode BeamManagementType_CBM_r17", err)
		}
	}
	return nil
}
