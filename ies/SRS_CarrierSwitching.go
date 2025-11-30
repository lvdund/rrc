package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SRS_CarrierSwitching struct {
	Srs_SwitchFromServCellIndex *int64                                     `lb:0,ub:31,optional`
	Srs_SwitchFromCarrier       SRS_CarrierSwitching_srs_SwitchFromCarrier `madatory`
	Srs_TPC_PDCCH_Group         *SRS_CarrierSwitching_srs_TPC_PDCCH_Group  `lb:1,ub:32,optional`
	MonitoringCells             []ServCellIndex                            `lb:1,ub:maxNrofServingCells,optional`
}

func (ie *SRS_CarrierSwitching) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Srs_SwitchFromServCellIndex != nil, ie.Srs_TPC_PDCCH_Group != nil, len(ie.MonitoringCells) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Srs_SwitchFromServCellIndex != nil {
		if err = w.WriteInteger(*ie.Srs_SwitchFromServCellIndex, &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode Srs_SwitchFromServCellIndex", err)
		}
	}
	if err = ie.Srs_SwitchFromCarrier.Encode(w); err != nil {
		return utils.WrapError("Encode Srs_SwitchFromCarrier", err)
	}
	if ie.Srs_TPC_PDCCH_Group != nil {
		if err = ie.Srs_TPC_PDCCH_Group.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_TPC_PDCCH_Group", err)
		}
	}
	if len(ie.MonitoringCells) > 0 {
		tmp_MonitoringCells := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
		for _, i := range ie.MonitoringCells {
			tmp_MonitoringCells.Value = append(tmp_MonitoringCells.Value, &i)
		}
		if err = tmp_MonitoringCells.Encode(w); err != nil {
			return utils.WrapError("Encode MonitoringCells", err)
		}
	}
	return nil
}

func (ie *SRS_CarrierSwitching) Decode(r *aper.AperReader) error {
	var err error
	var Srs_SwitchFromServCellIndexPresent bool
	if Srs_SwitchFromServCellIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_TPC_PDCCH_GroupPresent bool
	if Srs_TPC_PDCCH_GroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MonitoringCellsPresent bool
	if MonitoringCellsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Srs_SwitchFromServCellIndexPresent {
		var tmp_int_Srs_SwitchFromServCellIndex int64
		if tmp_int_Srs_SwitchFromServCellIndex, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode Srs_SwitchFromServCellIndex", err)
		}
		ie.Srs_SwitchFromServCellIndex = &tmp_int_Srs_SwitchFromServCellIndex
	}
	if err = ie.Srs_SwitchFromCarrier.Decode(r); err != nil {
		return utils.WrapError("Decode Srs_SwitchFromCarrier", err)
	}
	if Srs_TPC_PDCCH_GroupPresent {
		ie.Srs_TPC_PDCCH_Group = new(SRS_CarrierSwitching_srs_TPC_PDCCH_Group)
		if err = ie.Srs_TPC_PDCCH_Group.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_TPC_PDCCH_Group", err)
		}
	}
	if MonitoringCellsPresent {
		tmp_MonitoringCells := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
		fn_MonitoringCells := func() *ServCellIndex {
			return new(ServCellIndex)
		}
		if err = tmp_MonitoringCells.Decode(r, fn_MonitoringCells); err != nil {
			return utils.WrapError("Decode MonitoringCells", err)
		}
		ie.MonitoringCells = []ServCellIndex{}
		for _, i := range tmp_MonitoringCells.Value {
			ie.MonitoringCells = append(ie.MonitoringCells, *i)
		}
	}
	return nil
}
