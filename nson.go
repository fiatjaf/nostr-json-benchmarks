package benchmarks

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
)

/*
           nson size
             kind chars
               content chars
                   number of tags (let's say it's two)
                     number of items on the first tag (let's say it's three)
                       number of chars on the first item
                         number of chars on the second item
                           number of chars on the third item
                             number of items on the second tag (let's say it's two)
                               number of chars on the first item
                                 number of chars on the second item
   "nson":"xxkkccccttnn112233nn1122"
*/

func decodeNson(data string) *Event {
	evt := Event{}

	// static fields
	evt.ID = data[7 : 7+64]
	evt.PubKey = data[83 : 83+64]
	evt.Sig = data[156 : 156+128]
	ts, _ := strconv.ParseInt(data[299:299+10], 10, 64)
	evt.CreatedAt = Timestamp(ts)

	// nson values
	nsonSizeBytes, _ := hex.DecodeString(data[318 : 318+2])
	nsonSize := int(nsonSizeBytes[0])
	nsonDescriptors, _ := hex.DecodeString(data[320 : 320+nsonSize])
	fmt.Println(nsonSizeBytes, nsonSize, nsonDescriptors)

	// dynamic fields
	kindChars := int(nsonDescriptors[0])
	kindStart := 320 + nsonSize + 9 // len(`","kind":`)
	evt.Kind, _ = strconv.Atoi(data[kindStart : kindStart+kindChars])

	contentChars := int(binary.BigEndian.Uint16(nsonDescriptors[1:3]))
	contentStart := kindStart + kindChars + 12 // len(`,"content":"`)
	evt.Content = data[contentStart : contentStart+contentChars]

	nTags := int(nsonDescriptors[3])
	evt.Tags = make(Tags, nTags)
	tagsStart := contentStart + contentChars + 9 // len(`","tags":`)
	fmt.Println(data[tagsStart-1 : tagsStart+10])

	nsonIndex := 3
	tagsIndex := tagsStart
	for t := 0; t < nTags; t++ {
		nsonIndex++
		tagsIndex += 1 // len(`[`) or len(`,`)
		nItems := int(nsonDescriptors[nsonIndex])
		fmt.Println(t, nItems, tagsIndex, data[tagsIndex:tagsIndex+10])
		tag := make(Tag, nItems)
		for n := 0; n < nItems; n++ {
			nsonIndex++
			itemStart := tagsIndex + 2 // len(`["`) or len(`,"`)
			itemChars := int(nsonDescriptors[nsonIndex])
			fmt.Println(" ", n, itemChars, tagsIndex, data[tagsIndex+2:tagsIndex+2+itemChars+1])
			tag[n] = data[itemStart : itemStart+itemChars]
			tagsIndex = itemStart + itemChars + 1 // len(`"`)
		}
		tagsIndex += 1 // len(`]`)
		evt.Tags[t] = tag
	}

	return &evt
}

func encodeNson(*Event) []byte {
	return nil
}
