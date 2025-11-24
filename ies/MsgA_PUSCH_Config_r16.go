package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MsgA_PUSCH_Config_r16 struct {
	MsgA_PUSCH_ResourceGroupA_r16 *MsgA_PUSCH_Resource_r16                          `optional`
	MsgA_PUSCH_ResourceGroupB_r16 *MsgA_PUSCH_Resource_r16                          `optional`
	MsgA_TransformPrecoder_r16    *MsgA_PUSCH_Config_r16_msgA_TransformPrecoder_r16 `optional`
	MsgA_DataScramblingIndex_r16  *int64                                            `lb:0,ub:1023,optional`
	MsgA_DeltaPreamble_r16        *int64                                            `lb:-1,ub:6,optional`
}

func (ie *MsgA_PUSCH_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MsgA_PUSCH_ResourceGroupA_r16 != nil, ie.MsgA_PUSCH_ResourceGroupB_r16 != nil, ie.MsgA_TransformPrecoder_r16 != nil, ie.MsgA_DataScramblingIndex_r16 != nil, ie.MsgA_DeltaPreamble_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MsgA_PUSCH_ResourceGroupA_r16 != nil {
		if err = ie.MsgA_PUSCH_ResourceGroupA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_PUSCH_ResourceGroupA_r16", err)
		}
	}
	if ie.MsgA_PUSCH_ResourceGroupB_r16 != nil {
		if err = ie.MsgA_PUSCH_ResourceGroupB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_PUSCH_ResourceGroupB_r16", err)
		}
	}
	if ie.MsgA_TransformPrecoder_r16 != nil {
		if err = ie.MsgA_TransformPrecoder_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_TransformPrecoder_r16", err)
		}
	}
	if ie.MsgA_DataScramblingIndex_r16 != nil {
		if err = w.WriteInteger(*ie.MsgA_DataScramblingIndex_r16, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode MsgA_DataScramblingIndex_r16", err)
		}
	}
	if ie.MsgA_DeltaPreamble_r16 != nil {
		if err = w.WriteInteger(*ie.MsgA_DeltaPreamble_r16, &uper.Constraint{Lb: -1, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode MsgA_DeltaPreamble_r16", err)
		}
	}
	return nil
}

func (ie *MsgA_PUSCH_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var MsgA_PUSCH_ResourceGroupA_r16Present bool
	if MsgA_PUSCH_ResourceGroupA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_PUSCH_ResourceGroupB_r16Present bool
	if MsgA_PUSCH_ResourceGroupB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_TransformPrecoder_r16Present bool
	if MsgA_TransformPrecoder_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_DataScramblingIndex_r16Present bool
	if MsgA_DataScramblingIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_DeltaPreamble_r16Present bool
	if MsgA_DeltaPreamble_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MsgA_PUSCH_ResourceGroupA_r16Present {
		ie.MsgA_PUSCH_ResourceGroupA_r16 = new(MsgA_PUSCH_Resource_r16)
		if err = ie.MsgA_PUSCH_ResourceGroupA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_PUSCH_ResourceGroupA_r16", err)
		}
	}
	if MsgA_PUSCH_ResourceGroupB_r16Present {
		ie.MsgA_PUSCH_ResourceGroupB_r16 = new(MsgA_PUSCH_Resource_r16)
		if err = ie.MsgA_PUSCH_ResourceGroupB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_PUSCH_ResourceGroupB_r16", err)
		}
	}
	if MsgA_TransformPrecoder_r16Present {
		ie.MsgA_TransformPrecoder_r16 = new(MsgA_PUSCH_Config_r16_msgA_TransformPrecoder_r16)
		if err = ie.MsgA_TransformPrecoder_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_TransformPrecoder_r16", err)
		}
	}
	if MsgA_DataScramblingIndex_r16Present {
		var tmp_int_MsgA_DataScramblingIndex_r16 int64
		if tmp_int_MsgA_DataScramblingIndex_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode MsgA_DataScramblingIndex_r16", err)
		}
		ie.MsgA_DataScramblingIndex_r16 = &tmp_int_MsgA_DataScramblingIndex_r16
	}
	if MsgA_DeltaPreamble_r16Present {
		var tmp_int_MsgA_DeltaPreamble_r16 int64
		if tmp_int_MsgA_DeltaPreamble_r16, err = r.ReadInteger(&uper.Constraint{Lb: -1, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode MsgA_DeltaPreamble_r16", err)
		}
		ie.MsgA_DeltaPreamble_r16 = &tmp_int_MsgA_DeltaPreamble_r16
	}
	return nil
}
