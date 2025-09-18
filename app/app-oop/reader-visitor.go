package app_oop

// ReaderVisitor represents a type that implements visitor pattern to visit Reader interface
type ReaderVisitor interface {
	// Visit Readin
	DoForReadin(r *Readin)
	// Visit ReadinUnzip
	DoForReadinUnzip(r *ReadinUnzip)
}
