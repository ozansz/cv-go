package builder

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type ContactType string

const (
	ContactTypeEmail    ContactType = "email"
	ContactTypeGithub   ContactType = "github"
	ContactTypeLinkedIn ContactType = "linkedin"
	ContactTypeTwitter  ContactType = "twitter"
	ContactTypeWebsite  ContactType = "website"
	ContactTypePhone    ContactType = "phone"
)

var (
	// Formats are used to format the contact URL if there were no given.
	// We expect the contact text to be the username or website without proto
	//   if no URL was given to use the formatters correctly.
	contactURLFormats = map[ContactType]string{
		ContactTypeEmail:    "mailto:%s",
		ContactTypeGithub:   "https://github.com/%s",
		ContactTypeLinkedIn: "https://linkedin.com/in/%s",
		ContactTypeTwitter:  "https://twitter.com/%s",
		ContactTypeWebsite:  "https://%s",
		ContactTypePhone:    "tel:%s",
	}

	contactIcons = map[ContactType]string{
		ContactTypeEmail:    "mail-outline",
		ContactTypeGithub:   "logo-github",
		ContactTypeLinkedIn: "logo-linkedin",
		ContactTypeTwitter:  "logo-twitter",
		ContactTypeWebsite:  "link-outline",
		ContactTypePhone:    "call-outline",
	}
)

type Config struct {
	Header      CVHeader        `yaml:"header"`
	Experiences []*CVExperience `yaml:"jobs"`
	Education   []*CVEducation  `yaml:"education"`
	Projects    []*CVProject    `yaml:"projects,omitempty"`

	Meta *MetaConfig `yaml:"meta,omitempty"`
}

type CVHeader struct {
	Name        string     `yaml:"name"`
	Title       string     `yaml:"title"`
	Description string     `yaml:"tldr"`
	Contacts    []*Contact `yaml:"contacts"`
}

func (h *CVHeader) populateContactData() {
	// Populate the URLs and icons
	for _, c := range h.Contacts {
		if c.URL == nil {
			url := fmt.Sprintf(contactURLFormats[c.Type], c.Text)
			c.URL = &url
		}

		c.Icon = contactIcons[c.Type]
	}
}

type Contact struct {
	Type ContactType `yaml:"type"`
	Text string      `yaml:"text"`
	URL  *string     `yaml:"link,omitempty"`
	Icon string
}

type CVExperience struct {
	Title        string   `yaml:"title"`
	Company      string   `yaml:"company"`
	Description  string   `yaml:"tldr"`
	Location     *string  `yaml:"location,omitempty"`
	StartDate    string   `yaml:"start"`
	EndDate      *string  `yaml:"end,omitempty"`
	BulletPoints []string `yaml:"notable,omitempty"`

	DescriptionLines []string
}

type CVProject struct {
	Name        string         `yaml:"name"`
	Description string         `yaml:"tldr"`
	StartDate   *string        `yaml:"start,omitempty"`
	EndDate     *string        `yaml:"end,omitempty"`
	GithubRepo  *string        `yaml:"repo,omitempty"`
	Link        *CVProjectLink `yaml:"link,omitempty"`

	DescriptionLines []string
}

type CVEducation struct {
	Title        string   `yaml:"title"`
	School       string   `yaml:"school"`
	Location     string   `yaml:"location"`
	CGPA         *string  `yaml:"cgpa,omitempty"`
	StartDate    string   `yaml:"start"`
	EndDate      *string  `yaml:"end,omitempty"`
	Description  string   `yaml:"tldr"`
	BulletPoints []string `yaml:"notable,omitempty"`

	DescriptionLines []string
}

type CVProjectLink struct {
	URL   string `yaml:"url"`
	Title string `yaml:"title"`
}

type MetaConfig struct {
	CSS []*CSSConfig `yaml:"css,omitempty"`
}

type CSSConfig struct {
	URL  *string `yaml:"url,omitempty"`
	File *string `yaml:"file,omitempty"`
}

func (m *MetaConfig) validateCSSImports() error {
	for index, c := range m.CSS {
		if c.File == nil && c.URL == nil {
			return fmt.Errorf("CSS import #%d do not have URL or file import specified", index)
		}
	}
	return nil
}

func ParseConfig(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, err
	}

	c.Header.populateContactData()
	if c.Meta != nil {
		if err := c.Meta.validateCSSImports(); err != nil {
			return nil, err
		}
	}

	for _, p := range c.Projects {
		p.DescriptionLines = strings.Split(p.Description, "\n")
	}
	for _, p := range c.Experiences {
		p.DescriptionLines = strings.Split(p.Description, "\n")
	}

	return &c, nil
}
