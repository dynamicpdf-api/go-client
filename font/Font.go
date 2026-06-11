package font

import (
	"encoding/binary"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/dynamicpdf-api/go-client/v2/resource"
	"github.com/google/uuid"
)

// Represents font.
type Font struct {
	//Resource resource.Resource `json:"-"`

	FontResource resource.FontResource `json:"-"`

	ResourceName string `json:"resourceName,omitempty"`

	Embed bool `json:"embed"`

	Subset bool `json:"subset"`

	Name string `json:"name"`
}

var (
	loadRequired                 = true
	pathToFontsResourceDirectory string
	fontDetails                  []FontInformation
	lockObject                   sync.Mutex
)

func init() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Error during OS system font path detection: %v\n", r)
			pathToFontsResourceDirectory = ""
		}
	}()

	// -----------------
	// Windows
	// -----------------
	if runtime.GOOS == "windows" {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Error detecting Windows font directory: %v\n", r)
			}
		}()

		if windir := os.Getenv("WINDIR"); windir != "" {
			pathToFontsResourceDirectory = filepath.Join(windir, "Fonts")
			return
		}
	}

	// -----------------
	// macOS
	// -----------------
	if runtime.GOOS == "darwin" {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Error detecting macOS font directories: %v\n", r)
			}
		}()

		macPaths := []string{
			"/System/Library/Fonts", "/Library/Fonts", filepath.Join(os.Getenv("HOME"), "Library", "Fonts"),
		}

		for _, p := range macPaths {
			if dirExists(p) {
				pathToFontsResourceDirectory = p
				return
			}
		}
	}

	// -----------------
	// Linux
	// -----------------
	if runtime.GOOS == "linux" {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Error detecting Linux font directories: %v\n", r)
			}
		}()

		linuxPaths := []string{
			"/usr/share/fonts", "/usr/local/share/fonts", filepath.Join(os.Getenv("HOME"), ".fonts"), filepath.Join(os.Getenv("HOME"), ".local", "share", "fonts"),
		}

		for _, p := range linuxPaths {
			if dirExists(p) {
				pathToFontsResourceDirectory = p
				return
			}
		}
	}

	// Fallback
	pathToFontsResourceDirectory = ""
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// Initializes a new instance of the `Font` class
func NewFont() *Font {
	var p Font
	p.Name = uuid.NewString()
	return &p
}

// NewFontResource initializes a new Font with given font file provided with path.
func NewFontResource(path string, resourceName string) *Font {
	if resourceName == "" {
		resourceName = filepath.Base(path)
	}
	p := Font{FontResource: *resource.NewFontWithResourcePath(path, resourceName)}
	p.Name = uuid.NewString()
	p.ResourceName = resourceName
	return &p
}

func NewFontWithResource(path string) *Font {
	return NewFontResource(path, "")
}

// NewFontFromStreamWithResourceName initializes a new Font with given font file provided as stream.
func NewFontFromStreamWithResourceName(stream io.Reader, resourceName string) *Font {
	if resourceName == "" {
		resourceName = uuid.New().String()
	}
	p := Font{FontResource: *resource.NewFontWithStreamResource(stream, resourceName)}
	p.Name = uuid.NewString()
	p.ResourceName = resourceName
	return &p
}

func NewFontFromStream(stream io.Reader) *Font {
	return NewFontFromStreamWithResourceName(stream, "")
}

func NewFontCloudResource(cloudResourceName string) *Font {
	var p Font
	p.ResourceName = cloudResourceName
	p.Name = uuid.NewString()
	return &p
}

// Gets the Times Italic core font with Latin 1 encoding.
func TimesItalic() Font {
	return Font{Name: "timesItalic"}
}

// Gets the Times Roman core font with Latin 1 encoding.
func TimesRoman() Font {
	return Font{Name: "timesRoman"}
}

// Gets the Times Bold core font with Latin 1 encoding.
func TimesBold() Font {
	return Font{Name: "timesBold"}
}

// Gets the Times Bold Italic core font with Latin 1 encoding.
func TimesBoldItalic() Font {
	return Font{Name: "timesBoldItalic"}
}

// Gets the Helvetica core font with Latin 1 encoding.
func Helvetica() Font {
	return Font{Name: "helvetica"}
}

// Gets the Helvetica Bold core font with Latin 1 encoding.
func HelveticaBold() Font {
	return Font{Name: "helveticaBold"}
}

// Gets the Helvetica Oblique core font with Latin 1 encoding.
func HelveticaOblique() Font {
	return Font{Name: "helveticaOblique"}
}

/** Gets the Helvetica Bold Oblique core font with Latin 1 encoding. */
func HelveticaBoldOblique() Font {
	return Font{Name: "helveticaBoldOblique"}
}

// Gets the Courier core font with Latin 1 encoding.
func Courier() Font {
	return Font{Name: "courier"}
}

// Gets the Courier Bold core font with Latin 1 encoding.
func CourierBold() Font {
	return Font{Name: "courierBold"}
}

// Gets the Courier Oblique core font with Latin 1 encoding.
func CourierOblique() Font {
	return Font{Name: "courierOblique"}
}

