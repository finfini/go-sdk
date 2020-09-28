package ekyc

import (
	"github.com/finfini/sdk/internal/difflib"
)

// Similarity contains data resulted from string similarity calculations
type Similarity struct {
	Ratio  float64
	Passed bool
}

// StringSimilarity calculate similarity between 2 strings, and check the ratio for given threshold.
func StringSimilarity(a, b string, threshold float64) Similarity {
	ratio := difflib.NewMatcher(
		difflib.ListifyString(a),
		difflib.ListifyString(b),
	).Ratio() * 100

	pass := false
	if ratio >= threshold {
		pass = true
	}

	return Similarity{
		Passed: pass,
		Ratio:  ratio,
	}
}
