package web

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLanguageSelectionAndPersistence(t *testing.T) {
	handler := LocalizedHandler("templates")
	for _, locale := range []string{"en", "zh-CN", "zh-TW"} {
		first := httptest.NewRecorder()
		handler.ServeHTTP(first, httptest.NewRequest("GET", "/carneades/?lang="+locale, nil))
		if first.Code != 200 || first.Header().Get("Content-Language") != locale {
			t.Fatalf("locale %s: %d", locale, first.Code)
		}
		if !strings.Contains(first.Body.String(), `lang="`+locale+`"`) {
			t.Fatalf("missing document language: %s", locale)
		}
		cookies := first.Result().Cookies()
		if len(cookies) != 1 {
			t.Fatal("missing language cookie")
		}
		for _, page := range []string{"help", "eval-form", "dung-form", "dataprot", "imprint"} {
			next := httptest.NewRequest("GET", "/carneades/"+page, nil)
			next.AddCookie(cookies[0])
			result := httptest.NewRecorder()
			handler.ServeHTTP(result, next)
			if result.Header().Get("Content-Language") != locale || !strings.Contains(result.Body.String(), `lang="`+locale+`"`) {
				t.Fatalf("language lost on %s/%s", locale, page)
			}
		}
	}
}
func TestInvalidLanguageCannotSelectAnArbitraryDirectory(t *testing.T) {
	handler := LocalizedHandler("templates/zh-TW")
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, httptest.NewRequest("GET", "/carneades/?lang=../../private", nil))
	if response.Header().Get("Content-Language") != "zh-TW" {
		t.Fatal("invalid locale changed the configured default")
	}
}
