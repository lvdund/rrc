package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MulticastRLC_BearerConfig_r17 struct {
	ServedMBS_RadioBearer_r17 MRB_Identity_r17                                `madatory`
	IsPTM_Entity_r17          *MulticastRLC_BearerConfig_r17_isPTM_Entity_r17 `optional`
}

func (ie *MulticastRLC_BearerConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.IsPTM_Entity_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ServedMBS_RadioBearer_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ServedMBS_RadioBearer_r17", err)
	}
	if ie.IsPTM_Entity_r17 != nil {
		if err = ie.IsPTM_Entity_r17.Encode(w); err != nil {
			return utils.WrapError("Encode IsPTM_Entity_r17", err)
		}
	}
	return nil
}

func (ie *MulticastRLC_BearerConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var IsPTM_Entity_r17Present bool
	if IsPTM_Entity_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ServedMBS_RadioBearer_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ServedMBS_RadioBearer_r17", err)
	}
	if IsPTM_Entity_r17Present {
		ie.IsPTM_Entity_r17 = new(MulticastRLC_BearerConfig_r17_isPTM_Entity_r17)
		if err = ie.IsPTM_Entity_r17.Decode(r); err != nil {
			return utils.WrapError("Decode IsPTM_Entity_r17", err)
		}
	}
	return nil
}
