package bowl

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type Cases struct {
	input  []string
	output []string
}

func visit(cases *Cases) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() {
			return nil
		}
		if strings.Contains(path, "input") {
			cases.input = append(cases.input, path)
			return nil
		}
		cases.output = append(cases.output, path)
		return nil
	}
}

func getFiles(directory string) Cases {
	var cases Cases

	if err := filepath.Walk(directory, visit(&cases)); err != nil {
		panic(err)
	}
	return cases
}

func TestGetScoreWithUseCase1(t *testing.T) {
	directory := "./use-cases/use-cases-1"
	files := getFiles(directory)

	for i, file := range files.input {
		var frames []int
		var expectedScore Score

		content, err := ioutil.ReadFile(file)
		if err != nil {
			t.Fatalf("unexpected err: %s.", err.Error())
		}

		if err := json.Unmarshal(content, &frames); err != nil {
			t.Fatalf("Unexpected error: %s.", err.Error())
		}

		outputContent, err := ioutil.ReadFile(files.output[i])
		if err != nil {
			t.Fatalf("unexpected err: %s.", err.Error())
		}

		if err := json.Unmarshal(outputContent, &expectedScore); err != nil {
			t.Fatalf("Unexpected error: %s.", err.Error())
		}

		if got := GetScore(frames); got.Total != expectedScore.Total {
			t.Fatalf("Unexpected score for %s. Expected: %d - found: %d.", file, expectedScore.Total, got.Total)
		}
	}
}

func TestGetScoreWithUseCase2(t *testing.T) {
	directory := "./use-cases/use-cases-2"
	files := getFiles(directory)

	for i, file := range files.input {
		var frames []int
		var expectedScore Score

		content, err := ioutil.ReadFile(file)
		if err != nil {
			t.Fatalf("unexpected err: %s.", err.Error())
		}

		if err := json.Unmarshal(content, &frames); err != nil {
			t.Fatalf("Unexpected error: %s.", err.Error())
		}

		outputContent, err := ioutil.ReadFile(files.output[i])
		if err != nil {
			t.Fatalf("unexpected err: %s.", err.Error())
		}

		if err := json.Unmarshal(outputContent, &expectedScore); err != nil {
			t.Fatalf("Unexpected error: %s.", err.Error())
		}

		got := GetScore(frames)
		t.Logf("Score: %+v", got)
		if got.Total != expectedScore.Total {
			t.Fatalf("Unexpected score for %s. Expected: %d - found: %d.", file, expectedScore.Total, got.Total)
		}
		if got.Valid != expectedScore.Valid {
			t.Fatalf("Unexpected validity for %s. Expected: %t - found: %t.", file, expectedScore.Valid, got.Valid)
		}
		if got.ErrorMessage != expectedScore.ErrorMessage {
			t.Fatalf("Unexpected error message for %s. Expected: %+v - found: %+v.", file, expectedScore.ErrorMessage, got.ErrorMessage)
		}
	}
}
