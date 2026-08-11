package pdf

import (
	"bytes"
	"io"
	"strings"

	"github.com/ledongthuc/pdf"
)

func ExtractText(file io.Reader) (string, error) {

	pdfBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	reader, err := pdf.NewReader(
		bytes.NewReader(pdfBytes),
		int64(len(pdfBytes)),
	)
	if err != nil {
		return "", err
	}

	var text strings.Builder

	for pageNumber := 1; pageNumber <= reader.NumPage(); pageNumber++ {

		page := reader.Page(pageNumber)

		if page.V.IsNull() {
			continue
		}

		pageText, err := page.GetPlainText(nil)
		if err != nil {
			return "", err
		}

		text.WriteString(pageText)
		text.WriteString("\n")
	}

	return text.String(), nil
}