package entities

import (
	"time"

	"github.com/google/uuid"
)

// FormField represents a question/field in a form
type FormField struct {
	ID          uuid.UUID         `db:"id" json:"id"`
	FormID      uuid.UUID         `db:"form_id" json:"form_id"`
	Label       map[string]string `db:"label" json:"label"`                       // {"en": "Name", "ar": "الاسم"}
	Placeholder map[string]string `db:"placeholder" json:"placeholder,omitempty"` // Optional
	HelpText    map[string]string `db:"help_text" json:"help_text,omitempty"`     // Optional
	Required    bool              `db:"required" json:"required"`
	Options     map[string]any    `db:"options" json:"options,omitempty"` // JSONB for select/radio/checkbox
	FieldOrder  int               `db:"field_order" json:"field_order"`   // Position/order in the form
	Type        int16             `db:"type" json:"type"`                 // text, textarea, select, radio, etc.
	CreatedAt   time.Time         `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time         `db:"updated_at" json:"updated_at"`
}

// TableName returns the DB table name
func (FormField) TableName() string {
	return "form_fields"
}

// HasOptions checks if the field has selectable options
func (ff *FormField) HasOptions() bool {
	return len(ff.Options) > 0
}

// IsRequired checks if the field is mandatory
func (ff *FormField) IsRequired() bool {
	return ff.Required
}

// GetLabel returns the label in the requested language, defaults to English
func (ff *FormField) GetLabel(lang string) string {
	if val, ok := ff.Label[lang]; ok && val != "" {
		return val
	}
	if val, ok := ff.Label["en"]; ok {
		return val
	}
	return ""
}

// GetPlaceholder returns the placeholder in the requested language, defaults to English
func (ff *FormField) GetPlaceholder(lang string) string {
	if ff.Placeholder == nil {
		return ""
	}
	if val, ok := ff.Placeholder[lang]; ok && val != "" {
		return val
	}
	if val, ok := ff.Placeholder["en"]; ok {
		return val
	}
	return ""
}

// GetHelpText returns the help text in the requested language, defaults to English
func (ff *FormField) GetHelpText(lang string) string {
	if ff.HelpText == nil {
		return ""
	}
	if val, ok := ff.HelpText[lang]; ok && val != "" {
		return val
	}
	if val, ok := ff.HelpText["en"]; ok {
		return val
	}
	return ""
}
