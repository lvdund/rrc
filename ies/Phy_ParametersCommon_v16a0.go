package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersCommon_v16a0 struct {
	Srs_PeriodicityAndOffsetExt_r16 *Phy_ParametersCommon_v16a0_srs_PeriodicityAndOffsetExt_r16 `optional`
}

func (ie *Phy_ParametersCommon_v16a0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Srs_PeriodicityAndOffsetExt_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Srs_PeriodicityAndOffsetExt_r16 != nil {
		if err = ie.Srs_PeriodicityAndOffsetExt_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_PeriodicityAndOffsetExt_r16", err)
		}
	}
	return nil
}

func (ie *Phy_ParametersCommon_v16a0) Decode(r *uper.UperReader) error {
	var err error
	var Srs_PeriodicityAndOffsetExt_r16Present bool
	if Srs_PeriodicityAndOffsetExt_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Srs_PeriodicityAndOffsetExt_r16Present {
		ie.Srs_PeriodicityAndOffsetExt_r16 = new(Phy_ParametersCommon_v16a0_srs_PeriodicityAndOffsetExt_r16)
		if err = ie.Srs_PeriodicityAndOffsetExt_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_PeriodicityAndOffsetExt_r16", err)
		}
	}
	return nil
}
