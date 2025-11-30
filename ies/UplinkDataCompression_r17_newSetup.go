package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UplinkDataCompression_r17_newSetup struct {
	BufferSize_r17 UplinkDataCompression_r17_newSetup_bufferSize_r17  `madatory`
	Dictionary_r17 *UplinkDataCompression_r17_newSetup_dictionary_r17 `optional`
}

func (ie *UplinkDataCompression_r17_newSetup) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Dictionary_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.BufferSize_r17.Encode(w); err != nil {
		return utils.WrapError("Encode BufferSize_r17", err)
	}
	if ie.Dictionary_r17 != nil {
		if err = ie.Dictionary_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dictionary_r17", err)
		}
	}
	return nil
}

func (ie *UplinkDataCompression_r17_newSetup) Decode(r *aper.AperReader) error {
	var err error
	var Dictionary_r17Present bool
	if Dictionary_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.BufferSize_r17.Decode(r); err != nil {
		return utils.WrapError("Decode BufferSize_r17", err)
	}
	if Dictionary_r17Present {
		ie.Dictionary_r17 = new(UplinkDataCompression_r17_newSetup_dictionary_r17)
		if err = ie.Dictionary_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Dictionary_r17", err)
		}
	}
	return nil
}
