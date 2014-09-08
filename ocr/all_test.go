package ocr_test

import "testing"
import "github.com/otiai10/ocr-kcwidget/ocr"
import . "github.com/otiai10/mint"

func TestProcess(t *testing.T) {
	result, err := ocr.Process("hogeeee")
	Expect(t, err).ToBe(nil)
	Expect(t, result).ToBe("")
}
