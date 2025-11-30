package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_v1700_searchSpaceType_r17 struct {
	Common_r17 *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17 `optional`
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Common_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Common_r17 != nil {
		if err = ie.Common_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Common_r17", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17) Decode(r *aper.AperReader) error {
	var err error
	var Common_r17Present bool
	if Common_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Common_r17Present {
		ie.Common_r17 = new(SearchSpaceExt_v1700_searchSpaceType_r17_common_r17)
		if err = ie.Common_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Common_r17", err)
		}
	}
	return nil
}
