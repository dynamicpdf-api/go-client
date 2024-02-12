package element

import (
	"encoding/json"

	"github.com/dynamicpdf-api/go-client/font"
	"github.com/dynamicpdf-api/go-client/resource"
)

// Base class from which all page elements are derived.
type ElementCollector interface {
	InputValue() string
	ElementType() ElementType
	TextFont() font.Font
	Resources() resource.Resource
}
type Element struct {
	ElementCollector `json:"-"`
	// Gets and sets placement of the page element on the page.
	Placement elementPlacement `json:"placement,omitempty"`
	// Gets or sets the X coordinate of the page element.
	XOffset float32 `json:"xOffset,omitempty"`
	// Gets or sets the Y coordinate of the page element.
	YOffset float32 `json:"yOffset,omitempty"`
	// Gets or sets the boolean value specifying whether the element should be added to even pages or not.
	EvenPages bool `json:"evenPages,omitempty"`
	// Gets or sets the boolean value specifying whether the element should be added to odd pages or not.
	OddPages bool `json:"oddPages,omitempty"`

	grayLevel                   int
	width                       float32
	height                      float32
	x2Offset                    float32
	y2Offset                    float32
	borderWidth                 float32
	cornerRadius                float32
	processTilde                bool
	symbolSize                  int
	aztecErrorCorrection        int
	readerInitialization        bool
	uccEan128                   bool
	includeCheckDigit           bool
	scaleX                      float32
	scaleY                      float32
	maxHeight                   float32
	maxWidth                    float32
	appendCheckDigit            int
	fontSize                    float32
	columns                     int
	yDimension                  float32
	compactPdf417               bool
	errorCorrection             ErrorCorrection
	compaction                  compaction
	version                     int
	fnc1                        QrCodeFnc1
	stackedGs1DataBarType       stackedGs1DataBarType
	rowHeight                   float32
	expandedStackedSegmentCount int
	textValue                   string
	value                       string
	typeOfElement               ElementType
	lineStyleName               string
	resourceName                string
	font                        font.Font
	fillColorName               string
	borderColorName             string
	borderStyleName             string
	dataMatrixSymbolSize        DataMatrixSymbolSize
	dataMatrixEncodingType      DataMatrixEncodingType
	dataMatrixFunctionCharacter DataMatrixFunctionCharacter
	valueType                   ValueType
	colorName                   string
	gs1DataBarType              Gs1DataBarType
	fontName                    string
	xDimension                  float32
	showText                    bool
	textColorName               string
	resource                    resource.Resource
}

func newElement(value string, placement elementPlacement, xOffset float32, yOffset float32) *Element {
	var p Element
	p.textValue = value
	p.Placement = placement
	p.XOffset = xOffset
	p.YOffset = yOffset
	return &p
}
func (e *Element) Resources() resource.Resource {
	return e.resource
}

func (e *Element) TextFont() font.Font {
	return e.font
}

