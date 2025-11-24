package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_partialFreqSounding_r17 struct {
	StartRBIndexFScaling_r17 SRS_Resource_partialFreqSounding_r17_startRBIndexFScaling_r17  `lb:0,ub:1,madatory`
	EnableStartRBHopping_r17 *SRS_Resource_partialFreqSounding_r17_enableStartRBHopping_r17 `optional`
}

func (ie *SRS_Resource_partialFreqSounding_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.EnableStartRBHopping_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.StartRBIndexFScaling_r17.Encode(w); err != nil {
		return utils.WrapError("Encode StartRBIndexFScaling_r17", err)
	}
	if ie.EnableStartRBHopping_r17 != nil {
		if err = ie.EnableStartRBHopping_r17.Encode(w); err != nil {
			return utils.WrapError("Encode EnableStartRBHopping_r17", err)
		}
	}
	return nil
}

func (ie *SRS_Resource_partialFreqSounding_r17) Decode(r *uper.UperReader) error {
	var err error
	var EnableStartRBHopping_r17Present bool
	if EnableStartRBHopping_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.StartRBIndexFScaling_r17.Decode(r); err != nil {
		return utils.WrapError("Decode StartRBIndexFScaling_r17", err)
	}
	if EnableStartRBHopping_r17Present {
		ie.EnableStartRBHopping_r17 = new(SRS_Resource_partialFreqSounding_r17_enableStartRBHopping_r17)
		if err = ie.EnableStartRBHopping_r17.Decode(r); err != nil {
			return utils.WrapError("Decode EnableStartRBHopping_r17", err)
		}
	}
	return nil
}
