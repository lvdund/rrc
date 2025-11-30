package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BeamFailureDetectionSet_r17 struct {
	BfdResourcesToAddModList_r17    []BeamLinkMonitoringRS_r17                                   `lb:1,ub:maxNrofBFDResourcePerSet_r17,optional`
	BfdResourcesToReleaseList_r17   []BeamLinkMonitoringRS_Id_r17                                `lb:1,ub:maxNrofBFDResourcePerSet_r17,optional`
	BeamFailureInstanceMaxCount_r17 *BeamFailureDetectionSet_r17_beamFailureInstanceMaxCount_r17 `optional`
	BeamFailureDetectionTimer_r17   *BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17   `optional`
}

func (ie *BeamFailureDetectionSet_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.BfdResourcesToAddModList_r17) > 0, len(ie.BfdResourcesToReleaseList_r17) > 0, ie.BeamFailureInstanceMaxCount_r17 != nil, ie.BeamFailureDetectionTimer_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.BfdResourcesToAddModList_r17) > 0 {
		tmp_BfdResourcesToAddModList_r17 := utils.NewSequence[*BeamLinkMonitoringRS_r17]([]*BeamLinkMonitoringRS_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofBFDResourcePerSet_r17}, false)
		for _, i := range ie.BfdResourcesToAddModList_r17 {
			tmp_BfdResourcesToAddModList_r17.Value = append(tmp_BfdResourcesToAddModList_r17.Value, &i)
		}
		if err = tmp_BfdResourcesToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode BfdResourcesToAddModList_r17", err)
		}
	}
	if len(ie.BfdResourcesToReleaseList_r17) > 0 {
		tmp_BfdResourcesToReleaseList_r17 := utils.NewSequence[*BeamLinkMonitoringRS_Id_r17]([]*BeamLinkMonitoringRS_Id_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofBFDResourcePerSet_r17}, false)
		for _, i := range ie.BfdResourcesToReleaseList_r17 {
			tmp_BfdResourcesToReleaseList_r17.Value = append(tmp_BfdResourcesToReleaseList_r17.Value, &i)
		}
		if err = tmp_BfdResourcesToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode BfdResourcesToReleaseList_r17", err)
		}
	}
	if ie.BeamFailureInstanceMaxCount_r17 != nil {
		if err = ie.BeamFailureInstanceMaxCount_r17.Encode(w); err != nil {
			return utils.WrapError("Encode BeamFailureInstanceMaxCount_r17", err)
		}
	}
	if ie.BeamFailureDetectionTimer_r17 != nil {
		if err = ie.BeamFailureDetectionTimer_r17.Encode(w); err != nil {
			return utils.WrapError("Encode BeamFailureDetectionTimer_r17", err)
		}
	}
	return nil
}

func (ie *BeamFailureDetectionSet_r17) Decode(r *aper.AperReader) error {
	var err error
	var BfdResourcesToAddModList_r17Present bool
	if BfdResourcesToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BfdResourcesToReleaseList_r17Present bool
	if BfdResourcesToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BeamFailureInstanceMaxCount_r17Present bool
	if BeamFailureInstanceMaxCount_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BeamFailureDetectionTimer_r17Present bool
	if BeamFailureDetectionTimer_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if BfdResourcesToAddModList_r17Present {
		tmp_BfdResourcesToAddModList_r17 := utils.NewSequence[*BeamLinkMonitoringRS_r17]([]*BeamLinkMonitoringRS_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofBFDResourcePerSet_r17}, false)
		fn_BfdResourcesToAddModList_r17 := func() *BeamLinkMonitoringRS_r17 {
			return new(BeamLinkMonitoringRS_r17)
		}
		if err = tmp_BfdResourcesToAddModList_r17.Decode(r, fn_BfdResourcesToAddModList_r17); err != nil {
			return utils.WrapError("Decode BfdResourcesToAddModList_r17", err)
		}
		ie.BfdResourcesToAddModList_r17 = []BeamLinkMonitoringRS_r17{}
		for _, i := range tmp_BfdResourcesToAddModList_r17.Value {
			ie.BfdResourcesToAddModList_r17 = append(ie.BfdResourcesToAddModList_r17, *i)
		}
	}
	if BfdResourcesToReleaseList_r17Present {
		tmp_BfdResourcesToReleaseList_r17 := utils.NewSequence[*BeamLinkMonitoringRS_Id_r17]([]*BeamLinkMonitoringRS_Id_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofBFDResourcePerSet_r17}, false)
		fn_BfdResourcesToReleaseList_r17 := func() *BeamLinkMonitoringRS_Id_r17 {
			return new(BeamLinkMonitoringRS_Id_r17)
		}
		if err = tmp_BfdResourcesToReleaseList_r17.Decode(r, fn_BfdResourcesToReleaseList_r17); err != nil {
			return utils.WrapError("Decode BfdResourcesToReleaseList_r17", err)
		}
		ie.BfdResourcesToReleaseList_r17 = []BeamLinkMonitoringRS_Id_r17{}
		for _, i := range tmp_BfdResourcesToReleaseList_r17.Value {
			ie.BfdResourcesToReleaseList_r17 = append(ie.BfdResourcesToReleaseList_r17, *i)
		}
	}
	if BeamFailureInstanceMaxCount_r17Present {
		ie.BeamFailureInstanceMaxCount_r17 = new(BeamFailureDetectionSet_r17_beamFailureInstanceMaxCount_r17)
		if err = ie.BeamFailureInstanceMaxCount_r17.Decode(r); err != nil {
			return utils.WrapError("Decode BeamFailureInstanceMaxCount_r17", err)
		}
	}
	if BeamFailureDetectionTimer_r17Present {
		ie.BeamFailureDetectionTimer_r17 = new(BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17)
		if err = ie.BeamFailureDetectionTimer_r17.Decode(r); err != nil {
			return utils.WrapError("Decode BeamFailureDetectionTimer_r17", err)
		}
	}
	return nil
}
