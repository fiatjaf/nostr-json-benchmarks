/* globals Deno */

import {assertObjectMatch} from 'https://deno.land/std@0.183.0/testing/asserts.ts'

const eventString = `{"id":"ae1fc7154296569d87ca4663f6bdf448c217d1590d28c85d158557b8b43b4d69","pubkey":"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798","sig":"94e10947814b1ebe38af42300ecd90c7642763896c4f69506ae97bfdf54eec3c0c21df96b7d95daa74ff3d414b1d758ee95fc258125deebc31df0c6ba9396a51","created_at":1683660344,"nson":"2805000b0203000100400005040001004000000014","kind":30023,"content":"hello hello","tags":[["e","b6de44a9dd47d1c000f795ea0453046914f44ba7d5e369608b04867a575ea83e","reply"],["p","c26f7b252cea77a5b94f42b1a4771021be07d4df766407e47738605f7e3ab774","","wss://relay.damus.io"]]}`
const eventBinary = hexToBytes(
  `ae1fc7154296569d87ca4663f6bdf448c217d1590d28c85d158557b8b43b4d6979be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f8179894e10947814b1ebe38af42300ecd90c7642763896c4f69506ae97bfdf54eec3c0c21df96b7d95daa74ff3d414b1d758ee95fc258125deebc31df0c6ba9396a51645a9e387547000b68656c6c6f2068656c6c6f0203000165000040623664653434613964643437643163303030663739356561303435333034363931346634346261376435653336393630386230343836376135373565613833650000057265706c790004000170000040633236663762323532636561373761356239346634326231613437373130323162653037643464663736363430376534373733383630356637653361623737340000000000147773733a2f2f72656c61792e64616d75732e696f`
)

let utf8 = new TextDecoder('utf-8')

Deno.test('nson and leaner match JSON.parse', () => {
  let jp = JSON.parse(eventString)
  delete jp.nson

  let [dn] = decodeNson(eventString)
  assertObjectMatch(dn, jp, 'nson does not match')

  let [ld] = leanerDecode(eventBinary)
  assertObjectMatch(
    ld,
    {
      ...jp,
      id: Uint8Array.from(hexToBytes(jp.id)),
      pubkey: Uint8Array.from(hexToBytes(jp.pubkey)),
      sig: Uint8Array.from(hexToBytes(jp.sig))
    },
    'binary does not match'
  )
})

Deno.bench('JSON.parse', () => {
  JSON.parse(eventString)
})

Deno.bench('decodeNson', () => {
  decodeNson(eventString)
})

Deno.bench('leanerDecode', () => {
  leanerDecode(eventBinary)
})

function decodeNson(data) {
  let evt = {}
  let err = null

  try {
    evt.id = data.substring(7, 7 + 64)
    evt.pubkey = data.substring(83, 83 + 64)
    evt.sig = data.substring(156, 156 + 128)
    let ts = parseInt(data.substring(299, 299 + 10), 10)
    evt.created_at = ts

    let nsonSizeBytes = hexToBytes(data.substring(318, 318 + 2))
    let nsonSize = nsonSizeBytes[0]
    let nsonDescriptors = hexToBytes(data.substring(320, 320 + nsonSize))

    let kindChars = nsonDescriptors[0]
    let kindStart = 320 + nsonSize + 9
    evt.kind = parseInt(data.substring(kindStart, kindStart + kindChars))

    let contentChars = (nsonDescriptors[1] << 8) | nsonDescriptors[2]
    let contentStart = kindStart + kindChars + 12
    evt.content = data.substring(contentStart, contentStart + contentChars)

    let nTags = nsonDescriptors[3]
    evt.tags = new Array(nTags)
    let tagsStart = contentStart + contentChars + 9

    let nsonIndex = 3
    let tagsIndex = tagsStart
    for (let t = 0; t < nTags; t++) {
      nsonIndex++
      tagsIndex += 1
      let nItems = nsonDescriptors[nsonIndex]
      let tag = new Array(nItems)
      for (let n = 0; n < nItems; n++) {
        nsonIndex++
        let itemStart = tagsIndex + 2
        let itemChars =
          (nsonDescriptors[nsonIndex] << 8) | nsonDescriptors[nsonIndex + 1]
        nsonIndex++
        tag[n] = data.substring(itemStart, itemStart + itemChars)
        tagsIndex = itemStart + itemChars + 1
      }
      tagsIndex += 1
      evt.tags[t] = tag
    }
  } catch (e) {
    err = new Error('failed to decode nson: ' + e)
    console.log(err)
  }

  return [evt, err]
}

function leanerDecode(data) {
  let evt = {}
  let err = null
  let buffer = Uint8Array.from(data).buffer

  try {
    evt.id = new Uint8Array(buffer, 0, 32)
    evt.pubkey = new Uint8Array(buffer, 32, 32)
    evt.sig = new Uint8Array(buffer, 64, 64)
    evt.created_at =
      (data[128] << 24) | (data[129] << 16) | (data[130] << 8) | data[131]
    evt.kind = (data[132] << 8) | data[133]
    let contentLength = (data[134] << 8) | data[135]
    evt.content = utf8.decode(new Uint8Array(buffer, 136, contentLength))

    let curr = 136 + contentLength
    let ntags = data[curr]
    evt.tags = new Array(ntags)

    for (let t = 0; t < evt.tags.length; t++) {
      curr = curr + 1
      let nItems = data[curr]
      let tag = new Array(nItems)
      for (let i = 0; i < tag.length; i++) {
        curr = curr + 1
        let itemSize = (data[curr] << 8) | data[curr + 1]
        let itemStart = curr + 2
        let item = utf8.decode(new Uint8Array(buffer, itemStart, itemSize))
        tag[i] = item
        curr = itemStart + itemSize
      }
      evt.tags[t] = tag
    }
  } catch (e) {
    err = new Error('failed to decode leaner: ' + e)
    console.log(err)
  }

  return [evt, err]
}

function hexToBytes(hex) {
  let bytes = []
  for (let i = 0; i < hex.length; i += 2) {
    bytes.push(parseInt(hex.substr(i, 2), 16))
  }
  return bytes
}
