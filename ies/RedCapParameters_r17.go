package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RedCapParameters_r17 struct {
	SupportOfRedCap_r17       *RedCapParameters_r17_supportOfRedCap_r17       `optional`
	SupportOf16DRB_RedCap_r17 *RedCapParameters_r17_supportOf16DRB_RedCap_r17 `optional`
}

func (ie *RedCapParameters_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SupportOfRedCap_r17 != nil, ie.SupportOf16DRB_RedCap_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SupportOfRedCap_r17 != nil {
		if err = ie.SupportOfRedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SupportOfRedCap_r17", err)
		}
	}
	if ie.SupportOf16DRB_RedCap_r17 != nil {
		if err = ie.SupportOf16DRB_RedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SupportOf16DRB_RedCap_r17", err)
		}
	}
	return nil
}

func (ie *RedCapParameters_r17) Decode(r *aper.AperReader) error {
	var err error
	var SupportOfRedCap_r17Present bool
	if SupportOfRedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportOf16DRB_RedCap_r17Present bool
	if SupportOf16DRB_RedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SupportOfRedCap_r17Present {
		ie.SupportOfRedCap_r17 = new(RedCapParameters_r17_supportOfRedCap_r17)
		if err = ie.SupportOfRedCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SupportOfRedCap_r17", err)
		}
	}
	if SupportOf16DRB_RedCap_r17Present {
		ie.SupportOf16DRB_RedCap_r17 = new(RedCapParameters_r17_supportOf16DRB_RedCap_r17)
		if err = ie.SupportOf16DRB_RedCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SupportOf16DRB_RedCap_r17", err)
		}
	}
	return nil
}
