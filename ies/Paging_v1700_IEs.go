package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Paging_v1700_IEs struct {
	PagingRecordList_v1700 *PagingRecordList_v1700 `optional`
	PagingGroupList_r17    *PagingGroupList_r17    `optional`
	NonCriticalExtension   interface{}             `optional`
}

func (ie *Paging_v1700_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.PagingRecordList_v1700 != nil, ie.PagingGroupList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PagingRecordList_v1700 != nil {
		if err = ie.PagingRecordList_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode PagingRecordList_v1700", err)
		}
	}
	if ie.PagingGroupList_r17 != nil {
		if err = ie.PagingGroupList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode PagingGroupList_r17", err)
		}
	}
	return nil
}

func (ie *Paging_v1700_IEs) Decode(r *aper.AperReader) error {
	var err error
	var PagingRecordList_v1700Present bool
	if PagingRecordList_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PagingGroupList_r17Present bool
	if PagingGroupList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if PagingRecordList_v1700Present {
		ie.PagingRecordList_v1700 = new(PagingRecordList_v1700)
		if err = ie.PagingRecordList_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode PagingRecordList_v1700", err)
		}
	}
	if PagingGroupList_r17Present {
		ie.PagingGroupList_r17 = new(PagingGroupList_r17)
		if err = ie.PagingGroupList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PagingGroupList_r17", err)
		}
	}
	return nil
}
