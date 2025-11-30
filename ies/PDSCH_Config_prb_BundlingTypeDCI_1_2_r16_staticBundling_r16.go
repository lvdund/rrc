package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_staticBundling_r16 struct {
	BundleSize_r16 *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_staticBundling_r16_bundleSize_r16 `optional`
}

func (ie *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_staticBundling_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.BundleSize_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.BundleSize_r16 != nil {
		if err = ie.BundleSize_r16.Encode(w); err != nil {
			return utils.WrapError("Encode BundleSize_r16", err)
		}
	}
	return nil
}

func (ie *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_staticBundling_r16) Decode(r *aper.AperReader) error {
	var err error
	var BundleSize_r16Present bool
	if BundleSize_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if BundleSize_r16Present {
		ie.BundleSize_r16 = new(PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_staticBundling_r16_bundleSize_r16)
		if err = ie.BundleSize_r16.Decode(r); err != nil {
			return utils.WrapError("Decode BundleSize_r16", err)
		}
	}
	return nil
}
