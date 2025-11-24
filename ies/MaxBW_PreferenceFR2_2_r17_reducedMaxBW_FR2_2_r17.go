package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MaxBW_PreferenceFR2_2_r17_reducedMaxBW_FR2_2_r17 struct {
	ReducedBW_FR2_2_DL_r17 *ReducedAggregatedBandwidth_r17 `optional`
	ReducedBW_FR2_2_UL_r17 *ReducedAggregatedBandwidth_r17 `optional`
}

func (ie *MaxBW_PreferenceFR2_2_r17_reducedMaxBW_FR2_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ReducedBW_FR2_2_DL_r17 != nil, ie.ReducedBW_FR2_2_UL_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ReducedBW_FR2_2_DL_r17 != nil {
		if err = ie.ReducedBW_FR2_2_DL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedBW_FR2_2_DL_r17", err)
		}
	}
	if ie.ReducedBW_FR2_2_UL_r17 != nil {
		if err = ie.ReducedBW_FR2_2_UL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedBW_FR2_2_UL_r17", err)
		}
	}
	return nil
}

func (ie *MaxBW_PreferenceFR2_2_r17_reducedMaxBW_FR2_2_r17) Decode(r *uper.UperReader) error {
	var err error
	var ReducedBW_FR2_2_DL_r17Present bool
	if ReducedBW_FR2_2_DL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReducedBW_FR2_2_UL_r17Present bool
	if ReducedBW_FR2_2_UL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ReducedBW_FR2_2_DL_r17Present {
		ie.ReducedBW_FR2_2_DL_r17 = new(ReducedAggregatedBandwidth_r17)
		if err = ie.ReducedBW_FR2_2_DL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedBW_FR2_2_DL_r17", err)
		}
	}
	if ReducedBW_FR2_2_UL_r17Present {
		ie.ReducedBW_FR2_2_UL_r17 = new(ReducedAggregatedBandwidth_r17)
		if err = ie.ReducedBW_FR2_2_UL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedBW_FR2_2_UL_r17", err)
		}
	}
	return nil
}
