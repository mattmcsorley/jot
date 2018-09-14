package internal

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

type Journal struct {
	docService DocumentService
}

func NewJournal(documentService DocumentService) *Journal {
	return &Journal{documentService}
}

func (j *Journal) ViewDocument(daysAgo int) {
	doc, _ := j.docService.LoadDocument(getFileName(0))
	buffer := bytes.Buffer{}
	buffer.Write(doc.Contents)
	cmd := exec.Command("less")
	cmd.Stdin = &buffer
	cmd.Stdout = os.Stdout
	cmd.Run()

}

func (j *Journal) SaveContent(c string, templatePath string) {
	doc, _ := j.docService.LoadDocument(getFileName(0))

	if len(doc.Contents) < 1 {
		data := struct {
			Title string
		}{
			getTitle(),
		}
		tmpl := template.Must(template.ParseFiles(templatePath))
		tmpl.Execute(doc, data)

		j.docService.SaveDocument(doc)
	}

	doc, _ = j.docService.LoadDocument(getFileName(0))

	if len(c) < 1 {
		contents, _ := getEditorContent(doc.Contents)
		doc.Contents = contents
	} else {
		doc.Contents = append(doc.Contents, c...)
		doc.Contents = append(doc.Contents, '\n')
	}
	j.docService.SaveDocument(doc)

	j.ViewDocument(0)
}

func getFileName(daysAgo int) string {
	now := time.Now()
	now.AddDate(0, 0, daysAgo*-1)
	return now.Format("20060102") + ".md"
}

func getTitle() string {
	now := time.Now()
	return now.Format("Mon Jan _2 2006")
}

func getEditorContent(contents []byte) ([]byte, error) {
	tmpfile, _ := ioutil.TempFile("", "journal")
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write(contents); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("vim", tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
	nContents, err := ioutil.ReadFile(tmpfile.Name())
	return nContents, err
}
