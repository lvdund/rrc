package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PortIndexFor8Ranks_portIndex4 struct {
	Rank1_4 *PortIndex4  `optional`
	Rank2_4 []PortIndex4 `lb:2,ub:2,optional`
	Rank3_4 []PortIndex4 `lb:3,ub:3,optional`
	Rank4_4 []PortIndex4 `lb:4,ub:4,optional`
}

func (ie *PortIndexFor8Ranks_portIndex4) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Rank1_4 != nil, len(ie.Rank2_4) > 0, len(ie.Rank3_4) > 0, len(ie.Rank4_4) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Rank1_4 != nil {
		if err = ie.Rank1_4.Encode(w); err != nil {
			return utils.WrapError("Encode Rank1_4", err)
		}
	}
	if len(ie.Rank2_4) > 0 {
		tmp_Rank2_4 := utils.NewSequence[*PortIndex4]([]*PortIndex4{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		for _, i := range ie.Rank2_4 {
			tmp_Rank2_4.Value = append(tmp_Rank2_4.Value, &i)
		}
		if err = tmp_Rank2_4.Encode(w); err != nil {
			return utils.WrapError("Encode Rank2_4", err)
		}
	}
	if len(ie.Rank3_4) > 0 {
		tmp_Rank3_4 := utils.NewSequence[*PortIndex4]([]*PortIndex4{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		for _, i := range ie.Rank3_4 {
			tmp_Rank3_4.Value = append(tmp_Rank3_4.Value, &i)
		}
		if err = tmp_Rank3_4.Encode(w); err != nil {
			return utils.WrapError("Encode Rank3_4", err)
		}
	}
	if len(ie.Rank4_4) > 0 {
		tmp_Rank4_4 := utils.NewSequence[*PortIndex4]([]*PortIndex4{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		for _, i := range ie.Rank4_4 {
			tmp_Rank4_4.Value = append(tmp_Rank4_4.Value, &i)
		}
		if err = tmp_Rank4_4.Encode(w); err != nil {
			return utils.WrapError("Encode Rank4_4", err)
		}
	}
	return nil
}

func (ie *PortIndexFor8Ranks_portIndex4) Decode(r *uper.UperReader) error {
	var err error
	var Rank1_4Present bool
	if Rank1_4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rank2_4Present bool
	if Rank2_4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rank3_4Present bool
	if Rank3_4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rank4_4Present bool
	if Rank4_4Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Rank1_4Present {
		ie.Rank1_4 = new(PortIndex4)
		if err = ie.Rank1_4.Decode(r); err != nil {
			return utils.WrapError("Decode Rank1_4", err)
		}
	}
	if Rank2_4Present {
		tmp_Rank2_4 := utils.NewSequence[*PortIndex4]([]*PortIndex4{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		fn_Rank2_4 := func() *PortIndex4 {
			return new(PortIndex4)
		}
		if err = tmp_Rank2_4.Decode(r, fn_Rank2_4); err != nil {
			return utils.WrapError("Decode Rank2_4", err)
		}
		ie.Rank2_4 = []PortIndex4{}
		for _, i := range tmp_Rank2_4.Value {
			ie.Rank2_4 = append(ie.Rank2_4, *i)
		}
	}
	if Rank3_4Present {
		tmp_Rank3_4 := utils.NewSequence[*PortIndex4]([]*PortIndex4{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		fn_Rank3_4 := func() *PortIndex4 {
			return new(PortIndex4)
		}
		if err = tmp_Rank3_4.Decode(r, fn_Rank3_4); err != nil {
			return utils.WrapError("Decode Rank3_4", err)
		}
		ie.Rank3_4 = []PortIndex4{}
		for _, i := range tmp_Rank3_4.Value {
			ie.Rank3_4 = append(ie.Rank3_4, *i)
		}
	}
	if Rank4_4Present {
		tmp_Rank4_4 := utils.NewSequence[*PortIndex4]([]*PortIndex4{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		fn_Rank4_4 := func() *PortIndex4 {
			return new(PortIndex4)
		}
		if err = tmp_Rank4_4.Decode(r, fn_Rank4_4); err != nil {
			return utils.WrapError("Decode Rank4_4", err)
		}
		ie.Rank4_4 = []PortIndex4{}
		for _, i := range tmp_Rank4_4.Value {
			ie.Rank4_4 = append(ie.Rank4_4, *i)
		}
	}
	return nil
}
