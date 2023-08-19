package builder

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

const (
	DefaultConfigPath = "cv.yaml"
	DefaultOutputPath = "cv.pdf"

	httpTimeout       = 10 * time.Second
	tempStyleFileName = "cv-go-style"
	tempHTMLFileName  = "cv-go-out"
)

type Builder interface {
	Build() (string, error)

	setOutputPath(string)
	setConfigPath(string)
	setNoPDF()
}

type Option func(Builder)

func WithOutputPath(path string) Option {
	return func(b Builder) {
		b.setOutputPath(path)
	}
}

func WithConfigPath(path string) Option {
	return func(b Builder) {
		b.setConfigPath(path)
	}
}

func WithNoPDF() Option {
	return func(b Builder) {
		b.setNoPDF()
	}
}

type builder struct {
	conf     *Config
	outPath  string
	confPath string
	noPDF    bool
	httpCl   *http.Client
}

func New(opts ...Option) (Builder, error) {
	b := &builder{
		outPath:  DefaultOutputPath,
		confPath: DefaultConfigPath,
		httpCl: &http.Client{
			Timeout: httpTimeout,
		},
	}
	for _, o := range opts {
		o(b)
	}
	conf, err := ParseConfig(b.confPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	b.conf = conf
	return b, nil
}

func (b *builder) Build() (string, error) {
	cssFile, err := b.getPDFStyle()
	if err != nil {
		return "", fmt.Errorf("failed to fetch and merge CSS styles: %w", err)
	}

	cors := func(fs http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", "*")

			fs.ServeHTTP(w, r)
		}
	}

	ts := httptest.NewServer(cors(http.FileServer(http.Dir("./"))))
	defer ts.Close()

	tmpl, err := template.New("cv-template").Parse(cvTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse CV template: %w", err)
	}

	styleFile := fmt.Sprintf("%s/%s", ts.URL, cssFile)
	if b.noPDF {
		styleFile = "/" + cssFile
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, CVTemplateData{
		Conf:      b.conf,
		StyleFile: styleFile,
	}); err != nil {
		return "", fmt.Errorf("failed to execute CV template: %w", err)
	}

	// htmlFile := fmt.Sprintf("%s-%s.html", tempHTMLFileName, uuid.New().String()[:8])
	htmlFile := tempHTMLFileName + ".html"
	if err := os.WriteFile(htmlFile, buf.Bytes(), 0666); err != nil {
		return "", fmt.Errorf("failed to create temporary CV HTML: %w", err)
	}

	if b.noPDF {
		return htmlFile, nil
	}

	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %w", err)
	}
	htmlPath := fmt.Sprintf("file://%s/%s", cwd, htmlFile)

	if err := saveURLToPDF(htmlPath, b.outPath, 0.8); err != nil {
		return "", fmt.Errorf("failed to save CV as PDF: %w", err)
	}

	return b.outPath, nil
}

func (b *builder) setOutputPath(path string) {
	b.outPath = path
}

func (b *builder) setConfigPath(path string) {
	b.confPath = path
}

func (b *builder) setNoPDF() {
	b.noPDF = true
}

func (b *builder) getPDFStyle() (string, error) {
	if b.conf.Meta == nil || len(b.conf.Meta.CSS) == 0 {
		return "", nil
	}

	var sb strings.Builder

	for index, c := range b.conf.Meta.CSS {
		if c.File != nil {
			b, err := os.ReadFile(*c.File)
			if err != nil {
				return "", fmt.Errorf("failed to read file content of CSS import #%d: %w", index, err)
			}
			sb.Write(b)
			sb.WriteRune('\n')
		} else if c.URL != nil {
			req, err := http.NewRequest(http.MethodGet, *c.URL, nil)
			if err != nil {
				return "", fmt.Errorf("failed to create HTTP request for CSS import #%d: %w", index, err)
			}
			res, err := b.httpCl.Do(req)
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

	filename := fmt.Sprintf("%s-%s.css", tempStyleFileName, uuid.New().String()[:8])
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

func saveURLToPDF(url string, outPath string, scale float64) error {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf), chromedp.WithLogf(log.Printf))
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(url, &buf, scale)); err != nil {
		return err
	}

	if err := os.WriteFile(outPath, buf, 0644); err != nil {
		return err
	}

	return nil
}

func printToPDF(url string, res *[]byte, scale float64) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(10 * time.Second), // Wait for ionicons to load fully
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				WithScale(scale).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
