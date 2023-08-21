package themes

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type Collection interface {
	Collect(*http.Client) (string, error)
	URL(...string) Collection
	Inline(...string) Collection
	TempFilePrefix(string)
}

const (
	defaultTempFilePrefix = "cv-go-style"
)

var (
	defaultTheme = NewCollection().
			URL(sazakBaseThemeCSS).
			Inline(sazakStyleCustomization)

	allThemes = map[string]Collection{
		"sazak": defaultTheme,
		"sazak-light": NewCollection().
			URL(sazakLightBaseThemeCSS).
			Inline(sazakLightBaseStyleCustomization).
			Inline(sazakLightStyleCustomization),
	}
)

func GetThemeCollection(name string) Collection {
	theme, ok := allThemes[name]
	if !ok {
		return nil
	}
	return theme
}

func GetDefaultTheme() Collection {
	return defaultTheme
}

type collection struct {
	Items            []*cssItem
	OutputFilePrefix string
}

type cssItem struct {
	Type    CSSItemType
	Content string
}

type CSSItemType int

const (
	CSSItemTypeFileURL CSSItemType = iota
	CSSItemTypeInlineContent
)

func NewCollection() Collection {
	return &collection{
		Items:            make([]*cssItem, 0),
		OutputFilePrefix: defaultTempFilePrefix,
	}
}

func (c *collection) URL(urls ...string) Collection {
	for _, u := range urls {
		c.Items = append(c.Items, &cssItem{
			Type:    CSSItemTypeFileURL,
			Content: u,
		})
	}
	return c
}

func (c *collection) Inline(contents ...string) Collection {
	for _, content := range contents {
		c.Items = append(c.Items, &cssItem{
			Type:    CSSItemTypeInlineContent,
			Content: content,
		})
	}
	return c
}

func (c *collection) TempFilePrefix(prefix string) {
	c.OutputFilePrefix = prefix
}

func (c *collection) Collect(cl *http.Client) (string, error) {
	var sb strings.Builder

	for index, c := range c.Items {
		switch c.Type {
		case CSSItemTypeInlineContent:
			sb.WriteString(c.Content)
			sb.WriteRune('\n')
		case CSSItemTypeFileURL:
			req, err := http.NewRequest(http.MethodGet, c.Content, nil)
			if err != nil {
				return "", fmt.Errorf("failed to create HTTP request for CSS import #%d: %w", index, err)
			}
			res, err := cl.Do(req)
			if err != nil {
				return "", fmt.Errorf("failed to do HTTP request for CSS import #%d: %w", index, err)
			}
			defer res.Body.Close()

			data, err := io.ReadAll(res.Body)
			if err != nil {
				return "", fmt.Errorf("failed to read HTTP response body for CSS import #%d: %w", index, err)
			}

			sb.Write(data)
			sb.WriteRune('\n')
		}
	}

	filename := fmt.Sprintf("%s-%s.css", c.OutputFilePrefix, uuid.New().String()[:8])
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := f.WriteString(sb.String()); err != nil {
		return "", err
	}
	slog.Info("Successfully merged style files")

	return filename, nil
}
