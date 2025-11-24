package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PortIndexFor8Ranks_portIndex8 struct {
	Rank1_8 *PortIndex8  `optional`
	Rank2_8 []PortIndex8 `lb:2,ub:2,optional`
	Rank3_8 []PortIndex8 `lb:3,ub:3,optional`
	Rank4_8 []PortIndex8 `lb:4,ub:4,optional`
	Rank5_8 []PortIndex8 `lb:5,ub:5,optional`
	Rank6_8 []PortIndex8 `lb:6,ub:6,optional`
	Rank7_8 []PortIndex8 `lb:7,ub:7,optional`
	Rank8_8 []PortIndex8 `lb:8,ub:8,optional`
}

func (ie *PortIndexFor8Ranks_portIndex8) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Rank1_8 != nil, len(ie.Rank2_8) > 0, len(ie.Rank3_8) > 0, len(ie.Rank4_8) > 0, len(ie.Rank5_8) > 0, len(ie.Rank6_8) > 0, len(ie.Rank7_8) > 0, len(ie.Rank8_8) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Rank1_8 != nil {
		if err = ie.Rank1_8.Encode(w); err != nil {
			return utils.WrapError("Encode Rank1_8", err)
		}
	}
	if len(ie.Rank2_8) > 0 {
		tmp_Rank2_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		for _, i := range ie.Rank2_8 {
			tmp_Rank2_8.Value = append(tmp_Rank2_8.Value, &i)
		}
		if err = tmp_Rank2_8.Encode(w); err != nil {
			return utils.WrapError("Encode Rank2_8", err)
		}
	}
	if len(ie.Rank3_8) > 0 {
		tmp_Rank3_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		for _, i := range ie.Rank3_8 {
			tmp_Rank3_8.Value = append(tmp_Rank3_8.Value, &i)
		}
		if err = tmp_Rank3_8.Encode(w); err != nil {
			return utils.WrapError("Encode Rank3_8", err)
		}
	}
	if len(ie.Rank4_8) > 0 {
		tmp_Rank4_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		for _, i := range ie.Rank4_8 {
			tmp_Rank4_8.Value = append(tmp_Rank4_8.Value, &i)
		}
		if err = tmp_Rank4_8.Encode(w); err != nil {
			return utils.WrapError("Encode Rank4_8", err)
		}
	}
	if len(ie.Rank5_8) > 0 {
		tmp_Rank5_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 5, Ub: 5}, false)
		for _, i := range ie.Rank5_8 {
			tmp_Rank5_8.Value = append(tmp_Rank5_8.Value, &i)
		}
		if err = tmp_Rank5_8.Encode(w); err != nil {
			return utils.WrapError("Encode Rank5_8", err)
		}
	}
	if len(ie.Rank6_8) > 0 {
		tmp_Rank6_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 6, Ub: 6}, false)
		for _, i := range ie.Rank6_8 {
			tmp_Rank6_8.Value = append(tmp_Rank6_8.Value, &i)
		}
		if err = tmp_Rank6_8.Encode(w); err != nil {
			return utils.WrapError("Encode Rank6_8", err)
		}
	}
	if len(ie.Rank7_8) > 0 {
		tmp_Rank7_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 7, Ub: 7}, false)
		for _, i := range ie.Rank7_8 {
			tmp_Rank7_8.Value = append(tmp_Rank7_8.Value, &i)
		}
		if err = tmp_Rank7_8.Encode(w); err != nil {
			return utils.WrapError("Encode Rank7_8", err)
		}
	}
	if len(ie.Rank8_8) > 0 {
		tmp_Rank8_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 8, Ub: 8}, false)
		for _, i := range ie.Rank8_8 {
			tmp_Rank8_8.Value = append(tmp_Rank8_8.Value, &i)
		}
		if err = tmp_Rank8_8.Encode(w); err != nil {
			return utils.WrapError("Encode Rank8_8", err)
		}
	}
	return nil
}

