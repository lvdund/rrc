package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1630 struct {
	SimulTX_SRS_AntSwitchingInterBandUL_CA_r16 *SimulSRS_ForAntennaSwitching_r16                                `optional`
	BeamManagementType_r16                     *CA_ParametersNR_v1630_beamManagementType_r16                    `optional`
	IntraBandFreqSeparationUL_AggBW_GapBW_r16  *CA_ParametersNR_v1630_intraBandFreqSeparationUL_AggBW_GapBW_r16 `optional`
	InterCA_NonAlignedFrame_B_r16              *CA_ParametersNR_v1630_interCA_NonAlignedFrame_B_r16             `optional`
}

func (ie *CA_ParametersNR_v1630) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SimulTX_SRS_AntSwitchingInterBandUL_CA_r16 != nil, ie.BeamManagementType_r16 != nil, ie.IntraBandFreqSeparationUL_AggBW_GapBW_r16 != nil, ie.InterCA_NonAlignedFrame_B_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SimulTX_SRS_AntSwitchingInterBandUL_CA_r16 != nil {
		if err = ie.SimulTX_SRS_AntSwitchingInterBandUL_CA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SimulTX_SRS_AntSwitchingInterBandUL_CA_r16", err)
		}
	}
	if ie.BeamManagementType_r16 != nil {
		if err = ie.BeamManagementType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode BeamManagementType_r16", err)
		}
	}
	if ie.IntraBandFreqSeparationUL_AggBW_GapBW_r16 != nil {
		if err = ie.IntraBandFreqSeparationUL_AggBW_GapBW_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IntraBandFreqSeparationUL_AggBW_GapBW_r16", err)
		}
	}
	if ie.InterCA_NonAlignedFrame_B_r16 != nil {
		if err = ie.InterCA_NonAlignedFrame_B_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterCA_NonAlignedFrame_B_r16", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1630) Decode(r *aper.AperReader) error {
	var err error
	var SimulTX_SRS_AntSwitchingInterBandUL_CA_r16Present bool
	if SimulTX_SRS_AntSwitchingInterBandUL_CA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BeamManagementType_r16Present bool
	if BeamManagementType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraBandFreqSeparationUL_AggBW_GapBW_r16Present bool
	if IntraBandFreqSeparationUL_AggBW_GapBW_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterCA_NonAlignedFrame_B_r16Present bool
	if InterCA_NonAlignedFrame_B_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SimulTX_SRS_AntSwitchingInterBandUL_CA_r16Present {
		ie.SimulTX_SRS_AntSwitchingInterBandUL_CA_r16 = new(SimulSRS_ForAntennaSwitching_r16)
		if err = ie.SimulTX_SRS_AntSwitchingInterBandUL_CA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SimulTX_SRS_AntSwitchingInterBandUL_CA_r16", err)
		}
	}
	if BeamManagementType_r16Present {
		ie.BeamManagementType_r16 = new(CA_ParametersNR_v1630_beamManagementType_r16)
		if err = ie.BeamManagementType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode BeamManagementType_r16", err)
		}
	}
	if IntraBandFreqSeparationUL_AggBW_GapBW_r16Present {
		ie.IntraBandFreqSeparationUL_AggBW_GapBW_r16 = new(CA_ParametersNR_v1630_intraBandFreqSeparationUL_AggBW_GapBW_r16)
		if err = ie.IntraBandFreqSeparationUL_AggBW_GapBW_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IntraBandFreqSeparationUL_AggBW_GapBW_r16", err)
		}
	}
	if InterCA_NonAlignedFrame_B_r16Present {
		ie.InterCA_NonAlignedFrame_B_r16 = new(CA_ParametersNR_v1630_interCA_NonAlignedFrame_B_r16)
		if err = ie.InterCA_NonAlignedFrame_B_r16.Decode(r); err != nil {
			return utils.WrapError("Decode InterCA_NonAlignedFrame_B_r16", err)
		}
	}
	return nil
}
