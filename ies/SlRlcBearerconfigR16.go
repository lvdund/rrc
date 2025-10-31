package ies

// SL-RLC-BearerConfig-r16 ::= SEQUENCE
// Extensible
type SlRlcBearerconfigR16 struct {
	SlRlcBearerconfigindexR16    SlRlcBearerconfigindexR16
	SlServedradiobearerR16       *SlrbUuConfigindexR16
	SlRlcConfigR16               *SlRlcConfigR16
	SlMacLogicalchannelconfigR16 *SlLogicalchannelconfigR16
}
