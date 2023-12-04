package tests

import (
	"github.com/eco-challenge/service"
	"os"
	"testing"
)

func TestFileManager(t *testing.T) {
	f := service.NewFileManager()

	t.Run("🧪 Expect new FileManager not fail", func(t *testing.T) {
		if f == nil {
			t.Error("Expected mailer to be not nil")
		}
	})

	t.Run("🧪 Test GetMimeTypeFromMagicNumber", func(t *testing.T) {
		testFiles := []struct {
			filePath     string
			expectedMime string
			expectError  bool
		}{
			{"./source/img.jpg", "image/jpeg", false},
			{"./source/img.png", "image/png", false},
			{"./source/img.gif", "image/gif", false},
			{"./source/img.webp", "image/webp", false},
		}

		for _, tc := range testFiles {
			file, err := os.Open(tc.filePath)
			if err != nil {
				t.Errorf("Failed to open test file, %s", err)
			}

			magicNumberBytes := make([]byte, 12)
			if _, err := file.Read(magicNumberBytes); err != nil {
				t.Errorf("Failed to read test file, %s", err)
			}
			if err != nil {
				t.Errorf("Failed to open test file, %s", err)
			}
			mime, err := f.CheckMimeType(magicNumberBytes)
			if err != nil {
				t.Errorf("Error raise by GetMimeTypeFromMagicNumber, %s", err)
			}

			if mime != tc.expectedMime {
				t.Errorf("Error, mimetype comming from GetMimeTypeFromMagicNumber is wrong !")
			}

			err = file.Close()
			if err != nil {
				t.Errorf("Failed to close test file, %s", err)
			}
		}
	})

	t.Run("🧪 Test VerifyMimeType", func(t *testing.T) {
		res := f.VerifyMimeType("image/jpeg", []string{"image/jpeg"})
		if !res {
			t.Errorf("Valid format must return true")
		}

		res = f.VerifyMimeType("image/jpeg", []string{"image/png"})
		if res {
			t.Errorf("Wrong format must return false")
		}
	})
}
