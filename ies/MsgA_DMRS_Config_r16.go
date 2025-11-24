package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MsgA_DMRS_Config_r16 struct {
	MsgA_DMRS_AdditionalPosition_r16 *MsgA_DMRS_Config_r16_msgA_DMRS_AdditionalPosition_r16 `optional`
	MsgA_MaxLength_r16               *MsgA_DMRS_Config_r16_msgA_MaxLength_r16               `optional`
	MsgA_PUSCH_DMRS_CDM_Group_r16    *int64                                                 `lb:0,ub:1,optional`
	MsgA_PUSCH_NrofPorts_r16         *int64                                                 `lb:0,ub:1,optional`
	MsgA_ScramblingID0_r16           *int64                                                 `lb:0,ub:65535,optional`
	MsgA_ScramblingID1_r16           *int64                                                 `lb:0,ub:65535,optional`
}

func (ie *MsgA_DMRS_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MsgA_DMRS_AdditionalPosition_r16 != nil, ie.MsgA_MaxLength_r16 != nil, ie.MsgA_PUSCH_DMRS_CDM_Group_r16 != nil, ie.MsgA_PUSCH_NrofPorts_r16 != nil, ie.MsgA_ScramblingID0_r16 != nil, ie.MsgA_ScramblingID1_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MsgA_DMRS_AdditionalPosition_r16 != nil {
		if err = ie.MsgA_DMRS_AdditionalPosition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_DMRS_AdditionalPosition_r16", err)
		}
	}
	if ie.MsgA_MaxLength_r16 != nil {
		if err = ie.MsgA_MaxLength_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_MaxLength_r16", err)
		}
	}
	if ie.MsgA_PUSCH_DMRS_CDM_Group_r16 != nil {
		if err = w.WriteInteger(*ie.MsgA_PUSCH_DMRS_CDM_Group_r16, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode MsgA_PUSCH_DMRS_CDM_Group_r16", err)
		}
	}
	if ie.MsgA_PUSCH_NrofPorts_r16 != nil {
		if err = w.WriteInteger(*ie.MsgA_PUSCH_NrofPorts_r16, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode MsgA_PUSCH_NrofPorts_r16", err)
		}
	}
	if ie.MsgA_ScramblingID0_r16 != nil {
		if err = w.WriteInteger(*ie.MsgA_ScramblingID0_r16, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode MsgA_ScramblingID0_r16", err)
		}
	}
	if ie.MsgA_ScramblingID1_r16 != nil {
		if err = w.WriteInteger(*ie.MsgA_ScramblingID1_r16, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode MsgA_ScramblingID1_r16", err)
		}
	}
	return nil
}

func (ie *MsgA_DMRS_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var MsgA_DMRS_AdditionalPosition_r16Present bool
	if MsgA_DMRS_AdditionalPosition_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_MaxLength_r16Present bool
	if MsgA_MaxLength_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_PUSCH_DMRS_CDM_Group_r16Present bool
	if MsgA_PUSCH_DMRS_CDM_Group_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_PUSCH_NrofPorts_r16Present bool
	if MsgA_PUSCH_NrofPorts_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_ScramblingID0_r16Present bool
	if MsgA_ScramblingID0_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_ScramblingID1_r16Present bool
	if MsgA_ScramblingID1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MsgA_DMRS_AdditionalPosition_r16Present {
		ie.MsgA_DMRS_AdditionalPosition_r16 = new(MsgA_DMRS_Config_r16_msgA_DMRS_AdditionalPosition_r16)
		if err = ie.MsgA_DMRS_AdditionalPosition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_DMRS_AdditionalPosition_r16", err)
		}
	}
	if MsgA_MaxLength_r16Present {
		ie.MsgA_MaxLength_r16 = new(MsgA_DMRS_Config_r16_msgA_MaxLength_r16)
		if err = ie.MsgA_MaxLength_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_MaxLength_r16", err)
		}
	}
	if MsgA_PUSCH_DMRS_CDM_Group_r16Present {
		var tmp_int_MsgA_PUSCH_DMRS_CDM_Group_r16 int64
		if tmp_int_MsgA_PUSCH_DMRS_CDM_Group_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode MsgA_PUSCH_DMRS_CDM_Group_r16", err)
		}
		ie.MsgA_PUSCH_DMRS_CDM_Group_r16 = &tmp_int_MsgA_PUSCH_DMRS_CDM_Group_r16
	}
	if MsgA_PUSCH_NrofPorts_r16Present {
		var tmp_int_MsgA_PUSCH_NrofPorts_r16 int64
		if tmp_int_MsgA_PUSCH_NrofPorts_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode MsgA_PUSCH_NrofPorts_r16", err)
		}
		ie.MsgA_PUSCH_NrofPorts_r16 = &tmp_int_MsgA_PUSCH_NrofPorts_r16
	}
	if MsgA_ScramblingID0_r16Present {
		var tmp_int_MsgA_ScramblingID0_r16 int64
		if tmp_int_MsgA_ScramblingID0_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode MsgA_ScramblingID0_r16", err)
		}
		ie.MsgA_ScramblingID0_r16 = &tmp_int_MsgA_ScramblingID0_r16
	}
	if MsgA_ScramblingID1_r16Present {
		var tmp_int_MsgA_ScramblingID1_r16 int64
		if tmp_int_MsgA_ScramblingID1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode MsgA_ScramblingID1_r16", err)
		}
		ie.MsgA_ScramblingID1_r16 = &tmp_int_MsgA_ScramblingID1_r16
	}
	return nil
}
