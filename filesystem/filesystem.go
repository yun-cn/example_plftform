package filesystem

import "io/ioutil"

// Storage ..
type Storage struct{}

// NewStorage ..
func NewStorage() *Storage {
	return &Storage{}
}

// Store stores a base64 encoded blob into the file system and returns it's path
func (s *Storage) Store(filename string, blob []byte) (string, error) {
	err := ioutil.WriteFile("../public/uploads/"+filename, blob, 0644)
	if err != nil {
		return "", err
	}

	return "/public/uploads/" + filename, err
}
