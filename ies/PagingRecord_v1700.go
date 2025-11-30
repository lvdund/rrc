package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PagingRecord_v1700 struct {
	PagingCause_r17 *PagingRecord_v1700_pagingCause_r17 `optional`
}

func (ie *PagingRecord_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.PagingCause_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PagingCause_r17 != nil {
		if err = ie.PagingCause_r17.Encode(w); err != nil {
			return utils.WrapError("Encode PagingCause_r17", err)
		}
	}
	return nil
}

func (ie *PagingRecord_v1700) Decode(r *aper.AperReader) error {
	var err error
	var PagingCause_r17Present bool
	if PagingCause_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if PagingCause_r17Present {
		ie.PagingCause_r17 = new(PagingRecord_v1700_pagingCause_r17)
		if err = ie.PagingCause_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PagingCause_r17", err)
		}
	}
	return nil
}
