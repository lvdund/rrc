package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_Config_prb_BundlingType_dynamicBundling struct {
	BundleSizeSet1 *PDSCH_Config_prb_BundlingType_dynamicBundling_bundleSizeSet1 `optional`
	BundleSizeSet2 *PDSCH_Config_prb_BundlingType_dynamicBundling_bundleSizeSet2 `optional`
}

func (ie *PDSCH_Config_prb_BundlingType_dynamicBundling) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.BundleSizeSet1 != nil, ie.BundleSizeSet2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.BundleSizeSet1 != nil {
		if err = ie.BundleSizeSet1.Encode(w); err != nil {
			return utils.WrapError("Encode BundleSizeSet1", err)
		}
	}
	if ie.BundleSizeSet2 != nil {
		if err = ie.BundleSizeSet2.Encode(w); err != nil {
			return utils.WrapError("Encode BundleSizeSet2", err)
		}
	}
	return nil
}

func (ie *PDSCH_Config_prb_BundlingType_dynamicBundling) Decode(r *uper.UperReader) error {
	var err error
	var BundleSizeSet1Present bool
	if BundleSizeSet1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BundleSizeSet2Present bool
	if BundleSizeSet2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if BundleSizeSet1Present {
		ie.BundleSizeSet1 = new(PDSCH_Config_prb_BundlingType_dynamicBundling_bundleSizeSet1)
		if err = ie.BundleSizeSet1.Decode(r); err != nil {
			return utils.WrapError("Decode BundleSizeSet1", err)
		}
	}
	if BundleSizeSet2Present {
		ie.BundleSizeSet2 = new(PDSCH_Config_prb_BundlingType_dynamicBundling_bundleSizeSet2)
		if err = ie.BundleSizeSet2.Decode(r); err != nil {
			return utils.WrapError("Decode BundleSizeSet2", err)
		}
	}
	return nil
}
