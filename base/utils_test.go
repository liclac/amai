package base

import (
	"testing"
)

func TestNormalizeHtmlNoHtml(t *testing.T) {
	if NormalizeHtml("Aaaa") != "Aaaa" {
		t.Fail()
	}
}

func TestNormalizeHtmlStyled(t *testing.T) {
	if NormalizeHtml("A<b>aa</b>a") != "Aaaa" {
		t.Fail()
	}
}

func TestNormalizeHtmlNewlines(t *testing.T) {
	if NormalizeHtml("Two<br />Lines") != "Two\nLines" {
		t.Fail()
	}
	if NormalizeHtml("Two<br/>Lines") != "Two\nLines" {
		t.Fail()
	}
	if NormalizeHtml("Two<br>Lines") != "Two\nLines" {
		t.Fail()
	}
}
