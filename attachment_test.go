package allure

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func ExampleBytes() {
	a := Bytes([]byte(`{"name": "value"}`)).As(DocumentJSON)

	fmt.Println(a.Type())
	// Output: application/json
}

// Mock Attachment for testing
type mockAttachment struct {
	content   []byte
	mediaType MediaType
	openError error
	sizeHint  int64
	hasSize   bool
}

func (m *mockAttachment) Open() (io.ReadCloser, error) {
	if m.openError != nil {
		return nil, m.openError
	}
	return io.NopCloser(bytes.NewReader(m.content)), nil
}

func (m *mockAttachment) Type() MediaType {
	return m.mediaType
}

func (m *mockAttachment) SizeHint() (int64, bool) {
	return m.sizeHint, m.hasSize
}

func TestAttachmentWriter_Write(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()

	tests := []struct {
		name        string
		attachment  Attachment
		deduplicate bool
		wantErr     bool
	}{
		{
			name: "basic write without deduplication",
			attachment: &mockAttachment{
				content:   []byte("test content"),
				mediaType: TextPlain,
			},
			deduplicate: false,
			wantErr:     false,
		},
		{
			name: "write with deduplication",
			attachment: &mockAttachment{
				content:   []byte("test content"),
				mediaType: TextPlain,
			},
			deduplicate: true,
			wantErr:     false,
		},
		{
			name: "open error",
			attachment: &mockAttachment{
				openError: os.ErrInvalid,
			},
			deduplicate: false,
			wantErr:     true,
		},
		{
			name: "empty media type should guess",
			attachment: &mockAttachment{
				content:   []byte("test"),
				mediaType: "",
			},
			deduplicate: false,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			aw := newAttachmentWriter()
			filename, mediaType, err := aw.Write(tmpDir, tt.attachment, tt.deduplicate)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if filename == "" {
					t.Error("expected filename, got empty")
				}

				if mediaType == "" {
					t.Error("expected media type, got empty")
				}

				// Verify file exists
				filePath := filepath.Join(tmpDir, filename)
				if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
					t.Errorf("file not created: %s", filePath)
				}
			}
		})
	}
}

func TestAttachmentWriter_Deduplication(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	aw := newAttachmentWriter()

	// Same content, should deduplicate
	attachment := &mockAttachment{
		content:   []byte("same content"),
		mediaType: "text/plain",
	}

	// First write
	filename1, mediaType1, err1 := aw.Write(tmpDir, attachment, true)
	if err1 != nil {
		t.Fatalf("first write failed: %v", err1)
	}

	// Second write with same content
	filename2, mediaType2, err2 := aw.Write(tmpDir, attachment, true)
	if err2 != nil {
		t.Fatalf("second write failed: %v", err2)
	}

	// Should get same filename
	if filename1 != filename2 {
		t.Errorf("deduplication failed: got filenames %q and %q", filename1, filename2)
	}

	if mediaType1 != mediaType2 {
		t.Errorf("media types differ: %q vs %q", mediaType1, mediaType2)
	}

	// Different content, should not deduplicate
	diffAttachment := &mockAttachment{
		content:   []byte("different content"),
		mediaType: "text/plain",
	}

	filename3, _, err3 := aw.Write(tmpDir, diffAttachment, true)
	if err3 != nil {
		t.Fatalf("third write failed: %v", err3)
	}

	if filename1 == filename3 {
		t.Error("different content should not deduplicate")
	}
}

func TestAttachmentWriter_ConcurrentDeduplication(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	aw := newAttachmentWriter()

	attachment := &mockAttachment{
		content:   []byte("concurrent test"),
		mediaType: TextPlain,
	}

	const goroutines = 10

	// Write same content concurrently
	results := make(chan string, goroutines)
	errors := make(chan error, goroutines)

	for range goroutines {
		go func() {
			filename, _, err := aw.Write(tmpDir, attachment, true)
			if err != nil {
				errors <- err

				return
			}

			results <- filename
		}()
	}

	var filenames []string

	for range goroutines {
		select {
		case err := <-errors:
			t.Fatalf("concurrent write failed: %v", err)

		case filename := <-results:
			filenames = append(filenames, filename)
		}
	}

	// All should return same filename
	first := filenames[0]

	for i, filename := range filenames {
		if filename != first {
			t.Errorf("concurrent deduplication failed at index %d: %q != %q", i, filename, first)
		}
	}

	// And there should not be any other files in the directory
	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("read tmp dir failed: %v", err)
	}

	if len(entries) != 1 {
		t.Fatalf("more than one entry found in the dir: %d", len(entries))
	}

	p := filepath.Join(tmpDir, entries[0].Name())

	data, err := os.ReadFile(p)
	if err != nil {
		t.Fatalf("read file failed: %v", err)
	}

	if !bytes.Equal(attachment.content, data) {
		t.Fatalf("content mismatch: %v", err)
	}
}