// Gets the Courier Bold Oblique core font with Latin 1 encoding.
func CourierBoldOblique() Font {
	return Font{Name: "courierBoldOblique"}
}

// Gets the Symbol core font.
func Symbol() Font {
	return Font{Name: "symbol"}
}

// Gets the Zapf Dingbats core font.
func ZapfDingbats() Font {
	return Font{Name: "zapfDingbats"}
}

func getGoogleFontText(name string, weight int, italic bool) string {
	if italic {
		return name + ":" + strconv.Itoa(weight) + "italic"
	} else {
		return name + ":" + strconv.Itoa(weight)
	}
}

/*
Initializes a new instance of the Font class using the Google font name.
  - @param fontName The name of the google font.
  - @return The font from the google.
*/
func Google(fontName string) *Font {
	font := NewFont()
	font.Name = fontName
	return font
}

/*
Initializes a new instance of the Font class using the google font name, weight and italic.
  - @param fontName The name of the google font.
  - @param weight The weight of the font.
  - @param italic The italic property of the font.
  - @return The font from the google.
*/
func GoogleFontWithWeight(fontName string, weight int, italic bool) *Font {
	font := NewFont()
	font.Name = getGoogleFontText(fontName, weight, italic)
	return font
}

/*
Initializes a new instance of the Font class using the google font name, bold and italic.
  - @param fontName The name of the google font.
  - @param bold If true font weight will be taken as 700 otherwise 400.
  - @param italic The italic property of the font.
  - @return The font from the google.
*/
func GoogleFontWithStyle(fontName string, bold bool, italic bool) *Font {
	font := NewFont()
	weight := 400
	if bold {
		weight = 700
	}
	font.Name = getGoogleFontText(fontName, weight, italic)
	return font
}

/*
Initializes a new instance of the Font class using the font from the global storage.
  - @param fontName The name of the font to get from the global storage..
  - @return The font from the global.
*/
func Global(fontName string) *Font {
	font := NewFont()
	font.Name = fontName
	return font
}

/*
Initializes a new instance of the <code>Font</code> class using the system font name and resource name.
  - @param fontName The name of the font in the system.
  - @param resourceName The resource name for the font.
  - @return The font from system.
*/
func FontFromSystemWithResourceName(fontName string, resourceName string) *Font {
	if fontName == "" {
		return nil
	}

	fontName = strings.ReplaceAll(fontName, "-", "")
	fontName = strings.ReplaceAll(fontName, " ", "")

	if loadRequired {
		LoadFonts()
	}

	for _, fontDetail := range fontDetails {
		if strings.EqualFold(fontDetail.fontName, fontName) {
			fontRes := NewFontResource(fontDetail.filePath, resourceName)
			return fontRes
		}
	}
	return nil
}

/*
Initializes a new instance of the <code>Font</code> class using the system font name and resource name.
  - @param fontName The name of the font in the system.
  - @return The font from system.
*/
func FontFromSystem(fontName string) *Font {
	return FontFromSystemWithResourceName(fontName, "")
}

func LoadFonts() {
	lockObject.Lock()
	if !loadRequired {
		return
	}

	if pathToFontsResourceDirectory != "" {
		files, err := os.ReadDir(pathToFontsResourceDirectory)
		if err != nil {
			log.Printf("Error reading font directory: %v\n", err)
			return
		}

		for _, file := range files {
			if file.IsDir() || !(strings.HasSuffix(strings.ToLower(file.Name()), ".ttf") || strings.HasSuffix(strings.ToLower(file.Name()), ".otf")) {
				continue
			}

			fullPath := filepath.Join(pathToFontsResourceDirectory, file.Name())
			f, err := os.Open(fullPath)
			if err != nil {
				log.Printf("Failed to open font file: %s, error: %v\n", fullPath, err)
				continue
			}

			name := ReadFontNameTable(f)
			f.Close()

			if name != nil && name.fontName != "" {
				fontDetails = append(fontDetails, FontInformation{
					fontName: name.fontName,
					filePath: fullPath,
				})
			}
		}
		loadRequired = false
	}
	defer lockObject.Unlock()
}

func ReadFontNameTable(reader io.ReadSeeker) *FullNameTable {
	var nameTable *FullNameTable

	defer func() {
		_ = recover() // ignore errors like C# catch {}
	}()

	// Skip scaler type (first 4 bytes)
	reader.Seek(4, io.SeekStart)

	// Read number of tables (2 bytes)
	buf := make([]byte, 2)
	reader.Read(buf)
	intTableCount := int(buf[0])<<8 | int(buf[1])

	if intTableCount > 0 {
		// Each table directory entry is 16 bytes
		bytTableDirectory := make([]byte, intTableCount*16)

		reader.Seek(12, io.SeekStart)
		io.ReadFull(reader, bytTableDirectory)

		for i := 0; i < len(bytTableDirectory); i += 16 {
			switch binary.BigEndian.Uint32(bytTableDirectory[i:]) {
			case 1851878757: //"0x6E616D65" or "name"
				nameTable, _ = newFullNameTable(reader, bytTableDirectory, i)
			}
		}
	}

	return nameTable
}
