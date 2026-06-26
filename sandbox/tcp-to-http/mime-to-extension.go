package main

var mimeToExtension = map[string]string{
	// Documents
	"application/pdf":  "pdf",
	"text/plain":       "txt",
	"text/html":        "html",
	"text/css":         "css",
	"text/csv":         "csv",
	"application/json": "json",
	"application/xml":  "xml",
	"text/xml":         "xml",

	// Microsoft Office
	"application/msword": "doc",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": "docx",
	"application/vnd.ms-excel": "xls",
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         "xlsx",
	"application/vnd.ms-powerpoint":                                             "ppt",
	"application/vnd.openxmlformats-officedocument.presentationml.presentation": "pptx",

	// Images
	"image/jpeg":    "jpg",
	"image/png":     "png",
	"image/gif":     "gif",
	"image/webp":    "webp",
	"image/svg+xml": "svg",
	"image/bmp":     "bmp",
	"image/tiff":    "tiff",
	"image/x-icon":  "ico",
	"image/heic":    "heic",

	// Audio
	"audio/mpeg": "mp3",
	"audio/wav":  "wav",
	"audio/ogg":  "ogg",
	"audio/flac": "flac",
	"audio/aac":  "aac",

	// Video
	"video/mp4":        "mp4",
	"video/mpeg":       "mpeg",
	"video/webm":       "webm",
	"video/x-msvideo":  "avi",
	"video/quicktime":  "mov",
	"video/x-matroska": "mkv",

	// Archives
	"application/zip":              "zip",
	"application/x-rar-compressed": "rar",
	"application/x-7z-compressed":  "7z",
	"application/gzip":             "gz",
	"application/x-tar":            "tar",

	// Other
	"application/octet-stream": "bin",
}
