# rrc-gen

Go types and PER (Packed Encoding Rules) codec for the 3GPP NR (5G) **Radio Resource Control (RRC)** protocol,
generated from the ASN.1 specification **3GPP TS 38.331 V19.2.0 (2026-03)**.

## RRC PDU messages

Every RRC PDU / top-level message from TS 38.331 is generated as a Go struct in package `ies`.
Each struct exposes `Encode(e *per.Encoder) error` and `Decode(d *per.Decoder) error`.

| RRC message (ASN.1)             | Go type                  | File                                          |
|---------------------------------|--------------------------|-----------------------------------------------|
| `BCCH-BCH-Message`              | `BCCH_BCH_Message`       | [ies/BCCH_BCH_Message.gen.go](ies/BCCH_BCH_Message.gen.go) |
| `BCCH-DL-SCH-Message`           | `BCCH_DL_SCH_Message`    | [ies/BCCH_DL_SCH_Message.gen.go](ies/BCCH_DL_SCH_Message.gen.go) |
| `DL-CCCH-Message`               | `DL_CCCH_Message`        | [ies/DL_CCCH_Message.gen.go](ies/DL_CCCH_Message.gen.go) |
| `DL-DCCH-Message`               | `DL_DCCH_Message`        | [ies/DL_DCCH_Message.gen.go](ies/DL_DCCH_Message.gen.go) |
| `PCCH-Message`                  | `PCCH_Message`           | [ies/PCCH_Message.gen.go](ies/PCCH_Message.gen.go) |
| `UL-CCCH-Message`               | `UL_CCCH_Message`        | [ies/UL_CCCH_Message.gen.go](ies/UL_CCCH_Message.gen.go) |
| `UL-CCCH1-Message`              | `UL_CCCH1_Message`       | [ies/UL_CCCH1_Message.gen.go](ies/UL_CCCH1_Message.gen.go) |
| `UL-DCCH-Message`               | `UL_DCCH_Message`        | [ies/UL_DCCH_Message.gen.go](ies/UL_DCCH_Message.gen.go) |

In addition to these PDUs, package `ies` also contains every referenced Information Element
(MIB, SIB1, RRCReconfiguration, RRCSetupRequest, CellGroupConfig, Measurements, …) — see the
[`ies/`](ies/) directory for the full list.

## How to use

### Decode an RRC PDU (PER bytes → Go struct)

```go
package main

import (
    "fmt"

    "github.com/lvdund/asn1go/per"
    "github.com/lvdund/rrc/ies"
)

func main() {
    raw := []byte{ /* PER-encoded BCCH-BCH-Message bytes here */ }

    dec := per.NewDecoder(raw)
    var msg ies.BCCH_BCH_Message
    if err := msg.Decode(dec); err != nil {
        panic(err)
    }

    fmt.Printf("MIB = %#v\n", msg.Message)
}
```

### Encode an RRC PDU (Go struct → PER bytes)

```go
enc := per.NewEncoder()
msg := ies.BCCH_BCH_Message{
    Message: ies.BCCH_BCH_MessageType{
        // MIB fields...
    },
}
if err := msg.Encode(enc); err != nil {
    panic(err)
}
bytes := enc.Bytes()
```

### Generic helper

The `Encode`/`Decode` methods follow the same contract for every generated type, so you can write
a small generic wrapper once and reuse it for all RRC messages:

```go
type Codec interface {
    Encode(e *per.Encoder) error
    Decode(d *per.Decoder) error
}

func Encode[T Codec](msg T) ([]byte, error) {
    e := per.NewEncoder()
    if err := msg.Encode(e); err != nil {
        return nil, err
    }
    return e.Bytes(), nil
}

func Decode[T Codec](raw []byte) (T, error) {
    var msg T
    if err := msg.Decode(per.NewDecoder(raw)); err != nil {
        return msg, err
    }
    return msg, nil
}
```
## License

See [LICENSE](LICENSE).