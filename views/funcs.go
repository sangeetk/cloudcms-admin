package views

import (
	"fmt"
	"strings"
	"time"

	"git.urantiatech.com/cloudcms/cloudcms/item"
)

// DriveX mapping
var DriveX string

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

// ContentDateTimeValue gets value of date field from the content map
func ContentDateTimeValue(content map[string]interface{}, field string) string {
	if c, ok := content[field]; ok {
		t, err := time.Parse(time.RFC3339, c.(string))
		if err != nil {
			return ""
		}
		return t.Format("2006-01-02 15:04:05")
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
			if c == nil {
				return ""
			}
			return c.(string)
		}
	}
	return ""
}

// ContentListValue
func ContentListValue(content map[string]interface{}, field string) string {
	if c, ok := content[field]; ok {

		switch c.(type) {
		case []interface{}:
			var list []string
			for _, t := range c.([]interface{}) {
				list = append(list, t.(string))
			}
			return strings.Join(list, ",\n")

		default:
			if c == nil {
				return ""
			}
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
			Type: content[field+".type"].(string),
			Size: int64(content[field+".size"].(float64)),
			URI:  content[field+".uri"].(string),
		}
	} else if fmap, ok := content[field]; ok {
		filemap := fmap.(map[string]interface{})
		file = &item.File{
			Name: filemap["name"].(string),
			Type: filemap["type"].(string),
			Size: int64(filemap["size"].(float64)),
			URI:  filemap["uri"].(string),
		}
	}
	return file
}

// CurrentDate in YYYY-MM-DD
func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}

// CurrentDateTime in YYYY-MM-DD HH:MI:SS
func CurrentDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// CloudDrive mapping
func CloudDrive(s string) string {
	return strings.Replace(s, "drive", DriveX, 1)
}

// Pager contains the info about pagination buttons
type Pager struct {
	PageNum int
	Title   string
	URI     string
	Active  bool
}

func PagerFn(maxPagers, currentPage, totalItems, itemsPerPage int, query string) []Pager {
	var pagers = []Pager{}

	totalPages := totalItems / itemsPerPage
	if totalItems%itemsPerPage != 0 {
		totalPages++
	}

	var uri string
	if query == "" {
		uri = fmt.Sprintf("?")
	} else {
		query = strings.Replace(query, " ", "+", -1)
		uri = fmt.Sprintf("?q=%s&", query)
	}

	if currentPage > 1 {
		pagers = append(pagers, Pager{0, "Prev", fmt.Sprintf("%sp=%d", uri, currentPage-1), false})
	}

	start := 0
	if currentPage > maxPagers/2 {
		start = currentPage - maxPagers/2
	}

	if totalPages-start < maxPagers && totalPages > maxPagers {
		start = totalPages - maxPagers
	}

	for i := start; i < maxPagers+start && i < totalPages; i++ {
		var p Pager
		p.PageNum = i + 1
		p.Title = fmt.Sprint(p.PageNum)
		p.URI = fmt.Sprintf("%sp=%d", uri, p.PageNum)
		p.Active = p.PageNum == currentPage
		pagers = append(pagers, p)
	}

	if start+maxPagers < totalPages {
		pagers = append(pagers, Pager{0, "...", "#", false})
	}

	if currentPage < totalPages {
		pagers = append(pagers, Pager{0, "Next", fmt.Sprintf("%sp=%d", uri, currentPage+1), false})
	}

	return pagers
}
