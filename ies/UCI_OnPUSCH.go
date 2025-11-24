package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UCI_OnPUSCH struct {
	BetaOffsets *UCI_OnPUSCH_betaOffsets `lb:4,ub:4,optional`
	Scaling     UCI_OnPUSCH_scaling      `madatory`
}

func (ie *UCI_OnPUSCH) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.BetaOffsets != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.BetaOffsets != nil {
		if err = ie.BetaOffsets.Encode(w); err != nil {
			return utils.WrapError("Encode BetaOffsets", err)
		}
	}
	if err = ie.Scaling.Encode(w); err != nil {
		return utils.WrapError("Encode Scaling", err)
	}
	return nil
}

func (ie *UCI_OnPUSCH) Decode(r *uper.UperReader) error {
	var err error
	var BetaOffsetsPresent bool
	if BetaOffsetsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if BetaOffsetsPresent {
		ie.BetaOffsets = new(UCI_OnPUSCH_betaOffsets)
		if err = ie.BetaOffsets.Decode(r); err != nil {
			return utils.WrapError("Decode BetaOffsets", err)
		}
	}
	if err = ie.Scaling.Decode(r); err != nil {
		return utils.WrapError("Decode Scaling", err)
	}
	return nil
}
