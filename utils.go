package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/karrick/godirwalk"
)

// getAllFiles walks through the code directory and logs all the files
func getAllFiles() ([]string, error) {
	fileCount := 0

	allFiles := make([]string, 0)
	if err := godirwalk.Walk(codePath, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			// Following string operation is not most performant way
			// of doing this, but common enough to warrant a simple
			// example here:
			if strings.Contains(osPathname, ".git") {
				return godirwalk.SkipThis
			}
			if !de.IsDir() {
				allFiles = append(allFiles, osPathname)
				fileCount++
			}
			return nil
		},
		Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
	}); err != nil {
		return nil, err
	}
	fmt.Println("Total files: ", fileCount)

	return allFiles, nil
}

func createIssue(filePath string, lineNumber, column int) {
	vcsPath := path.Base(filePath)
	actualLineNumber := lineNumber + 1

	issue := Issue{
		Code:  "I001",
		Title: "Possible TODO comment found",
		Location: Location{
			Path: vcsPath,
			Position: Position{
				Begin: Coordinate{
					Line:   actualLineNumber,
					Column: column,
				},
				End: Coordinate{
					Line: actualLineNumber,
				},
			},
		},
	}
	issues = append(issues, issue)
}

func prepareResult() MacroResult {
	result := MacroResult{}
	result.Issues = issues
	result.IsPassed = false

	if len(issues) > 0 {
		result.IsPassed = true
	}

	return result
}

func writeMacroResult(result MacroResult) error {
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return err
	}

	f, err := os.Create(path.Join(toolboxPath, "analysis_results.json"))
	if err != nil {
		return err
	}

	defer f.Close()
	if _, err2 := f.Write(resultJSON); err2 != nil {
		return err
	}
	return nil
}
