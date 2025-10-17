package flags

// --- cmd/root.go ---
var (
	ForceTranslation string
	VerboseOutput    bool
	UseOop           bool
)

// --- cmd/cli.go ---
var (
	OutputToConsole bool

	UseEvalLib     bool
	UseFilterRegex bool

	Unzip             bool
	Archive           bool
	DataFileInArchive string

	Decrypt bool
	Encrypt bool
	KeyPath string
)
