package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NTN_Config_r17 struct {
	EpochTime_r17                  *EpochTime_r17                                 `optional`
	Ntn_UlSyncValidityDuration_r17 *NTN_Config_r17_ntn_UlSyncValidityDuration_r17 `optional`
	CellSpecificKoffset_r17        *int64                                         `lb:1,ub:1023,optional`
	Kmac_r17                       *int64                                         `lb:1,ub:512,optional`
	Ta_Info_r17                    *TA_Info_r17                                   `optional`
	Ntn_PolarizationDL_r17         *NTN_Config_r17_ntn_PolarizationDL_r17         `optional`
	Ntn_PolarizationUL_r17         *NTN_Config_r17_ntn_PolarizationUL_r17         `optional`
	EphemerisInfo_r17              *EphemerisInfo_r17                             `optional`
	Ta_Report_r17                  *NTN_Config_r17_ta_Report_r17                  `optional`
}

func (ie *NTN_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.EpochTime_r17 != nil, ie.Ntn_UlSyncValidityDuration_r17 != nil, ie.CellSpecificKoffset_r17 != nil, ie.Kmac_r17 != nil, ie.Ta_Info_r17 != nil, ie.Ntn_PolarizationDL_r17 != nil, ie.Ntn_PolarizationUL_r17 != nil, ie.EphemerisInfo_r17 != nil, ie.Ta_Report_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.EpochTime_r17 != nil {
		if err = ie.EpochTime_r17.Encode(w); err != nil {
			return utils.WrapError("Encode EpochTime_r17", err)
		}
	}
	if ie.Ntn_UlSyncValidityDuration_r17 != nil {
		if err = ie.Ntn_UlSyncValidityDuration_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ntn_UlSyncValidityDuration_r17", err)
		}
	}
	if ie.CellSpecificKoffset_r17 != nil {
		if err = w.WriteInteger(*ie.CellSpecificKoffset_r17, &uper.Constraint{Lb: 1, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode CellSpecificKoffset_r17", err)
		}
	}
	if ie.Kmac_r17 != nil {
		if err = w.WriteInteger(*ie.Kmac_r17, &uper.Constraint{Lb: 1, Ub: 512}, false); err != nil {
			return utils.WrapError("Encode Kmac_r17", err)
		}
	}
	if ie.Ta_Info_r17 != nil {
		if err = ie.Ta_Info_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ta_Info_r17", err)
		}
	}
	if ie.Ntn_PolarizationDL_r17 != nil {
		if err = ie.Ntn_PolarizationDL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ntn_PolarizationDL_r17", err)
		}
	}
	if ie.Ntn_PolarizationUL_r17 != nil {
		if err = ie.Ntn_PolarizationUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ntn_PolarizationUL_r17", err)
		}
	}
	if ie.EphemerisInfo_r17 != nil {
		if err = ie.EphemerisInfo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode EphemerisInfo_r17", err)
		}
	}
	if ie.Ta_Report_r17 != nil {
		if err = ie.Ta_Report_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ta_Report_r17", err)
		}
	}
	return nil
}

func (ie *NTN_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var EpochTime_r17Present bool
	if EpochTime_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ntn_UlSyncValidityDuration_r17Present bool
	if Ntn_UlSyncValidityDuration_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CellSpecificKoffset_r17Present bool
	if CellSpecificKoffset_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Kmac_r17Present bool
	if Kmac_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ta_Info_r17Present bool
	if Ta_Info_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ntn_PolarizationDL_r17Present bool
	if Ntn_PolarizationDL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ntn_PolarizationUL_r17Present bool
	if Ntn_PolarizationUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var EphemerisInfo_r17Present bool
	if EphemerisInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ta_Report_r17Present bool
	if Ta_Report_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if EpochTime_r17Present {
		ie.EpochTime_r17 = new(EpochTime_r17)
		if err = ie.EpochTime_r17.Decode(r); err != nil {
			return utils.WrapError("Decode EpochTime_r17", err)
		}
	}
	if Ntn_UlSyncValidityDuration_r17Present {
		ie.Ntn_UlSyncValidityDuration_r17 = new(NTN_Config_r17_ntn_UlSyncValidityDuration_r17)
		if err = ie.Ntn_UlSyncValidityDuration_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ntn_UlSyncValidityDuration_r17", err)
		}
	}
	if CellSpecificKoffset_r17Present {
		var tmp_int_CellSpecificKoffset_r17 int64
		if tmp_int_CellSpecificKoffset_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode CellSpecificKoffset_r17", err)
		}
		ie.CellSpecificKoffset_r17 = &tmp_int_CellSpecificKoffset_r17
	}
	if Kmac_r17Present {
		var tmp_int_Kmac_r17 int64
		if tmp_int_Kmac_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 512}, false); err != nil {
			return utils.WrapError("Decode Kmac_r17", err)
		}
		ie.Kmac_r17 = &tmp_int_Kmac_r17
	}
	if Ta_Info_r17Present {
		ie.Ta_Info_r17 = new(TA_Info_r17)
		if err = ie.Ta_Info_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ta_Info_r17", err)
		}
	}
	if Ntn_PolarizationDL_r17Present {
		ie.Ntn_PolarizationDL_r17 = new(NTN_Config_r17_ntn_PolarizationDL_r17)
		if err = ie.Ntn_PolarizationDL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ntn_PolarizationDL_r17", err)
		}
	}
	if Ntn_PolarizationUL_r17Present {
		ie.Ntn_PolarizationUL_r17 = new(NTN_Config_r17_ntn_PolarizationUL_r17)
		if err = ie.Ntn_PolarizationUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ntn_PolarizationUL_r17", err)
		}
	}
	if EphemerisInfo_r17Present {
		ie.EphemerisInfo_r17 = new(EphemerisInfo_r17)
		if err = ie.EphemerisInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode EphemerisInfo_r17", err)
		}
	}
	if Ta_Report_r17Present {
		ie.Ta_Report_r17 = new(NTN_Config_r17_ta_Report_r17)
		if err = ie.Ta_Report_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ta_Report_r17", err)
		}
	}
	return nil
}