func (ie *PortIndexFor8Ranks_portIndex8) Decode(r *uper.UperReader) error {
	var err error
	var Rank1_8Present bool
	if Rank1_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rank2_8Present bool
	if Rank2_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rank3_8Present bool
	if Rank3_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rank4_8Present bool
	if Rank4_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rank5_8Present bool
	if Rank5_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rank6_8Present bool
	if Rank6_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rank7_8Present bool
	if Rank7_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rank8_8Present bool
	if Rank8_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Rank1_8Present {
		ie.Rank1_8 = new(PortIndex8)
		if err = ie.Rank1_8.Decode(r); err != nil {
			return utils.WrapError("Decode Rank1_8", err)
		}
	}
	if Rank2_8Present {
		tmp_Rank2_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		fn_Rank2_8 := func() *PortIndex8 {
			return new(PortIndex8)
		}
		if err = tmp_Rank2_8.Decode(r, fn_Rank2_8); err != nil {
			return utils.WrapError("Decode Rank2_8", err)
		}
		ie.Rank2_8 = []PortIndex8{}
		for _, i := range tmp_Rank2_8.Value {
			ie.Rank2_8 = append(ie.Rank2_8, *i)
		}
	}
	if Rank3_8Present {
		tmp_Rank3_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		fn_Rank3_8 := func() *PortIndex8 {
			return new(PortIndex8)
		}
		if err = tmp_Rank3_8.Decode(r, fn_Rank3_8); err != nil {
			return utils.WrapError("Decode Rank3_8", err)
		}
		ie.Rank3_8 = []PortIndex8{}
		for _, i := range tmp_Rank3_8.Value {
			ie.Rank3_8 = append(ie.Rank3_8, *i)
		}
	}
	if Rank4_8Present {
		tmp_Rank4_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		fn_Rank4_8 := func() *PortIndex8 {
			return new(PortIndex8)
		}
		if err = tmp_Rank4_8.Decode(r, fn_Rank4_8); err != nil {
			return utils.WrapError("Decode Rank4_8", err)
		}
		ie.Rank4_8 = []PortIndex8{}
		for _, i := range tmp_Rank4_8.Value {
			ie.Rank4_8 = append(ie.Rank4_8, *i)
		}
	}
	if Rank5_8Present {
		tmp_Rank5_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 5, Ub: 5}, false)
		fn_Rank5_8 := func() *PortIndex8 {
			return new(PortIndex8)
		}
		if err = tmp_Rank5_8.Decode(r, fn_Rank5_8); err != nil {
			return utils.WrapError("Decode Rank5_8", err)
		}
		ie.Rank5_8 = []PortIndex8{}
		for _, i := range tmp_Rank5_8.Value {
			ie.Rank5_8 = append(ie.Rank5_8, *i)
		}
	}
	if Rank6_8Present {
		tmp_Rank6_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 6, Ub: 6}, false)
		fn_Rank6_8 := func() *PortIndex8 {
			return new(PortIndex8)
		}
		if err = tmp_Rank6_8.Decode(r, fn_Rank6_8); err != nil {
			return utils.WrapError("Decode Rank6_8", err)
		}
		ie.Rank6_8 = []PortIndex8{}
		for _, i := range tmp_Rank6_8.Value {
			ie.Rank6_8 = append(ie.Rank6_8, *i)
		}
	}
	if Rank7_8Present {
		tmp_Rank7_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 7, Ub: 7}, false)
		fn_Rank7_8 := func() *PortIndex8 {
			return new(PortIndex8)
		}
		if err = tmp_Rank7_8.Decode(r, fn_Rank7_8); err != nil {
			return utils.WrapError("Decode Rank7_8", err)
		}
		ie.Rank7_8 = []PortIndex8{}
		for _, i := range tmp_Rank7_8.Value {
			ie.Rank7_8 = append(ie.Rank7_8, *i)
		}
	}
	if Rank8_8Present {
		tmp_Rank8_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 8, Ub: 8}, false)
		fn_Rank8_8 := func() *PortIndex8 {
			return new(PortIndex8)
		}
		if err = tmp_Rank8_8.Decode(r, fn_Rank8_8); err != nil {
			return utils.WrapError("Decode Rank8_8", err)
		}
		ie.Rank8_8 = []PortIndex8{}
		for _, i := range tmp_Rank8_8.Value {
			ie.Rank8_8 = append(ie.Rank8_8, *i)
		}
	}
	return nil
}
