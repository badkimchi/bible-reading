package audio

import (
	"app/domains/account"
	"app/util/resp"
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"net/http"
	"os"
)

type AudioController struct {
	authServ account.IAuthService
}

func NewAudioController(authServ account.IAuthService) AudioController {
	return AudioController{
		authServ: authServ,
	}
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

	chapter := chi.URLParam(r, "chapter")
	fName := c.createFileName(chapter, r)

	// Create a new file in the current working directory
	dst, err := os.Create(fmt.Sprintf("/tmp/%s.ogg", fName))
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
	fileName := chi.URLParam(r, "file_name")
	http.ServeFile(w, r, fmt.Sprintf("/tmp/%s", fileName))
}
func (c *AudioController) createFileName(chapter string, r *http.Request) string {
	userID := c.authServ.CurrentUserID(r)
	return fmt.Sprintf("%s-%s", chapter, userID)
}
