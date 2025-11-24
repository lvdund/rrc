package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_r16 struct {
	ControlResourceSetId_r16   *ControlResourceSetId_r16               `optional`
	SearchSpaceType_r16        *SearchSpaceExt_r16_searchSpaceType_r16 `optional`
	SearchSpaceGroupIdList_r16 []int64                                 `lb:1,ub:2,e_lb:0,e_ub:1,optional,ext`
	FreqMonitorLocations_r16   *uper.BitString                         `lb:5,ub:5,optional,ext`
}

func (ie *SearchSpaceExt_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ControlResourceSetId_r16 != nil, ie.SearchSpaceType_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ControlResourceSetId_r16 != nil {
		if err = ie.ControlResourceSetId_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ControlResourceSetId_r16", err)
		}
	}
	if ie.SearchSpaceType_r16 != nil {
		if err = ie.SearchSpaceType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SearchSpaceType_r16", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_r16) Decode(r *uper.UperReader) error {
	var err error
	var ControlResourceSetId_r16Present bool
	if ControlResourceSetId_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SearchSpaceType_r16Present bool
	if SearchSpaceType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ControlResourceSetId_r16Present {
		ie.ControlResourceSetId_r16 = new(ControlResourceSetId_r16)
		if err = ie.ControlResourceSetId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ControlResourceSetId_r16", err)
		}
	}
	if SearchSpaceType_r16Present {
		ie.SearchSpaceType_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16)
		if err = ie.SearchSpaceType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SearchSpaceType_r16", err)
		}
	}
	return nil
}
