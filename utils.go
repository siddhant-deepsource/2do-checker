package main

import (
	"encoding/json"
	"os"
	"path"
)

func createIssue(filePath string, lineNumber, _ int) {
	actualLineNumber := lineNumber + 1

	issue := Diagnostic{
		Code:    "I001",
		Message: "Found a TODO comment",
		Range: Range{
			Start: Position{
				Line: actualLineNumber,
			},
			End: Position{
				Line: actualLineNumber,
			},
		},
		RelatedInformation: []DiagnosticRelatedInformation{
			{
				Location: Location{
					URI: filePath,
					Range: Range{
						Start: Position{
							Line: actualLineNumber,
						},
						End: Position{
							Line: actualLineNumber,
						},
					},
				},
				Message: "Found a TODO comment",
			},
		},
	}
	issues = append(issues, issue)
}

func createDummyIssue(filePath string, lineNumber, _ int) {
	actualLineNumber := lineNumber + 1

	issue := Diagnostic{
		Code:    "I002",
		Message: "Fix this please",
		Range: Range{
			Start: Position{
				Line: actualLineNumber,
			},
			End: Position{
				Line: actualLineNumber,
			},
		},
		RelatedInformation: []DiagnosticRelatedInformation{
			{
				Location: Location{
					URI: filePath,
					Range: Range{
						Start: Position{
							Line: actualLineNumber,
						},
						End: Position{
							Line: actualLineNumber,
						},
					},
				},
				Message: "Fix this please",
			},
		},
	}
	issues = append(issues, issue)
}

func prepareResult() AnalysisResult {
	result := AnalysisResult{}
	result.Issues = issues
	result.IsPassed = false
	namespace := Namespace{
		Key:   "2do-checker",
		Value: 20,
	}
	metric := Metric{
		MetricCode: "IDP",
	}
	metric.Namespaces = append(metric.Namespaces, namespace)
	result.Metrics = append(result.Metrics, metric)
	error := AnalyzerError{
		HMessage: "This is an error.",
		Level:    0,
	}
	result.Errors = append(result.Errors, error)

	if len(issues) > 0 {
		result.IsPassed = true
	}

	return result
}

func writeMacroResult(result *AnalysisResult) error {
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join(toolboxPath, "analysis_results.json"), resultJSON, 0o777)
}
