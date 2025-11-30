package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CFRA_TwoStep_r16 struct {
	OccasionsTwoStepRA_r16 *CFRA_TwoStep_r16_occasionsTwoStepRA_r16 `optional`
	MsgA_CFRA_PUSCH_r16    MsgA_PUSCH_Resource_r16                  `madatory`
	MsgA_TransMax_r16      *CFRA_TwoStep_r16_msgA_TransMax_r16      `optional`
	ResourcesTwoStep_r16   CFRA_TwoStep_r16_resourcesTwoStep_r16    `lb:1,ub:maxRA_SSB_Resources,madatory`
}

func (ie *CFRA_TwoStep_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.OccasionsTwoStepRA_r16 != nil, ie.MsgA_TransMax_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.OccasionsTwoStepRA_r16 != nil {
		if err = ie.OccasionsTwoStepRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OccasionsTwoStepRA_r16", err)
		}
	}
	if err = ie.MsgA_CFRA_PUSCH_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MsgA_CFRA_PUSCH_r16", err)
	}
	if ie.MsgA_TransMax_r16 != nil {
		if err = ie.MsgA_TransMax_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_TransMax_r16", err)
		}
	}
	if err = ie.ResourcesTwoStep_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ResourcesTwoStep_r16", err)
	}
	return nil
}

func (ie *CFRA_TwoStep_r16) Decode(r *aper.AperReader) error {
	var err error
	var OccasionsTwoStepRA_r16Present bool
	if OccasionsTwoStepRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_TransMax_r16Present bool
	if MsgA_TransMax_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if OccasionsTwoStepRA_r16Present {
		ie.OccasionsTwoStepRA_r16 = new(CFRA_TwoStep_r16_occasionsTwoStepRA_r16)
		if err = ie.OccasionsTwoStepRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OccasionsTwoStepRA_r16", err)
		}
	}
	if err = ie.MsgA_CFRA_PUSCH_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MsgA_CFRA_PUSCH_r16", err)
	}
	if MsgA_TransMax_r16Present {
		ie.MsgA_TransMax_r16 = new(CFRA_TwoStep_r16_msgA_TransMax_r16)
		if err = ie.MsgA_TransMax_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_TransMax_r16", err)
		}
	}
	if err = ie.ResourcesTwoStep_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ResourcesTwoStep_r16", err)
	}
	return nil
}
