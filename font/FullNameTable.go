package font

import (
	"encoding/binary"
	"io"
	"strings"
)

type FullNameTable struct {
	fontName string
}

func newFullNameTable(reader io.ReadSeeker, tableDirectory []byte, position int) (*FullNameTable, error) {
	table := &FullNameTable{}

	intOffset := ReadULong(tableDirectory, position+8)
	intLength := ReadULong(tableDirectory, position+12)

	data := make([]byte, intLength)
	reader.Seek(int64(intOffset), io.SeekStart)
	io.ReadFull(reader, data)

	dataStart := ReadUShort(data, 4)
	headerStart := 6
	headerEnd := ReadUShort(data, 2) * 12

	var fontName string

	for i := headerStart; i < headerEnd; i += 12 {
		if ReadUShort(data, i+6) == 4 { //4 is the Name ID for Full font name
			if ReadUShort(data, i) == 3 && ReadUShort(data, i+2) == 1 && ReadUShort(data, i+4) == 0x0409 { // Platform ID 3, Encoding ID 1, Language ID 0x0409 (Windows, Unicode, US English)
				offset := dataStart + ReadUShort(data, i+10)
				length := ReadUShort(data, i+8)
				fontName = decodeUTF16BE(data[offset : offset+length])
				break
			}
		}
	}

	// Fallback: platform ID 3, encoding ID 0, language ID 0x0409
	if fontName == "" {
		for i := headerStart; i < headerEnd; i += 12 {
			if ReadUShort(data, i+6) == 4 { //4 is the Name ID for Full font name
				if ReadUShort(data, i) == 3 && ReadUShort(data, i+2) == 0 && ReadUShort(data, i+4) == 0x0409 { //3 for PlatForm ID, 0 for Encoding ID and 0x0409 Language ID for English(United States)
					offset := dataStart + ReadUShort(data, i+10)
					length := ReadUShort(data, i+8)
					fontName = decodeUTF16BE(data[offset : offset+length])
					break
				}
			}
		}
	}

	// Clean font name
	fontName = strings.ReplaceAll(fontName, " ", "")
	fontName = strings.ReplaceAll(fontName, "-", "")
	table.fontName = fontName

	return table, nil
}

func ReadULong(data []byte, index int) int {
	intReturn := int(data[index])
	index++
	intReturn *= 0x100

	intReturn += int(data[index])
	index++
	intReturn *= 0x100

	intReturn += int(data[index])
	index++
	intReturn *= 0x100

	intReturn += int(data[index])
	return intReturn
}

func ReadUShort(data []byte, index int) int {
	return int(data[index])<<8 | int(data[index+1])
}

func ReadUShort1(data []byte, index int) int {
	return int(data[index]) | int(data[index+1])<<8
}

func ReadShort(data []byte, index int) int {
	return int(data[index])<<8 | int(data[index+1])
}

func ReadByte(data []byte, index int) int {
	return int(data[index])
}

func ReadFixed(data []byte, index int) float32 {
	intInteger := int(data[index])
	if intInteger > 127 {
		intInteger -= 256
	}
	index++
	intInteger *= 0x100

	intInteger += int(data[index])
	index++
	intInteger *= 0x100

	intInteger += int(data[index])
	index++
	intInteger *= 0x100

	intInteger += int(data[index])
	return float32(intInteger / 0x10000)

}

func ReadFWord(data []byte, index int) int {
	intReturn := int(data[index])
	if intReturn > 127 {
		intReturn -= 256
	}
	intReturn = intReturn<<8 + int(data[index+1])
	return intReturn
}

func decodeUTF16BE(b []byte) string {
	u16s := make([]uint16, len(b)/2)
	for i := 0; i < len(u16s); i++ {
		u16s[i] = binary.BigEndian.Uint16(b[i*2:])
	}
	return string(runeSlice(u16s))
}

func runeSlice(u16s []uint16) []rune {
	runes := make([]rune, len(u16s))
	for i, v := range u16s {
		runes[i] = rune(v)
	}
	return runes
}
