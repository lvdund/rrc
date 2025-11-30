package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r17_codebookType_type1 struct {
	TypeI_SinglePanel_Group1_r17             *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17 `lb:6,ub:6,optional`
	TypeI_SinglePanel_Group2_r17             *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17 `lb:6,ub:6,optional`
	TypeI_SinglePanel_ri_RestrictionSTRP_r17 *aper.BitString                                                     `lb:8,ub:8,optional`
	TypeI_SinglePanel_ri_RestrictionSDM_r17  *aper.BitString                                                     `lb:4,ub:4,optional`
}

func (ie *CodebookConfig_r17_codebookType_type1) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.TypeI_SinglePanel_Group1_r17 != nil, ie.TypeI_SinglePanel_Group2_r17 != nil, ie.TypeI_SinglePanel_ri_RestrictionSTRP_r17 != nil, ie.TypeI_SinglePanel_ri_RestrictionSDM_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.TypeI_SinglePanel_Group1_r17 != nil {
		if err = ie.TypeI_SinglePanel_Group1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode TypeI_SinglePanel_Group1_r17", err)
		}
	}
	if ie.TypeI_SinglePanel_Group2_r17 != nil {
		if err = ie.TypeI_SinglePanel_Group2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode TypeI_SinglePanel_Group2_r17", err)
		}
	}
	if ie.TypeI_SinglePanel_ri_RestrictionSTRP_r17 != nil {
		if err = w.WriteBitString(ie.TypeI_SinglePanel_ri_RestrictionSTRP_r17.Bytes, uint(ie.TypeI_SinglePanel_ri_RestrictionSTRP_r17.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode TypeI_SinglePanel_ri_RestrictionSTRP_r17", err)
		}
	}
	if ie.TypeI_SinglePanel_ri_RestrictionSDM_r17 != nil {
		if err = w.WriteBitString(ie.TypeI_SinglePanel_ri_RestrictionSDM_r17.Bytes, uint(ie.TypeI_SinglePanel_ri_RestrictionSDM_r17.NumBits), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Encode TypeI_SinglePanel_ri_RestrictionSDM_r17", err)
		}
	}
	return nil
}

func (ie *CodebookConfig_r17_codebookType_type1) Decode(r *aper.AperReader) error {
	var err error
	var TypeI_SinglePanel_Group1_r17Present bool
	if TypeI_SinglePanel_Group1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TypeI_SinglePanel_Group2_r17Present bool
	if TypeI_SinglePanel_Group2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TypeI_SinglePanel_ri_RestrictionSTRP_r17Present bool
	if TypeI_SinglePanel_ri_RestrictionSTRP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TypeI_SinglePanel_ri_RestrictionSDM_r17Present bool
	if TypeI_SinglePanel_ri_RestrictionSDM_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if TypeI_SinglePanel_Group1_r17Present {
		ie.TypeI_SinglePanel_Group1_r17 = new(CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17)
		if err = ie.TypeI_SinglePanel_Group1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode TypeI_SinglePanel_Group1_r17", err)
		}
	}
	if TypeI_SinglePanel_Group2_r17Present {
		ie.TypeI_SinglePanel_Group2_r17 = new(CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17)
		if err = ie.TypeI_SinglePanel_Group2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode TypeI_SinglePanel_Group2_r17", err)
		}
	}
	if TypeI_SinglePanel_ri_RestrictionSTRP_r17Present {
		var tmp_bs_TypeI_SinglePanel_ri_RestrictionSTRP_r17 []byte
		var n_TypeI_SinglePanel_ri_RestrictionSTRP_r17 uint
		if tmp_bs_TypeI_SinglePanel_ri_RestrictionSTRP_r17, n_TypeI_SinglePanel_ri_RestrictionSTRP_r17, err = r.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode TypeI_SinglePanel_ri_RestrictionSTRP_r17", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_TypeI_SinglePanel_ri_RestrictionSTRP_r17,
			NumBits: uint64(n_TypeI_SinglePanel_ri_RestrictionSTRP_r17),
		}
		ie.TypeI_SinglePanel_ri_RestrictionSTRP_r17 = &tmp_bitstring
	}
	if TypeI_SinglePanel_ri_RestrictionSDM_r17Present {
		var tmp_bs_TypeI_SinglePanel_ri_RestrictionSDM_r17 []byte
		var n_TypeI_SinglePanel_ri_RestrictionSDM_r17 uint
		if tmp_bs_TypeI_SinglePanel_ri_RestrictionSDM_r17, n_TypeI_SinglePanel_ri_RestrictionSDM_r17, err = r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode TypeI_SinglePanel_ri_RestrictionSDM_r17", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_TypeI_SinglePanel_ri_RestrictionSDM_r17,
			NumBits: uint64(n_TypeI_SinglePanel_ri_RestrictionSDM_r17),
		}
		ie.TypeI_SinglePanel_ri_RestrictionSDM_r17 = &tmp_bitstring
	}
	return nil
}
