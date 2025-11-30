package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_ConfigCommon struct {
	Pdsch_TimeDomainAllocationList *PDSCH_TimeDomainResourceAllocationList `optional`
}

func (ie *PDSCH_ConfigCommon) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Pdsch_TimeDomainAllocationList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pdsch_TimeDomainAllocationList != nil {
		if err = ie.Pdsch_TimeDomainAllocationList.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_TimeDomainAllocationList", err)
		}
	}
	return nil
}

func (ie *PDSCH_ConfigCommon) Decode(r *aper.AperReader) error {
	var err error
	var Pdsch_TimeDomainAllocationListPresent bool
	if Pdsch_TimeDomainAllocationListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Pdsch_TimeDomainAllocationListPresent {
		ie.Pdsch_TimeDomainAllocationList = new(PDSCH_TimeDomainResourceAllocationList)
		if err = ie.Pdsch_TimeDomainAllocationList.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_TimeDomainAllocationList", err)
		}
	}
	return nil
}
