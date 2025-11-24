package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_CC_SetIndex struct {
	Cc_SetIndex         *int64 `lb:0,ub:3,optional`
	Cc_IndexInOneCC_Set *int64 `lb:0,ub:7,optional`
}

func (ie *SRS_CC_SetIndex) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Cc_SetIndex != nil, ie.Cc_IndexInOneCC_Set != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Cc_SetIndex != nil {
		if err = w.WriteInteger(*ie.Cc_SetIndex, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Encode Cc_SetIndex", err)
		}
	}
	if ie.Cc_IndexInOneCC_Set != nil {
		if err = w.WriteInteger(*ie.Cc_IndexInOneCC_Set, &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Encode Cc_IndexInOneCC_Set", err)
		}
	}
	return nil
}

func (ie *SRS_CC_SetIndex) Decode(r *uper.UperReader) error {
	var err error
	var Cc_SetIndexPresent bool
	if Cc_SetIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Cc_IndexInOneCC_SetPresent bool
	if Cc_IndexInOneCC_SetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Cc_SetIndexPresent {
		var tmp_int_Cc_SetIndex int64
		if tmp_int_Cc_SetIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode Cc_SetIndex", err)
		}
		ie.Cc_SetIndex = &tmp_int_Cc_SetIndex
	}
	if Cc_IndexInOneCC_SetPresent {
		var tmp_int_Cc_IndexInOneCC_Set int64
		if tmp_int_Cc_IndexInOneCC_Set, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode Cc_IndexInOneCC_Set", err)
		}
		ie.Cc_IndexInOneCC_Set = &tmp_int_Cc_IndexInOneCC_Set
	}
	return nil
}
