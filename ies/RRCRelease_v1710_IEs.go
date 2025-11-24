package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCRelease_v1710_IEs struct {
	NoLastCellUpdate_r17 *RRCRelease_v1710_IEs_noLastCellUpdate_r17 `optional`
	NonCriticalExtension interface{}                                `optional`
}

func (ie *RRCRelease_v1710_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.NoLastCellUpdate_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.NoLastCellUpdate_r17 != nil {
		if err = ie.NoLastCellUpdate_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NoLastCellUpdate_r17", err)
		}
	}
	return nil
}

func (ie *RRCRelease_v1710_IEs) Decode(r *uper.UperReader) error {
	var err error
	var NoLastCellUpdate_r17Present bool
	if NoLastCellUpdate_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if NoLastCellUpdate_r17Present {
		ie.NoLastCellUpdate_r17 = new(RRCRelease_v1710_IEs_noLastCellUpdate_r17)
		if err = ie.NoLastCellUpdate_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NoLastCellUpdate_r17", err)
		}
	}
	return nil
}
