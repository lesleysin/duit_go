package duit_text_properties

import "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_utility_styles"

type TextStyle[T duit_utility_styles.Color] struct {
	Color         T          `json:"color,omitempty"`
	FontFamily    string     `json:"fontFamily,omitempty"`
	FontWeight    FontWeight `json:"fontWeight,omitempty"`
	FontSize      float32    `json:"fontSize,omitempty"`
	LetterSpacing float32    `json:"letterSpacing,omitempty"`
	WordSpacing   float32    `json:"wordSpacing,omitempty"`
	Height        float32    `json:"height,omitempty"`
}
