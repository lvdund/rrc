package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17 struct {
	Inter_SN_condPSCellChangeFDD_TDD_NRDC_r17    *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFDD_TDD_NRDC_r17    `optional`
	Inter_SN_condPSCellChangeFR1_FR2_NRDC_r17    *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFR1_FR2_NRDC_r17    `optional`
	Inter_SN_condPSCellChangeFDD_TDD_ENDC_r17    *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFDD_TDD_ENDC_r17    `optional`
	Inter_SN_condPSCellChangeFR1_FR2_ENDC_r17    *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFR1_FR2_ENDC_r17    `optional`
	Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 `optional`
	Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 `optional`
	Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 `optional`
	Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 `optional`
	Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 `optional`
	Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 `optional`
}

func (ie *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Inter_SN_condPSCellChangeFDD_TDD_NRDC_r17 != nil, ie.Inter_SN_condPSCellChangeFR1_FR2_NRDC_r17 != nil, ie.Inter_SN_condPSCellChangeFDD_TDD_ENDC_r17 != nil, ie.Inter_SN_condPSCellChangeFR1_FR2_ENDC_r17 != nil, ie.Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 != nil, ie.Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 != nil, ie.Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 != nil, ie.Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 != nil, ie.Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 != nil, ie.Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Inter_SN_condPSCellChangeFDD_TDD_NRDC_r17 != nil {
		if err = ie.Inter_SN_condPSCellChangeFDD_TDD_NRDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Inter_SN_condPSCellChangeFDD_TDD_NRDC_r17", err)
		}
	}
	if ie.Inter_SN_condPSCellChangeFR1_FR2_NRDC_r17 != nil {
		if err = ie.Inter_SN_condPSCellChangeFR1_FR2_NRDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Inter_SN_condPSCellChangeFR1_FR2_NRDC_r17", err)
		}
	}
	if ie.Inter_SN_condPSCellChangeFDD_TDD_ENDC_r17 != nil {
		if err = ie.Inter_SN_condPSCellChangeFDD_TDD_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Inter_SN_condPSCellChangeFDD_TDD_ENDC_r17", err)
		}
	}
	if ie.Inter_SN_condPSCellChangeFR1_FR2_ENDC_r17 != nil {
		if err = ie.Inter_SN_condPSCellChangeFR1_FR2_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Inter_SN_condPSCellChangeFR1_FR2_ENDC_r17", err)
		}
	}
	if ie.Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 != nil {
		if err = ie.Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17", err)
		}
	}
	if ie.Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 != nil {
		if err = ie.Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17", err)
		}
	}
	if ie.Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 != nil {
		if err = ie.Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17", err)
		}
	}
	if ie.Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 != nil {
		if err = ie.Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17", err)
		}
	}
	if ie.Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 != nil {
		if err = ie.Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17", err)
		}
	}
	if ie.Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 != nil {
		if err = ie.Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17) Decode(r *uper.UperReader) error {
	var err error
	var Inter_SN_condPSCellChangeFDD_TDD_NRDC_r17Present bool
	if Inter_SN_condPSCellChangeFDD_TDD_NRDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Inter_SN_condPSCellChangeFR1_FR2_NRDC_r17Present bool
	if Inter_SN_condPSCellChangeFR1_FR2_NRDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Inter_SN_condPSCellChangeFDD_TDD_ENDC_r17Present bool
	if Inter_SN_condPSCellChangeFDD_TDD_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Inter_SN_condPSCellChangeFR1_FR2_ENDC_r17Present bool
	if Inter_SN_condPSCellChangeFR1_FR2_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17Present bool
	if Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17Present bool
	if Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17Present bool
	if Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17Present bool
	if Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17Present bool
	if Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17Present bool
	if Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Inter_SN_condPSCellChangeFDD_TDD_NRDC_r17Present {
		ie.Inter_SN_condPSCellChangeFDD_TDD_NRDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFDD_TDD_NRDC_r17)
		if err = ie.Inter_SN_condPSCellChangeFDD_TDD_NRDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Inter_SN_condPSCellChangeFDD_TDD_NRDC_r17", err)
		}
	}
	if Inter_SN_condPSCellChangeFR1_FR2_NRDC_r17Present {
		ie.Inter_SN_condPSCellChangeFR1_FR2_NRDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFR1_FR2_NRDC_r17)
		if err = ie.Inter_SN_condPSCellChangeFR1_FR2_NRDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Inter_SN_condPSCellChangeFR1_FR2_NRDC_r17", err)
		}
	}
	if Inter_SN_condPSCellChangeFDD_TDD_ENDC_r17Present {
		ie.Inter_SN_condPSCellChangeFDD_TDD_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFDD_TDD_ENDC_r17)
		if err = ie.Inter_SN_condPSCellChangeFDD_TDD_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Inter_SN_condPSCellChangeFDD_TDD_ENDC_r17", err)
		}
	}
	if Inter_SN_condPSCellChangeFR1_FR2_ENDC_r17Present {
		ie.Inter_SN_condPSCellChangeFR1_FR2_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFR1_FR2_ENDC_r17)
		if err = ie.Inter_SN_condPSCellChangeFR1_FR2_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Inter_SN_condPSCellChangeFR1_FR2_ENDC_r17", err)
		}
	}
	if Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17Present {
		ie.Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17)
		if err = ie.Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17", err)
		}
	}
	if Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17Present {
		ie.Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17)
		if err = ie.Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17", err)
		}
	}
	if Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17Present {
		ie.Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17)
		if err = ie.Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17", err)
		}
	}
	if Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17Present {
		ie.Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17)
		if err = ie.Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17", err)
		}
	}
	if Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17Present {
		ie.Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17)
		if err = ie.Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17", err)
		}
	}
	if Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17Present {
		ie.Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17)
		if err = ie.Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17", err)
		}
	}
	return nil
}
