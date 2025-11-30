package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplinkPerCC_v1540 struct {
	Mimo_NonCB_PUSCH *FeatureSetUplinkPerCC_v1540_mimo_NonCB_PUSCH `lb:1,ub:4,optional`
}

func (ie *FeatureSetUplinkPerCC_v1540) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Mimo_NonCB_PUSCH != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Mimo_NonCB_PUSCH != nil {
		if err = ie.Mimo_NonCB_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode Mimo_NonCB_PUSCH", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplinkPerCC_v1540) Decode(r *aper.AperReader) error {
	var err error
	var Mimo_NonCB_PUSCHPresent bool
	if Mimo_NonCB_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Mimo_NonCB_PUSCHPresent {
		ie.Mimo_NonCB_PUSCH = new(FeatureSetUplinkPerCC_v1540_mimo_NonCB_PUSCH)
		if err = ie.Mimo_NonCB_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode Mimo_NonCB_PUSCH", err)
		}
	}
	return nil
}
