package flags

// --- cmd/root.go ---
var (
	ForceTranslation string
	OutputToConsole  bool
	VerboseOutput    bool
	UseOop           bool

	UseEvalLib     bool
	UseFilterRegex bool

	Unzip             bool
	Archive           bool
	DataFileInArchive string

	Decrypt bool
	Encrypt bool
	KeyPath string
)
