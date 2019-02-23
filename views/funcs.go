package views

import (
	"fmt"
	"strings"
	"time"

	"git.urantiatech.com/cloudcms/cloudcms/item"
)

// Status prints status header field
func Status(s string) string {
	if s == "" {
		return "Published"
	}
	return s
}

// TimeToString converts timestamp to String
func TimeToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.String()
}

// UnixTimeToString converts unix timestamp to String
func UnixTimeToString(t float64) string {
	return TimeToString(time.Unix(int64(t), 0))
}

// TimeToDate converts timestamp to Date
func TimeToDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("Jan 2, 2006")
}

// UnixTimeToDate converts unix timestamp to Date
func UnixTimeToDate(t float64) string {
	return TimeToDate(time.Unix(int64(t), 0))
}

// AppendField appends another field to fields
func AppendField(fields string, field item.Field) string {
	if fields == "" {
		return strings.ToLower(field.Name)
	}
	return fields + "," + strings.ToLower(field.Name)
}

// ContentTextValue gets value of field from the content map
func ContentTextValue(content map[string]interface{}, field string) string {
	if c, ok := content[field]; ok {
		return fmt.Sprint(c)
	}
	return ""
}

// ContentDateValue gets value of date field from the content map
func ContentDateValue(content map[string]interface{}, field string) string {
	if c, ok := content[field]; ok {
		t, err := time.Parse(time.RFC3339, c.(string))
		if err != nil {
			return ""
		}
		return t.Format("2006-01-02")
	}
	return ""
}

// ContentTagsValue gets value of date field from the content map
func ContentTagsValue(content map[string]interface{}, field string) string {
	if c, ok := content[field]; ok {

		switch c.(type) {
		case []interface{}:
			var tags []string
			for _, t := range c.([]interface{}) {
				tags = append(tags, t.(string))
			}
			return strings.Join(tags, ", ")

		default:
			return c.(string)
		}
	}
	return ""
}

// ContentFile gets value of file field from the content map
func ContentFile(content map[string]interface{}, field string) *item.File {
	var file *item.File
	if name, ok := content[field+".name"]; ok {
		file = &item.File{
			Name: name.(string),
			Size: int64(content[field+".size"].(float64)),
			URI:  content[field+".uri"].(string),
		}
	}
	return file
}

// CurrentDate in YYYY-MM-DD
func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}
