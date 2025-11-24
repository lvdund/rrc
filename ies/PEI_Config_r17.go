package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PEI_Config_r17 struct {
	Po_NumPerPEI_r17       PEI_Config_r17_po_NumPerPEI_r17      `madatory`
	PayloadSizeDCI_2_7_r17 int64                                `lb:1,ub:maxDCI_2_7_Size_r17,madatory`
	Pei_FrameOffset_r17    int64                                `lb:0,ub:16,madatory`
	SubgroupConfig_r17     SubgroupConfig_r17                   `madatory`
	LastUsedCellOnly_r17   *PEI_Config_r17_lastUsedCellOnly_r17 `optional`
}

func (ie *PEI_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.LastUsedCellOnly_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Po_NumPerPEI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Po_NumPerPEI_r17", err)
	}
	if err = w.WriteInteger(ie.PayloadSizeDCI_2_7_r17, &uper.Constraint{Lb: 1, Ub: maxDCI_2_7_Size_r17}, false); err != nil {
		return utils.WrapError("WriteInteger PayloadSizeDCI_2_7_r17", err)
	}
	if err = w.WriteInteger(ie.Pei_FrameOffset_r17, &uper.Constraint{Lb: 0, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger Pei_FrameOffset_r17", err)
	}
	if err = ie.SubgroupConfig_r17.Encode(w); err != nil {
		return utils.WrapError("Encode SubgroupConfig_r17", err)
	}
	if ie.LastUsedCellOnly_r17 != nil {
		if err = ie.LastUsedCellOnly_r17.Encode(w); err != nil {
			return utils.WrapError("Encode LastUsedCellOnly_r17", err)
		}
	}
	return nil
}

func (ie *PEI_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var LastUsedCellOnly_r17Present bool
	if LastUsedCellOnly_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Po_NumPerPEI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Po_NumPerPEI_r17", err)
	}
	var tmp_int_PayloadSizeDCI_2_7_r17 int64
	if tmp_int_PayloadSizeDCI_2_7_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxDCI_2_7_Size_r17}, false); err != nil {
		return utils.WrapError("ReadInteger PayloadSizeDCI_2_7_r17", err)
	}
	ie.PayloadSizeDCI_2_7_r17 = tmp_int_PayloadSizeDCI_2_7_r17
	var tmp_int_Pei_FrameOffset_r17 int64
	if tmp_int_Pei_FrameOffset_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger Pei_FrameOffset_r17", err)
	}
	ie.Pei_FrameOffset_r17 = tmp_int_Pei_FrameOffset_r17
	if err = ie.SubgroupConfig_r17.Decode(r); err != nil {
		return utils.WrapError("Decode SubgroupConfig_r17", err)
	}
	if LastUsedCellOnly_r17Present {
		ie.LastUsedCellOnly_r17 = new(PEI_Config_r17_lastUsedCellOnly_r17)
		if err = ie.LastUsedCellOnly_r17.Decode(r); err != nil {
			return utils.WrapError("Decode LastUsedCellOnly_r17", err)
		}
	}
	return nil
}
