// package ascii_art

// import (
// 	"strings"
// )

// func Generate(text string, bannerFile string) (string, error) {
// 	banner, err := LoadBanner(bannerFile)
// 	if err != nil {
// 		return "", err
// 	}

// 	_, err := ValidateInput(text)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Split text into lines
// 	lines := SplitText(text)

// 	var finalResult strings.Builder

// 	// Render each line separately
// 	for idx, line := range lines {
// 		if line == "" {
// 			finalResult.WriteString("\n")
// 			continue
// 		}

// 		rendered, err := RenderText(line, banner)
// 		if err != nil {
// 			return "", err
// 		}

// 		// join 8 ASCII rows into output
// 		for _, row := range rendered {
// 			finalResult.WriteString(row)
// 			finalResult.WriteString("\n")
// 		}

// 		// avoid extra newline at end except last line
// 		if idx != len(lines)-1 {
// 			finalResult.WriteString("\n")
// 		}
// 	}

// 	return finalResult.String(), nil
// }




//--------------------------------------------------------------------------------------------
package ascii_art

import (
	"fmt"
	"os"
	"strings"
)

func GenerateAscii(text, banner string) (string, error) {
	file, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("error reading banner", err)
		os.Exit(1)
	}

	content := strings.ReplaceAll(string(file), "\r\n", "\n")
	bannerLines := strings.Split(content, "\n")
	text = strings.ReplaceAll(text, "\r\n", "\n")
	words := strings.Split(text, "\n")
	build := strings.Builder{}

	for _, word := range words {
		if word == "" {
			build.WriteString("\n")
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range word {
				if ch < 32 || ch > 126 {
					fmt.Println("invalid character")
					return "", nil
				}

				index := (int(ch)-32)*9 + 1 + row

				build.WriteString(bannerLines[index])
			}
			build.WriteString("\n")
		}
	}
	return build.String(), nil
}
