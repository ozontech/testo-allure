package allure

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"io"
	"mime"
	"os"
	"path/filepath"
	"sync"

	"github.com/gabriel-vasile/mimetype"
	"github.com/ozontech/testo-allure/internal/syncutil"
	"github.com/ozontech/testo-allure/internal/uuid"
)

// MediaType stands for MIME type strings.
//
// See also [list of all official MIME types].
//
// [list of all official MIME types]: https://www.iana.org/assignments/media-types/media-types.xhtml
type MediaType string

func (m MediaType) String() string { return string(m) }

var extByMediaType = map[MediaType]string{
	ImagePNG:  ".png",
	ImageJPEG: ".jpeg",
	ImageWEBP: ".webp",
	ImageGIF:  ".gif",
	ImageSVG:  ".svg",
	ImageTIFF: ".tiff",
	ImageBMP:  ".bmp",

	VideoMP4:  ".mp4",
	VideoOGG:  ".ogg",
	VideoWebM: ".webm",

	TextPlain: ".txt",
	TextHTML:  ".html",

	TableCSV: ".csv",
	TableTSV: ".tsv",

	URIList: ".txt",

	DocumentXML:  ".xml",
	DocumentJSON: ".json",
	DocumentYAML: ".yaml",

	"application/octet-stream": ".bin",
}

// Ext returns a file extension for this media type.
// Returns empty string if extension is unknown.
func (m MediaType) Ext() string {
	// fast path for common media types.
	if ext, ok := extByMediaType[m]; ok {
		return ext
	}

	exts, _ := mime.ExtensionsByType(string(m))

	if len(exts) > 0 {
		return exts[0]
	}

	return ""
}

// Image types for attachments supported by Allure report.
//
// See also [screenshot attachments].
//
// [screenshot attachments]: https://allurereport.org/docs/attachments/#screenshots
const (
	ImagePNG  MediaType = "image/png"
	ImageJPEG MediaType = "image/jpeg"
	ImageWEBP MediaType = "image/webp"
	ImageGIF  MediaType = "image/gif"
	ImageSVG  MediaType = "image/svg+xml"
	ImageTIFF MediaType = "image/tiff"
	ImageBMP  MediaType = "image/bmp"
)

// Video types for attachments supported by Allure report.
//
// See also [video attachments].
//
// [video attachments]: https://allurereport.org/docs/attachments/#videos
const (
	VideoMP4  MediaType = "video/mp4"
	VideoOGG  MediaType = "video/ogg"
	VideoWebM MediaType = "video/webm"
)

// Text types for attachments supported by Allure report.
//
// See also [text attachments].
//
// [text attachments]: https://allurereport.org/docs/attachments/#text
const (
	TextPlain MediaType = "text/plain"
	TextHTML  MediaType = "text/html"
)

// Table types for attachments supported by Allure report.
//
// See also [table attachments].
//
// [table attachments]: https://allurereport.org/docs/attachments/#tables
const (
	TableCSV MediaType = "text/csv"
	TableTSV MediaType = "text/tab-separated-values"
)

// URIList is uri list type for attachments.
//
// See also [uri lists attachments].
//
// [uri lists attachments]: https://allurereport.org/docs/attachments/#uri-lists
const URIList MediaType = "text/uri-list"

// Document types for attachments supported by Allure report.
//
// See also [document attachments].
//
// [document attachments]: https://allurereport.org/docs/attachments/#documents
const (
	DocumentXML  MediaType = "text/xml"
	DocumentJSON MediaType = "application/json"
	DocumentYAML MediaType = "application/yaml"
)

// Attachment to add into report.
//
// See [Allure attachments] for more information.
//
// [Allure attachments]: https://allurereport.org/docs/attachments/
type Attachment interface {
	// Open attachment for reading.
	Open() (io.ReadCloser, error)

	// Type returns the media type of the content.
	Type() MediaType

	// SizeHint returns attachment size as in bytes count.
	SizeHint() (size int64, ok bool)
}

// AttachmentBytes is an attachment which stores its contents in-memory.
// Consider using [AttachmentFile] for large, static files.
type AttachmentBytes struct {
	Data      []byte
	MediaType MediaType
}

