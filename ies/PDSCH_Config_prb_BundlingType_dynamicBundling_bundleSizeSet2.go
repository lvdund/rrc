package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_Config_prb_BundlingType_dynamicBundling_bundleSizeSet2_Enum_n4       aper.Enumerated = 0
	PDSCH_Config_prb_BundlingType_dynamicBundling_bundleSizeSet2_Enum_wideband aper.Enumerated = 1
)

type PDSCH_Config_prb_BundlingType_dynamicBundling_bundleSizeSet2 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PDSCH_Config_prb_BundlingType_dynamicBundling_bundleSizeSet2) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PDSCH_Config_prb_BundlingType_dynamicBundling_bundleSizeSet2", err)
	}
	return nil
}

func (ie *PDSCH_Config_prb_BundlingType_dynamicBundling_bundleSizeSet2) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PDSCH_Config_prb_BundlingType_dynamicBundling_bundleSizeSet2", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
