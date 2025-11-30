package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultUTRA_FDD_r16 struct {
	PhysCellId_r16 PhysCellIdUTRA_FDD_r16                 `madatory`
	MeasResult_r16 *MeasResultUTRA_FDD_r16_measResult_r16 `lb:-5,ub:91,optional`
}

func (ie *MeasResultUTRA_FDD_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasResult_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PhysCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId_r16", err)
	}
	if ie.MeasResult_r16 != nil {
		if err = ie.MeasResult_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResult_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultUTRA_FDD_r16) Decode(r *aper.AperReader) error {
	var err error
	var MeasResult_r16Present bool
	if MeasResult_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PhysCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId_r16", err)
	}
	if MeasResult_r16Present {
		ie.MeasResult_r16 = new(MeasResultUTRA_FDD_r16_measResult_r16)
		if err = ie.MeasResult_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResult_r16", err)
		}
	}
	return nil
}
