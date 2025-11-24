package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SuccessHO_Report_r17 struct {
	SourceCellInfo_r17         *SuccessHO_Report_r17_sourceCellInfo_r17       `optional`
	TargetCellInfo_r17         *SuccessHO_Report_r17_targetCellInfo_r17       `optional`
	MeasResultNeighCells_r17   *SuccessHO_Report_r17_measResultNeighCells_r17 `optional`
	LocationInfo_r17           *LocationInfo_r16                              `optional`
	TimeSinceCHO_Reconfig_r17  *TimeSinceCHO_Reconfig_r17                     `optional`
	Shr_Cause_r17              *SHR_Cause_r17                                 `optional`
	Ra_InformationCommon_r17   *RA_InformationCommon_r16                      `optional`
	UpInterruptionTimeAtHO_r17 *UPInterruptionTimeAtHO_r17                    `optional`
	C_RNTI_r17                 *RNTI_Value                                    `optional`
}

func (ie *SuccessHO_Report_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SourceCellInfo_r17 != nil, ie.TargetCellInfo_r17 != nil, ie.MeasResultNeighCells_r17 != nil, ie.LocationInfo_r17 != nil, ie.TimeSinceCHO_Reconfig_r17 != nil, ie.Shr_Cause_r17 != nil, ie.Ra_InformationCommon_r17 != nil, ie.UpInterruptionTimeAtHO_r17 != nil, ie.C_RNTI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SourceCellInfo_r17 != nil {
		if err = ie.SourceCellInfo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SourceCellInfo_r17", err)
		}
	}
	if ie.TargetCellInfo_r17 != nil {
		if err = ie.TargetCellInfo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode TargetCellInfo_r17", err)
		}
	}
	if ie.MeasResultNeighCells_r17 != nil {
		if err = ie.MeasResultNeighCells_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultNeighCells_r17", err)
		}
	}
	if ie.LocationInfo_r17 != nil {
		if err = ie.LocationInfo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode LocationInfo_r17", err)
		}
	}
	if ie.TimeSinceCHO_Reconfig_r17 != nil {
		if err = ie.TimeSinceCHO_Reconfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode TimeSinceCHO_Reconfig_r17", err)
		}
	}
	if ie.Shr_Cause_r17 != nil {
		if err = ie.Shr_Cause_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Shr_Cause_r17", err)
		}
	}
	if ie.Ra_InformationCommon_r17 != nil {
		if err = ie.Ra_InformationCommon_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ra_InformationCommon_r17", err)
		}
	}
	if ie.UpInterruptionTimeAtHO_r17 != nil {
		if err = ie.UpInterruptionTimeAtHO_r17.Encode(w); err != nil {
			return utils.WrapError("Encode UpInterruptionTimeAtHO_r17", err)
		}
	}
	if ie.C_RNTI_r17 != nil {
		if err = ie.C_RNTI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode C_RNTI_r17", err)
		}
	}
	return nil
}

func (ie *SuccessHO_Report_r17) Decode(r *uper.UperReader) error {
	var err error
	var SourceCellInfo_r17Present bool
	if SourceCellInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TargetCellInfo_r17Present bool
	if TargetCellInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultNeighCells_r17Present bool
	if MeasResultNeighCells_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LocationInfo_r17Present bool
	if LocationInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TimeSinceCHO_Reconfig_r17Present bool
	if TimeSinceCHO_Reconfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Shr_Cause_r17Present bool
	if Shr_Cause_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_InformationCommon_r17Present bool
	if Ra_InformationCommon_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var UpInterruptionTimeAtHO_r17Present bool
	if UpInterruptionTimeAtHO_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var C_RNTI_r17Present bool
	if C_RNTI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SourceCellInfo_r17Present {
		ie.SourceCellInfo_r17 = new(SuccessHO_Report_r17_sourceCellInfo_r17)
		if err = ie.SourceCellInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SourceCellInfo_r17", err)
		}
	}
	if TargetCellInfo_r17Present {
		ie.TargetCellInfo_r17 = new(SuccessHO_Report_r17_targetCellInfo_r17)
		if err = ie.TargetCellInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode TargetCellInfo_r17", err)
		}
	}
	if MeasResultNeighCells_r17Present {
		ie.MeasResultNeighCells_r17 = new(SuccessHO_Report_r17_measResultNeighCells_r17)
		if err = ie.MeasResultNeighCells_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultNeighCells_r17", err)
		}
	}
	if LocationInfo_r17Present {
		ie.LocationInfo_r17 = new(LocationInfo_r16)
		if err = ie.LocationInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode LocationInfo_r17", err)
		}
	}
	if TimeSinceCHO_Reconfig_r17Present {
		ie.TimeSinceCHO_Reconfig_r17 = new(TimeSinceCHO_Reconfig_r17)
		if err = ie.TimeSinceCHO_Reconfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode TimeSinceCHO_Reconfig_r17", err)
		}
	}
	if Shr_Cause_r17Present {
		ie.Shr_Cause_r17 = new(SHR_Cause_r17)
		if err = ie.Shr_Cause_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Shr_Cause_r17", err)
		}
	}
	if Ra_InformationCommon_r17Present {
		ie.Ra_InformationCommon_r17 = new(RA_InformationCommon_r16)
		if err = ie.Ra_InformationCommon_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ra_InformationCommon_r17", err)
		}
	}
	if UpInterruptionTimeAtHO_r17Present {
		ie.UpInterruptionTimeAtHO_r17 = new(UPInterruptionTimeAtHO_r17)
		if err = ie.UpInterruptionTimeAtHO_r17.Decode(r); err != nil {
			return utils.WrapError("Decode UpInterruptionTimeAtHO_r17", err)
		}
	}
	if C_RNTI_r17Present {
		ie.C_RNTI_r17 = new(RNTI_Value)
		if err = ie.C_RNTI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode C_RNTI_r17", err)
		}
	}
	return nil
}
