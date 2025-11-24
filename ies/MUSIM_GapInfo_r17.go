package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MUSIM_GapInfo_r17 struct {
	Musim_Starting_SFN_AndSubframe_r17 *MUSIM_Starting_SFN_AndSubframe_r17                 `optional`
	Musim_GapLength_r17                *MUSIM_GapInfo_r17_musim_GapLength_r17              `optional`
	Musim_GapRepetitionAndOffset_r17   *MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17 `lb:0,ub:19,optional`
}

func (ie *MUSIM_GapInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Musim_Starting_SFN_AndSubframe_r17 != nil, ie.Musim_GapLength_r17 != nil, ie.Musim_GapRepetitionAndOffset_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Musim_Starting_SFN_AndSubframe_r17 != nil {
		if err = ie.Musim_Starting_SFN_AndSubframe_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Musim_Starting_SFN_AndSubframe_r17", err)
		}
	}
	if ie.Musim_GapLength_r17 != nil {
		if err = ie.Musim_GapLength_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Musim_GapLength_r17", err)
		}
	}
	if ie.Musim_GapRepetitionAndOffset_r17 != nil {
		if err = ie.Musim_GapRepetitionAndOffset_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Musim_GapRepetitionAndOffset_r17", err)
		}
	}
	return nil
}

func (ie *MUSIM_GapInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var Musim_Starting_SFN_AndSubframe_r17Present bool
	if Musim_Starting_SFN_AndSubframe_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Musim_GapLength_r17Present bool
	if Musim_GapLength_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Musim_GapRepetitionAndOffset_r17Present bool
	if Musim_GapRepetitionAndOffset_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Musim_Starting_SFN_AndSubframe_r17Present {
		ie.Musim_Starting_SFN_AndSubframe_r17 = new(MUSIM_Starting_SFN_AndSubframe_r17)
		if err = ie.Musim_Starting_SFN_AndSubframe_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Musim_Starting_SFN_AndSubframe_r17", err)
		}
	}
	if Musim_GapLength_r17Present {
		ie.Musim_GapLength_r17 = new(MUSIM_GapInfo_r17_musim_GapLength_r17)
		if err = ie.Musim_GapLength_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Musim_GapLength_r17", err)
		}
	}
	if Musim_GapRepetitionAndOffset_r17Present {
		ie.Musim_GapRepetitionAndOffset_r17 = new(MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17)
		if err = ie.Musim_GapRepetitionAndOffset_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Musim_GapRepetitionAndOffset_r17", err)
		}
	}
	return nil
}
