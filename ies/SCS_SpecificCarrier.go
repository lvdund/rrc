package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SCS_SpecificCarrier struct {
	OffsetToCarrier         int64             `lb:0,ub:2199,madatory`
	SubcarrierSpacing       SubcarrierSpacing `madatory`
	CarrierBandwidth        int64             `lb:1,ub:maxNrofPhysicalResourceBlocks,madatory`
	TxDirectCurrentLocation *int64            `lb:0,ub:4095,optional,ext-1`
}

func (ie *SCS_SpecificCarrier) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.TxDirectCurrentLocation != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.OffsetToCarrier, &aper.Constraint{Lb: 0, Ub: 2199}, false); err != nil {
		return utils.WrapError("WriteInteger OffsetToCarrier", err)
	}
	if err = ie.SubcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode SubcarrierSpacing", err)
	}
	if err = w.WriteInteger(ie.CarrierBandwidth, &aper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks}, false); err != nil {
		return utils.WrapError("WriteInteger CarrierBandwidth", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.TxDirectCurrentLocation != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SCS_SpecificCarrier", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.TxDirectCurrentLocation != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode TxDirectCurrentLocation optional
			if ie.TxDirectCurrentLocation != nil {
				if err = extWriter.WriteInteger(*ie.TxDirectCurrentLocation, &aper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
					return utils.WrapError("Encode TxDirectCurrentLocation", err)
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

func (ie *SCS_SpecificCarrier) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_OffsetToCarrier int64
	if tmp_int_OffsetToCarrier, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2199}, false); err != nil {
		return utils.WrapError("ReadInteger OffsetToCarrier", err)
	}
	ie.OffsetToCarrier = tmp_int_OffsetToCarrier
	if err = ie.SubcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode SubcarrierSpacing", err)
	}
	var tmp_int_CarrierBandwidth int64
	if tmp_int_CarrierBandwidth, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks}, false); err != nil {
		return utils.WrapError("ReadInteger CarrierBandwidth", err)
	}
	ie.CarrierBandwidth = tmp_int_CarrierBandwidth

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

			TxDirectCurrentLocationPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode TxDirectCurrentLocation optional
			if TxDirectCurrentLocationPresent {
				var tmp_int_TxDirectCurrentLocation int64
				if tmp_int_TxDirectCurrentLocation, err = extReader.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
					return utils.WrapError("Decode TxDirectCurrentLocation", err)
				}
				ie.TxDirectCurrentLocation = &tmp_int_TxDirectCurrentLocation
			}
		}
	}
	return nil
}
