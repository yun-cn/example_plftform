package noonde

// FileStorage stores a file somewhere and returns a publicly accessible URL.
type FileStorage interface {
	Store(filename string, blob []byte) (string, error)
}
