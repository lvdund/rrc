package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PortIndexFor8Ranks_portIndex2 struct {
	Rank1_2 *PortIndex2  `optional`
	Rank2_2 []PortIndex2 `lb:2,ub:2,optional`
}

func (ie *PortIndexFor8Ranks_portIndex2) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Rank1_2 != nil, len(ie.Rank2_2) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Rank1_2 != nil {
		if err = ie.Rank1_2.Encode(w); err != nil {
			return utils.WrapError("Encode Rank1_2", err)
		}
	}
	if len(ie.Rank2_2) > 0 {
		tmp_Rank2_2 := utils.NewSequence[*PortIndex2]([]*PortIndex2{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		for _, i := range ie.Rank2_2 {
			tmp_Rank2_2.Value = append(tmp_Rank2_2.Value, &i)
		}
		if err = tmp_Rank2_2.Encode(w); err != nil {
			return utils.WrapError("Encode Rank2_2", err)
		}
	}
	return nil
}

func (ie *PortIndexFor8Ranks_portIndex2) Decode(r *uper.UperReader) error {
	var err error
	var Rank1_2Present bool
	if Rank1_2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rank2_2Present bool
	if Rank2_2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Rank1_2Present {
		ie.Rank1_2 = new(PortIndex2)
		if err = ie.Rank1_2.Decode(r); err != nil {
			return utils.WrapError("Decode Rank1_2", err)
		}
	}
	if Rank2_2Present {
		tmp_Rank2_2 := utils.NewSequence[*PortIndex2]([]*PortIndex2{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		fn_Rank2_2 := func() *PortIndex2 {
			return new(PortIndex2)
		}
		if err = tmp_Rank2_2.Decode(r, fn_Rank2_2); err != nil {
			return utils.WrapError("Decode Rank2_2", err)
		}
		ie.Rank2_2 = []PortIndex2{}
		for _, i := range tmp_Rank2_2.Value {
			ie.Rank2_2 = append(ie.Rank2_2, *i)
		}
	}
	return nil
}
