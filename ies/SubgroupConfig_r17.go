package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SubgroupConfig_r17 struct {
	SubgroupsNumPerPO_r17   int64  `lb:1,ub:maxNrofPagingSubgroups_r17,madatory`
	SubgroupsNumForUEID_r17 *int64 `lb:1,ub:maxNrofPagingSubgroups_r17,optional`
}

func (ie *SubgroupConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SubgroupsNumForUEID_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.SubgroupsNumPerPO_r17, &uper.Constraint{Lb: 1, Ub: maxNrofPagingSubgroups_r17}, false); err != nil {
		return utils.WrapError("WriteInteger SubgroupsNumPerPO_r17", err)
	}
	if ie.SubgroupsNumForUEID_r17 != nil {
		if err = w.WriteInteger(*ie.SubgroupsNumForUEID_r17, &uper.Constraint{Lb: 1, Ub: maxNrofPagingSubgroups_r17}, false); err != nil {
			return utils.WrapError("Encode SubgroupsNumForUEID_r17", err)
		}
	}
	return nil
}

func (ie *SubgroupConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var SubgroupsNumForUEID_r17Present bool
	if SubgroupsNumForUEID_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_SubgroupsNumPerPO_r17 int64
	if tmp_int_SubgroupsNumPerPO_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofPagingSubgroups_r17}, false); err != nil {
		return utils.WrapError("ReadInteger SubgroupsNumPerPO_r17", err)
	}
	ie.SubgroupsNumPerPO_r17 = tmp_int_SubgroupsNumPerPO_r17
	if SubgroupsNumForUEID_r17Present {
		var tmp_int_SubgroupsNumForUEID_r17 int64
		if tmp_int_SubgroupsNumForUEID_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofPagingSubgroups_r17}, false); err != nil {
			return utils.WrapError("Decode SubgroupsNumForUEID_r17", err)
		}
		ie.SubgroupsNumForUEID_r17 = &tmp_int_SubgroupsNumForUEID_r17
	}
	return nil
}
