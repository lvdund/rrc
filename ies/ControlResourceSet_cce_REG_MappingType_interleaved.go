package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ControlResourceSet_cce_REG_MappingType_interleaved struct {
	Reg_BundleSize  ControlResourceSet_cce_REG_MappingType_interleaved_reg_BundleSize  `madatory`
	InterleaverSize ControlResourceSet_cce_REG_MappingType_interleaved_interleaverSize `madatory`
	ShiftIndex      *int64                                                             `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional`
}

func (ie *ControlResourceSet_cce_REG_MappingType_interleaved) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ShiftIndex != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Reg_BundleSize.Encode(w); err != nil {
		return utils.WrapError("Encode Reg_BundleSize", err)
	}
	if err = ie.InterleaverSize.Encode(w); err != nil {
		return utils.WrapError("Encode InterleaverSize", err)
	}
	if ie.ShiftIndex != nil {
		if err = w.WriteInteger(*ie.ShiftIndex, &uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Encode ShiftIndex", err)
		}
	}
	return nil
}

func (ie *ControlResourceSet_cce_REG_MappingType_interleaved) Decode(r *uper.UperReader) error {
	var err error
	var ShiftIndexPresent bool
	if ShiftIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Reg_BundleSize.Decode(r); err != nil {
		return utils.WrapError("Decode Reg_BundleSize", err)
	}
	if err = ie.InterleaverSize.Decode(r); err != nil {
		return utils.WrapError("Decode InterleaverSize", err)
	}
	if ShiftIndexPresent {
		var tmp_int_ShiftIndex int64
		if tmp_int_ShiftIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Decode ShiftIndex", err)
		}
		ie.ShiftIndex = &tmp_int_ShiftIndex
	}
	return nil
}
