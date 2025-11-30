package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CC_Group_r17 struct {
	ServCellIndexLower_r17  ServCellIndex                     `madatory`
	ServCellIndexHigher_r17 *ServCellIndex                    `optional`
	DefaultDC_Location_r17  DefaultDC_Location_r17            `madatory`
	OffsetToDefault_r17     *CC_Group_r17_offsetToDefault_r17 `lb:1,ub:maxNrofReqComDC_Location_r17,optional`
}

func (ie *CC_Group_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ServCellIndexHigher_r17 != nil, ie.OffsetToDefault_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ServCellIndexLower_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ServCellIndexLower_r17", err)
	}
	if ie.ServCellIndexHigher_r17 != nil {
		if err = ie.ServCellIndexHigher_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ServCellIndexHigher_r17", err)
		}
	}
	if err = ie.DefaultDC_Location_r17.Encode(w); err != nil {
		return utils.WrapError("Encode DefaultDC_Location_r17", err)
	}
	if ie.OffsetToDefault_r17 != nil {
		if err = ie.OffsetToDefault_r17.Encode(w); err != nil {
			return utils.WrapError("Encode OffsetToDefault_r17", err)
		}
	}
	return nil
}

func (ie *CC_Group_r17) Decode(r *aper.AperReader) error {
	var err error
	var ServCellIndexHigher_r17Present bool
	if ServCellIndexHigher_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OffsetToDefault_r17Present bool
	if OffsetToDefault_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ServCellIndexLower_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ServCellIndexLower_r17", err)
	}
	if ServCellIndexHigher_r17Present {
		ie.ServCellIndexHigher_r17 = new(ServCellIndex)
		if err = ie.ServCellIndexHigher_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ServCellIndexHigher_r17", err)
		}
	}
	if err = ie.DefaultDC_Location_r17.Decode(r); err != nil {
		return utils.WrapError("Decode DefaultDC_Location_r17", err)
	}
	if OffsetToDefault_r17Present {
		ie.OffsetToDefault_r17 = new(CC_Group_r17_offsetToDefault_r17)
		if err = ie.OffsetToDefault_r17.Decode(r); err != nil {
			return utils.WrapError("Decode OffsetToDefault_r17", err)
		}
	}
	return nil
}
