package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_ResourceExt_v1610_format_v1610_Choice_nothing uint64 = iota
	PUCCH_ResourceExt_v1610_format_v1610_Choice_Interlace1_v1610
	PUCCH_ResourceExt_v1610_format_v1610_Choice_Occ_v1610
)

type PUCCH_ResourceExt_v1610_format_v1610 struct {
	Choice           uint64
	Interlace1_v1610 int64 `lb:0,ub:9,madatory`
	Occ_v1610        *PUCCH_ResourceExt_v1610_format_v1610_occ_v1610
}

func (ie *PUCCH_ResourceExt_v1610_format_v1610) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_ResourceExt_v1610_format_v1610_Choice_Interlace1_v1610:
		if err = w.WriteInteger(int64(ie.Interlace1_v1610), &aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Interlace1_v1610", err)
		}
	case PUCCH_ResourceExt_v1610_format_v1610_Choice_Occ_v1610:
		if err = ie.Occ_v1610.Encode(w); err != nil {
			err = utils.WrapError("Encode Occ_v1610", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PUCCH_ResourceExt_v1610_format_v1610) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_ResourceExt_v1610_format_v1610_Choice_Interlace1_v1610:
		var tmp_int_Interlace1_v1610 int64
		if tmp_int_Interlace1_v1610, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Interlace1_v1610", err)
		}
		ie.Interlace1_v1610 = tmp_int_Interlace1_v1610
	case PUCCH_ResourceExt_v1610_format_v1610_Choice_Occ_v1610:
		ie.Occ_v1610 = new(PUCCH_ResourceExt_v1610_format_v1610_occ_v1610)
		if err = ie.Occ_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Occ_v1610", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
