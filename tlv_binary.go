package benchmarks

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

const (
	TLVID        uint8 = 0
	TLVKind      uint8 = 1
	TLVPubKey    uint8 = 3
	TLVCreatedAt uint8 = 4
	TLVContent   uint8 = 2
	TLVTags      uint8 = 5
	TLVSig       uint8 = 6
)

func decodeEventTLV(data []byte) *Event {
	var evt Event
	curr := 0
	for {
		t, v := readTLVEntry(data[curr:])
		if v == nil {
			break
		}

		switch t {
		case TLVID:
			evt.ID = hex.EncodeToString(v)
		case TLVPubKey:
			evt.PubKey = hex.EncodeToString(v)
		case TLVSig:
			evt.Sig = hex.EncodeToString(v)
		case TLVContent:
			evt.Content = string(v)
		case TLVKind:
			evt.Kind = int(binary.BigEndian.Uint32(v))
		case TLVCreatedAt:
			evt.CreatedAt = Timestamp(binary.BigEndian.Uint64(v))
		case TLVTags:
			evt.Tags = decodeTags(v)
		default:
			// ignore
		}

		curr = curr + 3 + len(v)
	}

	return &evt
}

func readTLVEntry(data []byte) (typ uint8, value []byte) {
	if len(data) < 2 {
		return 0, nil
	}

	typ = data[0]
	lenb := data[1:3]
	n := binary.BigEndian.Uint16(lenb)
	value = data[3 : 3+n]
	return
}

func writeTLVEntry(buf *bytes.Buffer, typ uint8, value []byte) {
	buf.WriteByte(typ)
	lenb := make([]byte, 2)
	binary.BigEndian.PutUint16(lenb, uint16(len(value)))
	buf.Write(lenb)
	buf.Write(value)
}

func encodeEventTLV(evt *Event) []byte {
	buf := bytes.NewBuffer([]byte{})

	id, _ := hex.DecodeString(evt.ID)
	writeTLVEntry(buf, TLVID, id)
	pubkey, _ := hex.DecodeString(evt.PubKey)
	writeTLVEntry(buf, TLVPubKey, pubkey)
	sig, _ := hex.DecodeString(evt.Sig)
	writeTLVEntry(buf, TLVSig, sig)
	writeTLVEntry(buf, TLVContent, []byte(evt.Content))
	kindb := make([]byte, 4)
	binary.BigEndian.PutUint32(kindb, uint32(evt.Kind))
	writeTLVEntry(buf, TLVKind, kindb)
	cab := make([]byte, 8)
	binary.BigEndian.PutUint64(cab, uint64(evt.CreatedAt))
	writeTLVEntry(buf, TLVCreatedAt, cab)
	tagsb := encodeTags(evt.Tags)
	writeTLVEntry(buf, TLVTags, tagsb)

	return buf.Bytes()
}

func encodeTags(tags Tags) []byte {
	buf := bytes.NewBuffer([]byte{})
	for _, tag := range tags {
		buf.WriteByte(uint8(len(tag)))
		for _, item := range tag {
			buf.WriteByte(uint8(len(item)))
			buf.WriteString(item)
		}
	}
	return buf.Bytes()
}

func decodeTags(buf []byte) Tags {
	tags := make(Tags, 0, 4)
	reader := bytes.NewReader(buf)
	for {
		n, err := reader.ReadByte()
		if err != nil {
			return tags
		}
		tag := make(Tag, n)
		for i := 0; i < int(n); i++ {
			itemSize, _ := reader.ReadByte()
			if itemSize == 0 {
				continue
			}
			data := make([]byte, itemSize)
			reader.Read(data)
			tag = append(tag, string(itemSize))
		}
		tags = append(tags, tag)
	}
}
