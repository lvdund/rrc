package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasTiming struct {
	FrequencyAndTiming *MeasTiming_frequencyAndTiming `optional`
	Ssb_ToMeasure      *SSB_ToMeasure                 `optional,ext-1`
	PhysCellId         *PhysCellId                    `optional,ext-1`
}

func (ie *MeasTiming) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Ssb_ToMeasure != nil || ie.PhysCellId != nil
	preambleBits := []bool{hasExtensions, ie.FrequencyAndTiming != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FrequencyAndTiming != nil {
		if err = ie.FrequencyAndTiming.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyAndTiming", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Ssb_ToMeasure != nil || ie.PhysCellId != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasTiming", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Ssb_ToMeasure != nil, ie.PhysCellId != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ssb_ToMeasure optional
			if ie.Ssb_ToMeasure != nil {
				if err = ie.Ssb_ToMeasure.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ssb_ToMeasure", err)
				}
			}
			// encode PhysCellId optional
			if ie.PhysCellId != nil {
				if err = ie.PhysCellId.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PhysCellId", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *MeasTiming) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var FrequencyAndTimingPresent bool
	if FrequencyAndTimingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FrequencyAndTimingPresent {
		ie.FrequencyAndTiming = new(MeasTiming_frequencyAndTiming)
		if err = ie.FrequencyAndTiming.Decode(r); err != nil {
			return utils.WrapError("Decode FrequencyAndTiming", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Ssb_ToMeasurePresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PhysCellIdPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ssb_ToMeasure optional
			if Ssb_ToMeasurePresent {
				ie.Ssb_ToMeasure = new(SSB_ToMeasure)
				if err = ie.Ssb_ToMeasure.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ssb_ToMeasure", err)
				}
			}
			// decode PhysCellId optional
			if PhysCellIdPresent {
				ie.PhysCellId = new(PhysCellId)
				if err = ie.PhysCellId.Decode(extReader); err != nil {
					return utils.WrapError("Decode PhysCellId", err)
				}
			}
		}
	}
	return nil
}
