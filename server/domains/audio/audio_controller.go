package audio

import (
	"app/util/resp"
	"fmt"
	"io"
	"net/http"
	"os"
)

type AudioController struct {
}

func NewAudioController() AudioController {
	return AudioController{}
}

func (c *AudioController) UploadAudio(w http.ResponseWriter, r *http.Request) {
	// The argument in ParseMultipartForm is the max memory size
	// that will be used to parse the form; overflow will be stored in temp files
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("audioFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Create a new file in the current working directory
	dst, err := os.Create("/tmp/uploaded_audio.ogg")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the filesystem at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Data(w, r, "success")
}

func (c *AudioController) GetAudio(w http.ResponseWriter, r *http.Request) {
	var abc = "asd"
	fmt.Println(abc)
	fmt.Println(abc)
	//w.Header().Set("Content-Type", "audio/ogg")
	//w.Header().Set("Content-Disposition", "attachment; filename=uploaded_audio.ogg")
	http.ServeFile(w, r, "/tmp/uploaded_audio.ogg")
}