func (p Element) MarshalJSON() ([]byte, error) {
	type Alias Element
	return json.Marshal(&struct {
		Width                       float32                     `json:"width,omitempty"`
		Height                      float32                     `json:"height,omitempty"`
		X2Offset                    float32                     `json:"x2Offset,omitempty"`
		Y2Offset                    float32                     `json:"y2Offset,omitempty"`
		BorderWidth                 float32                     `json:"borderWidth,omitempty"`
		CornerRadius                float32                     `json:"cornerRadius,omitempty"`
		ProcessTilde                bool                        `json:"processTilde,omitempty"`
		SymbolSize                  int                         `json:"symbolSize,omitempty"`
		AztecErrorCorrection        int                         `json:"aztecErrorCorrection,omitempty"`
		ReaderInitialization        bool                        `json:"readerInitialization,omitempty"`
		UccEan128                   bool                        `json:"uccEan128,omitempty"`
		IncludeCheckDigit           bool                        `json:"includeCheckDigit,omitempty"`
		ScaleX                      float32                     `json:"scaleX,omitempty"`
		ScaleY                      float32                     `json:"scaleY,omitempty"`
		MaxHeight                   float32                     `json:"MaxHeight,omitempty"`
		MaxWidth                    float32                     `json:"MaxWidth,omitempty"`
		AppendCheckDigit            int                         `json:"appendCheckDigit,omitempty"`
		FontSize                    float32                     `json:"fontSize,omitempty"`
		Columns                     int                         `json:"columns,omitempty"`
		YDimension                  float32                     `json:"yDimension,omitempty"`
		CompactPdf417               bool                        `json:"compactPdf417,omitempty"`
		ErrorCorrection             ErrorCorrection             `json:"errorCorrection,omitempty"`
		Compaction                  compaction                  `json:"compaction,omitempty"`
		Version                     int                         `json:"version,omitempty"`
		Fnc1                        QrCodeFnc1                  `json:"fnc1,omitempty"`
		StackedGs1DataBarType       stackedGs1DataBarType       `json:"stackedGs1DataBarType,omitempty"`
		RowHeight                   float32                     `json:"rowHeight,omitempty"`
		ExpandedStackedSegmentCount int                         `json:"expandedStackedSegmentCount,omitempty"`
		TextValue                   string                      `json:"text,omitempty"`
		Type                        ElementType                 `json:"type,omitempty"`
		LineStyleName               string                      `json:"lineStyle,omitempty"`
		BorderStyleName             string                      `json:"borderStyle,omitempty"`
		BorderColorName             string                      `json:"borderColor,omitempty"`
		FillColorName               string                      `json:"fillColor,omitempty"`
		Value                       string                      `json:"value,omitempty"`
		DataMatrixSymbolSize        DataMatrixSymbolSize        `json:"dataMatrixSymbolSize,omitempty"`
		DataMatrixEncodingType      DataMatrixEncodingType      `json:"dataMatrixEncodingType,omitempty"`
		DataMatrixFunctionCharacter DataMatrixFunctionCharacter `json:"dataMatrixFunctionCharacter,omitempty"`
		ValueType                   ValueType                   `json:"valueType,omitempty"`
		ColorName                   string                      `json:"color,omitempty"`
		Gs1DataBarType              Gs1DataBarType              `json:"gs1DataBarType,omitempty"`
		FontName                    string                      `json:"font,omitempty"`
		TextColor                   string                      `json:"textColor,omitempty"`
		ResourceName                string                      `json:"resourceName,omitempty"`
		Alias
	}{
		Width:                       p.width,
		Height:                      p.height,
		X2Offset:                    p.x2Offset,
		Y2Offset:                    p.y2Offset,
		BorderWidth:                 p.borderWidth,
		CornerRadius:                p.cornerRadius,
		ProcessTilde:                p.processTilde,
		SymbolSize:                  p.symbolSize,
		AztecErrorCorrection:        p.aztecErrorCorrection,
		ReaderInitialization:        p.readerInitialization,
		UccEan128:                   p.uccEan128,
		IncludeCheckDigit:           p.includeCheckDigit,
		ScaleX:                      p.scaleX,
		ScaleY:                      p.scaleY,
		MaxHeight:                   p.maxHeight,
		MaxWidth:                    p.maxWidth,
		AppendCheckDigit:            p.appendCheckDigit,
		FontSize:                    p.fontSize,
		Columns:                     p.columns,
		YDimension:                  p.yDimension,
		CompactPdf417:               p.compactPdf417,
		ErrorCorrection:             p.errorCorrection,
		Compaction:                  p.compaction,
		Version:                     p.version,
		Fnc1:                        p.fnc1,
		StackedGs1DataBarType:       p.stackedGs1DataBarType,
		RowHeight:                   p.rowHeight,
		ExpandedStackedSegmentCount: p.expandedStackedSegmentCount,
		TextValue:                   p.textValue,
		Type:                        p.typeOfElement,
		LineStyleName:               p.lineStyleName,
		BorderStyleName:             p.borderStyleName,
		BorderColorName:             p.borderColorName,
		FillColorName:               p.fillColorName,
		Value:                       p.value,
		DataMatrixSymbolSize:        p.dataMatrixSymbolSize,
		DataMatrixEncodingType:      p.dataMatrixEncodingType,
		DataMatrixFunctionCharacter: p.dataMatrixFunctionCharacter,
		ValueType:                   p.valueType,
		ColorName:                   p.colorName,
		Gs1DataBarType:              p.gs1DataBarType,
		FontName:                    p.fontName,
		TextColor:                   p.textColorName,
		ResourceName:                p.resourceName,
		Alias:                       (Alias)(p),
	})
}
