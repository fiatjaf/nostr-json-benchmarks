package benchmarks

import (
	"encoding/binary"
	"fmt"
)

func leanerDecode(data []byte) (evt *EventBinary, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to decode nson: %v", r)
		}
	}()

	evt = &EventBinary{}

	evt.ID = data[0:32]
	evt.PubKey = data[32:64]
	evt.Sig = data[64:128]
	evt.CreatedAt = Timestamp(binary.BigEndian.Uint32(data[128:132]))
	evt.Kind = int(binary.BigEndian.Uint16(data[132:134]))
	contentLength := int(binary.BigEndian.Uint16(data[134:136]))
	evt.Content = string(data[136 : 136+contentLength])

	curr := 136 + contentLength
	ntags := int(data[curr])
	evt.Tags = make(Tags, ntags)

	for t := range evt.Tags {
		curr = curr + 1
		nItems := int(data[curr])
		tag := make(Tag, nItems)
		for i := range tag {
			curr = curr + 1
			itemSize := int(binary.BigEndian.Uint16(data[curr : curr+2]))
			itemStart := curr + 2
			itemEnd := itemStart + itemSize
			item := string(data[itemStart:itemEnd])
			tag[i] = item
			curr = itemEnd
		}
		evt.Tags[t] = tag
	}

	return evt, err
}

func leanerEncode(evt *EventBinary) []byte {
	content := []byte(evt.Content)
	buf := make([]byte, 32+32+64+4+2+2+len(content)+65500)
	copy(buf[0:32], evt.ID)
	copy(buf[32:64], evt.PubKey)
	copy(buf[64:128], evt.Sig)
	binary.BigEndian.PutUint32(buf[128:132], uint32(evt.CreatedAt))
	binary.BigEndian.PutUint16(buf[132:134], uint16(evt.Kind))
	binary.BigEndian.PutUint16(buf[134:136], uint16(len(content)))
	copy(buf[136:], content)

	curr := 136 + len(content)
	buf[curr] = uint8(len(evt.Tags))
	for _, tag := range evt.Tags {
		curr++
		buf[curr] = uint8(len(tag))
		for _, item := range tag {
			curr++
			itemb := []byte(item)
			itemSize := len(itemb)
			binary.BigEndian.PutUint16(buf[curr:curr+2], uint16(itemSize))
			itemEnd := curr + 2 + itemSize
			copy(buf[curr+2:itemEnd], itemb)
			curr = itemEnd
		}
	}

	return buf
}
