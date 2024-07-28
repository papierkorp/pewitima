package filemanager

import (
    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/parser"
    "github.com/microcosm-cc/bluemonday"
    "io/ioutil"
)

func ParseMarkdownToHtml(path string) (string, error) {
    // cwd, err := os.Getwd()
    // if err != nil {
    //     return "", err
    // }

    // dataDir := filepath.Join(cwd, "data")
    // path := filepath.Join(dataDir, filename)

    content, err := ioutil.ReadFile(path)
    if err != nil {
        return "", err
    }

    extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
    parser := parser.NewWithExtensions(extensions)
    maybeUnsafeHTML := markdown.ToHTML(content, parser, nil)
    html := bluemonday.UGCPolicy().SanitizeBytes(maybeUnsafeHTML)

    return string(html), nil
}
