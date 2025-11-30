package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RedCap_ConfigCommonSIB_r17 struct {
	HalfDuplexRedCapAllowed_r17 *RedCap_ConfigCommonSIB_r17_halfDuplexRedCapAllowed_r17 `optional`
	CellBarredRedCap_r17        *RedCap_ConfigCommonSIB_r17_cellBarredRedCap_r17        `optional`
}

func (ie *RedCap_ConfigCommonSIB_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.HalfDuplexRedCapAllowed_r17 != nil, ie.CellBarredRedCap_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.HalfDuplexRedCapAllowed_r17 != nil {
		if err = ie.HalfDuplexRedCapAllowed_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HalfDuplexRedCapAllowed_r17", err)
		}
	}
	if ie.CellBarredRedCap_r17 != nil {
		if err = ie.CellBarredRedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CellBarredRedCap_r17", err)
		}
	}
	return nil
}

func (ie *RedCap_ConfigCommonSIB_r17) Decode(r *aper.AperReader) error {
	var err error
	var HalfDuplexRedCapAllowed_r17Present bool
	if HalfDuplexRedCapAllowed_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CellBarredRedCap_r17Present bool
	if CellBarredRedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if HalfDuplexRedCapAllowed_r17Present {
		ie.HalfDuplexRedCapAllowed_r17 = new(RedCap_ConfigCommonSIB_r17_halfDuplexRedCapAllowed_r17)
		if err = ie.HalfDuplexRedCapAllowed_r17.Decode(r); err != nil {
			return utils.WrapError("Decode HalfDuplexRedCapAllowed_r17", err)
		}
	}
	if CellBarredRedCap_r17Present {
		ie.CellBarredRedCap_r17 = new(RedCap_ConfigCommonSIB_r17_cellBarredRedCap_r17)
		if err = ie.CellBarredRedCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CellBarredRedCap_r17", err)
		}
	}
	return nil
}
