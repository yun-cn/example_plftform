package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"
// 	"strings"

// 	noonde "github.com/yanshiyason/noonde_platform"
// )

// func suggestHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := r.URL.Query()
// 	query := params.Get("q")

// 	suggestions := noonde.PlaceSuggest(query)
// 	bb, _ := json.Marshal(suggestions)

// 	w.Write(bb)
// }

// func main() {
// 	http.HandleFunc("/search", searchHandler)
// 	http.HandleFunc("/suggest", suggestHandler)
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		panic(err)
// 	}
// }

import (
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	serverPath := exPath + "/../apiserver"

	cmd := exec.Command("buffalo", "dev")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "PORT=8080")
	cmd.Dir = serverPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	cmd.Wait()
}
