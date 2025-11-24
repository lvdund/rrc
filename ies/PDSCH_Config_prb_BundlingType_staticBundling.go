package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_Config_prb_BundlingType_staticBundling struct {
	BundleSize *PDSCH_Config_prb_BundlingType_staticBundling_bundleSize `optional`
}

func (ie *PDSCH_Config_prb_BundlingType_staticBundling) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.BundleSize != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.BundleSize != nil {
		if err = ie.BundleSize.Encode(w); err != nil {
			return utils.WrapError("Encode BundleSize", err)
		}
	}
	return nil
}

func (ie *PDSCH_Config_prb_BundlingType_staticBundling) Decode(r *uper.UperReader) error {
	var err error
	var BundleSizePresent bool
	if BundleSizePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if BundleSizePresent {
		ie.BundleSize = new(PDSCH_Config_prb_BundlingType_staticBundling_bundleSize)
		if err = ie.BundleSize.Decode(r); err != nil {
			return utils.WrapError("Decode BundleSize", err)
		}
	}
	return nil
}
