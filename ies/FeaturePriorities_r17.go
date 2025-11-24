package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeaturePriorities_r17 struct {
	RedCapPriority_r17            *FeaturePriority_r17 `optional`
	SlicingPriority_r17           *FeaturePriority_r17 `optional`
	Msg3_Repetitions_Priority_r17 *FeaturePriority_r17 `optional`
	Sdt_Priority_r17              *FeaturePriority_r17 `optional`
}

func (ie *FeaturePriorities_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.RedCapPriority_r17 != nil, ie.SlicingPriority_r17 != nil, ie.Msg3_Repetitions_Priority_r17 != nil, ie.Sdt_Priority_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.RedCapPriority_r17 != nil {
		if err = ie.RedCapPriority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RedCapPriority_r17", err)
		}
	}
	if ie.SlicingPriority_r17 != nil {
		if err = ie.SlicingPriority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SlicingPriority_r17", err)
		}
	}
	if ie.Msg3_Repetitions_Priority_r17 != nil {
		if err = ie.Msg3_Repetitions_Priority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Msg3_Repetitions_Priority_r17", err)
		}
	}
	if ie.Sdt_Priority_r17 != nil {
		if err = ie.Sdt_Priority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sdt_Priority_r17", err)
		}
	}
	return nil
}

func (ie *FeaturePriorities_r17) Decode(r *uper.UperReader) error {
	var err error
	var RedCapPriority_r17Present bool
	if RedCapPriority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SlicingPriority_r17Present bool
	if SlicingPriority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Msg3_Repetitions_Priority_r17Present bool
	if Msg3_Repetitions_Priority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sdt_Priority_r17Present bool
	if Sdt_Priority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if RedCapPriority_r17Present {
		ie.RedCapPriority_r17 = new(FeaturePriority_r17)
		if err = ie.RedCapPriority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RedCapPriority_r17", err)
		}
	}
	if SlicingPriority_r17Present {
		ie.SlicingPriority_r17 = new(FeaturePriority_r17)
		if err = ie.SlicingPriority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SlicingPriority_r17", err)
		}
	}
	if Msg3_Repetitions_Priority_r17Present {
		ie.Msg3_Repetitions_Priority_r17 = new(FeaturePriority_r17)
		if err = ie.Msg3_Repetitions_Priority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Msg3_Repetitions_Priority_r17", err)
		}
	}
	if Sdt_Priority_r17Present {
		ie.Sdt_Priority_r17 = new(FeaturePriority_r17)
		if err = ie.Sdt_Priority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sdt_Priority_r17", err)
		}
	}
	return nil
}
