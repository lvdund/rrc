package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB16_r17 struct {
	FreqPriorityListSlicing_r17 *FreqPriorityListSlicing_r17 `optional`
	LateNonCriticalExtension    *[]byte                      `optional`
}

func (ie *SIB16_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.FreqPriorityListSlicing_r17 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FreqPriorityListSlicing_r17 != nil {
		if err = ie.FreqPriorityListSlicing_r17.Encode(w); err != nil {
			return utils.WrapError("Encode FreqPriorityListSlicing_r17", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB16_r17) Decode(r *aper.AperReader) error {
	var err error
	var FreqPriorityListSlicing_r17Present bool
	if FreqPriorityListSlicing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FreqPriorityListSlicing_r17Present {
		ie.FreqPriorityListSlicing_r17 = new(FreqPriorityListSlicing_r17)
		if err = ie.FreqPriorityListSlicing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode FreqPriorityListSlicing_r17", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