// Bytes creates a new bytes attachment from the given bytes or string.
//
// See [AttachmentBytes.As] to specify media type for preview in Allure report.
// If not specified, media type will be guessed based on the data contents with fallback
// to "application/octet-stream" when writing attachment to the report.
//
// Media type guesser is backed by [gabriel-vasile/mimetype].
//
//	allure.Bytes([]byte{...}).As(allure.DocumentJSON)
//	allure.Bytes(`{"name":"value"}`)
//	allure.Bytes("Hello, world!")
//
// [gabriel-vasile/mimetype]: https://github.com/gabriel-vasile/mimetype
func Bytes[BytesLike ~string | ~[]byte](data BytesLike) AttachmentBytes {
	return AttachmentBytes{
		Data: []byte(data),
	}
}

// As returns new attachment with the given media type.
func (b AttachmentBytes) As(mediaType MediaType) AttachmentBytes {
	b.MediaType = mediaType

	return b
}

// Open attachment for reading.
func (b AttachmentBytes) Open() (io.ReadCloser, error) {
	return io.NopCloser(bytes.NewReader(b.Data)), nil
}

// Type returns the media type of the content.
func (b AttachmentBytes) Type() MediaType { return b.MediaType }

// SizeHint returns bytes data len.
func (b AttachmentBytes) SizeHint() (int64, bool) {
	return int64(len(b.Data)), true
}

// AttachmentFile is an attachment
// stored in the file located at path.
type AttachmentFile struct {
	Path      string
	MediaType MediaType
}

// File creates a new attachment from the given file path.
//
// Media type will be guessed based on the file extension.
// If file extension is missing or unknown, media type will be guessed later based on
// the file contents when writing attachment to the report.
//
// Use [AttachmentFile.Bytes] method to convert it to [AttachmentBytes] immediately.
//
//	allure.File("/home/user/file.txt")
//	allure.File("/home/user/file.json").As(allure.DocumentYAML)
//	allure.File("/home/user/file")
func File(path string) AttachmentFile {
	ext := filepath.Ext(path)

	mediaType, _, _ := mime.ParseMediaType(mime.TypeByExtension(ext))

	return AttachmentFile{
		Path:      path,
		MediaType: MediaType(mediaType),
	}
}

// As returns new attachment with the given media type.
func (f AttachmentFile) As(mediaType MediaType) AttachmentFile {
	f.MediaType = mediaType

	return f
}

// Open attachment for reading.
func (f AttachmentFile) Open() (io.ReadCloser, error) {
	return os.OpenFile(f.Path, os.O_RDONLY, permFile)
}

// Type returns the media type of the content.
func (f AttachmentFile) Type() MediaType {
	return f.MediaType
}

// SizeHint returns size of the file pointer by path.
func (f AttachmentFile) SizeHint() (int64, bool) {
	info, err := os.Stat(f.Path)
	if err != nil {
		return 0, false
	}

	return info.Size(), true
}

var globalAttachmentWriter = newAttachmentWriter()

type hashSumForDir struct {
	HashSum uint64
	Dir     string
}

type attachmentWriteResult struct {
	Filename string
	Type     MediaType
}

type attachmentWriter struct {
	nameByHash sync.Map
	locks      syncutil.KeyedMutex[hashSumForDir]
	buffers    sync.Pool
}

func newAttachmentWriter() *attachmentWriter {
	return &attachmentWriter{
		buffers: sync.Pool{
			New: func() any { return new(bytes.Buffer) },
		},
	}
}

func (aw *attachmentWriter) Write(
	outputDir string,
	at Attachment,
	deduplicate bool,
) (filename string, mediaType MediaType, err error) {
	if deduplicate {
		return aw.writeDeduplicating(outputDir, at)
	}

	return aw.write(outputDir, at)
}

