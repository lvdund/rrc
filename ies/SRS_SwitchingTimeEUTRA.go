package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_SwitchingTimeEUTRA struct {
	SwitchingTimeDL *SRS_SwitchingTimeEUTRA_switchingTimeDL `optional`
	SwitchingTimeUL *SRS_SwitchingTimeEUTRA_switchingTimeUL `optional`
}

func (ie *SRS_SwitchingTimeEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SwitchingTimeDL != nil, ie.SwitchingTimeUL != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SwitchingTimeDL != nil {
		if err = ie.SwitchingTimeDL.Encode(w); err != nil {
			return utils.WrapError("Encode SwitchingTimeDL", err)
		}
	}
	if ie.SwitchingTimeUL != nil {
		if err = ie.SwitchingTimeUL.Encode(w); err != nil {
			return utils.WrapError("Encode SwitchingTimeUL", err)
		}
	}
	return nil
}

func (ie *SRS_SwitchingTimeEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var SwitchingTimeDLPresent bool
	if SwitchingTimeDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SwitchingTimeULPresent bool
	if SwitchingTimeULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SwitchingTimeDLPresent {
		ie.SwitchingTimeDL = new(SRS_SwitchingTimeEUTRA_switchingTimeDL)
		if err = ie.SwitchingTimeDL.Decode(r); err != nil {
			return utils.WrapError("Decode SwitchingTimeDL", err)
		}
	}
	if SwitchingTimeULPresent {
		ie.SwitchingTimeUL = new(SRS_SwitchingTimeEUTRA_switchingTimeUL)
		if err = ie.SwitchingTimeUL.Decode(r); err != nil {
			return utils.WrapError("Decode SwitchingTimeUL", err)
		}
	}
	return nil
}
