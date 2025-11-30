package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MaxBW_PreferenceFR2_2_r17 struct {
	ReducedMaxBW_FR2_2_r17 *MaxBW_PreferenceFR2_2_r17_reducedMaxBW_FR2_2_r17 `optional`
}

func (ie *MaxBW_PreferenceFR2_2_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ReducedMaxBW_FR2_2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ReducedMaxBW_FR2_2_r17 != nil {
		if err = ie.ReducedMaxBW_FR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedMaxBW_FR2_2_r17", err)
		}
	}
	return nil
}

func (ie *MaxBW_PreferenceFR2_2_r17) Decode(r *aper.AperReader) error {
	var err error
	var ReducedMaxBW_FR2_2_r17Present bool
	if ReducedMaxBW_FR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ReducedMaxBW_FR2_2_r17Present {
		ie.ReducedMaxBW_FR2_2_r17 = new(MaxBW_PreferenceFR2_2_r17_reducedMaxBW_FR2_2_r17)
		if err = ie.ReducedMaxBW_FR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedMaxBW_FR2_2_r17", err)
		}
	}
	return nil
}
