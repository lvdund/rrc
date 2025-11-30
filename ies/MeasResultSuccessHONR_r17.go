package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultSuccessHONR_r17 struct {
	MeasResult_r17 *MeasResultSuccessHONR_r17_measResult_r17 `optional`
}

func (ie *MeasResultSuccessHONR_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasResult_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasResult_r17 != nil {
		if err = ie.MeasResult_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResult_r17", err)
		}
	}
	return nil
}

func (ie *MeasResultSuccessHONR_r17) Decode(r *aper.AperReader) error {
	var err error
	var MeasResult_r17Present bool
	if MeasResult_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasResult_r17Present {
		ie.MeasResult_r17 = new(MeasResultSuccessHONR_r17_measResult_r17)
		if err = ie.MeasResult_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResult_r17", err)
		}
	}
	return nil
}