//nolint:funlen // TODO: make it shorter
func (aw *attachmentWriter) write(
	outputDir string,
	at Attachment,
) (filename string, mediaType MediaType, err error) {
	src, err := at.Open()
	if err != nil {
		return "", "", fmt.Errorf("open: %w", err)
	}

	//nolint:errcheck // even if it fails, it's not worth failing the writer
	defer src.Close()

	mediaType = at.Type()

	stem := uuid.New().String() + "-attachment"

	if mediaType == "" {
		file, err := os.CreateTemp("", stem+".tmp")
		if err != nil {
			return "", "", err
		}

		_, err = io.Copy(file, src)
		if err != nil {
			_ = file.Close()
			_ = os.Remove(file.Name())

			return "", "", err
		}

		_ = file.Close()

		mediaType = guessMediaTypeFile(file.Name())

		filename = stem + mediaType.Ext()

		newPath := filepath.Join(outputDir, filename)

		// rename is an atomic operation on unix platforms.
		err = os.Rename(file.Name(), newPath)
		if err != nil {
			_ = os.Remove(file.Name())

			// on windows it may be the case, since rename is not atomic there.
			_ = os.Remove(newPath)

			return "", "", err
		}

		return filename, mediaType, nil
	}

	filename = stem + mediaType.Ext()

	file, err := os.OpenFile(filepath.Join(outputDir, filename), os.O_WRONLY|os.O_CREATE, permFile)
	if err != nil {
		return "", "", err
	}

	defer func() { _ = file.Close() }()

	_, err = io.Copy(file, src)
	if err != nil {
		_ = os.Remove(file.Name())

		return "", "", err
	}

	return filename, mediaType, nil
}

//nolint:funlen // TODO: make it shorter
func (aw *attachmentWriter) writeDeduplicating(
	outputDir string,
	at Attachment,
) (filename string, mediaType MediaType, err error) {
	src, err := at.Open()
	if err != nil {
		return "", "", fmt.Errorf("open: %w", err)
	}

	//nolint:errcheck // even if it fails, it's not worth failing the writer
	defer src.Close()

	buffer := aw.buffers.Get().(*bytes.Buffer)

	defer func() {
		buffer.Reset()
		aw.buffers.Put(buffer)
	}()

	if size, ok := at.SizeHint(); ok && size > 0 {
		buffer.Grow(int(size))
	}

	h := fnv.New64a()

	_, err = io.Copy(io.MultiWriter(h, buffer), src)
	if err != nil {
		return "", "", fmt.Errorf("read: %w", err)
	}

	hashSum := hashSumForDir{
		HashSum: h.Sum64(),
		Dir:     outputDir,
	}

	unlock := aw.locks.Lock(hashSum)
	defer unlock()

	if v, ok := aw.nameByHash.Load(hashSum); ok {
		res := v.(attachmentWriteResult)

		return res.Filename, res.Type, nil
	}

	data := buffer.Bytes()

	mediaType = at.Type()
	if mediaType == "" {
		mediaType = guessMediaType(data)
	}

	filename = uuid.New().String() + "-attachment" + mediaType.Ext()

	err = os.WriteFile(filepath.Join(outputDir, filename), buffer.Bytes(), permFile)
	if err != nil {
		return "", "", fmt.Errorf("write: %w", err)
	}

	aw.nameByHash.Store(hashSum, attachmentWriteResult{
		Filename: filename,
		Type:     mediaType,
	})

	return filename, mediaType, nil
}

func guessMediaType(data []byte) MediaType {
	detected := mimetype.Detect(data)

	m, _, _ := mime.ParseMediaType(detected.String())

	return MediaType(m)
}

func guessMediaTypeFile(path string) MediaType {
	detected, _ := mimetype.DetectFile(path)

	m, _, _ := mime.ParseMediaType(detected.String())

	return MediaType(m)
}

func trimmedAttachment(
	data []byte,
	mediaType MediaType,
	limit int64,
) AttachmentBytes {
	if len(data) <= int(limit) {
		return Bytes(data).As(mediaType)
	}

	// we can't use format like "want %d, got %d" because len(data)
	// isn't always a "full" attachment.

	suffix := fmt.Sprintf("...\n\n...size exceeds %d bytes limit", limit)

	// TODO: it's possible to avoid copying
	// if we would store some sort of flag
	// rather than appending text to bytes slice.

	buf := make([]byte, int(limit)+len(suffix))

	copy(buf, data[:limit])
	copy(buf[limit:], suffix)

	return Bytes(buf).As(TextPlain)
}
