package http // import "github.com/attoleap/objekt/internal/adapter/http"

import (
	"encoding/json"
	"net/http"

	"github.com/attoleap/objekt/internal/core/domain"
	"github.com/attoleap/objekt/internal/core/port"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
)

type FileHandler struct {
	log    *zerolog.Logger
	router *httprouter.Router
	svc    port.FileService
}

func NewFileHandler(log *zerolog.Logger, router *httprouter.Router, svc port.FileService) *FileHandler {
	return &FileHandler{
		log:    log,
		router: router,
		svc:    svc,
	}
}

type createFileRequest struct {
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	BucketName string `json:"bucket_name"`
	MimeType   string `json:"mime_type"`
}

func (h *FileHandler) AddRoutes() {
	h.router.POST("/files", contentTypeMiddleware(h.CreateFile, []string{ContentTypeJSON}))
	h.router.DELETE("/files/:id", h.DeleteFile)
	h.router.GET("/files/:id", h.GetFile)
	h.router.GET("/buckets/:bucket_id/files", h.GetFilesByBucketName)
}

func (h *FileHandler) CreateFile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var requestBody createFileRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		h.log.Err(err).Msg("failed to decode request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file := &domain.File{
		Name:       requestBody.Name,
		Size:       requestBody.Size,
		BucketName: requestBody.BucketName,
		MimeType:   requestBody.MimeType,
	}

	file, err = h.svc.CreateFile(r.Context(), file)
	if err != nil {
		h.log.Err(err).Msg("failed to create file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(HeaderContentType, ContentTypeJSON)
	w.Header().Set(HeaderLocation, "/files/"+file.ID.String())
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(file); err != nil {
		h.log.Err(err).Msg("failed to encode response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *FileHandler) DeleteFile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	err := h.svc.DeleteFile(r.Context(), id)
	if err != nil {
		h.log.Err(err).Msg("failed to delete file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *FileHandler) GetFile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	file, err := h.svc.GetFile(r.Context(), id)
	if err != nil {
		h.log.Err(err).Msg("failed to get file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(HeaderContentType, ContentTypeJSON)
	if err = json.NewEncoder(w).Encode(file); err != nil {
		h.log.Err(err).Msg("failed to encode response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *FileHandler) GetFilesByBucketName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bucketID := p.ByName("bucket_id")
	files, err := h.svc.GetFilesByBucketID(r.Context(), bucketID)
	if err != nil {
		h.log.Err(err).Msg("failed to get files by bucket ID")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(HeaderContentType, ContentTypeJSON)
	if err = json.NewEncoder(w).Encode(files); err != nil {
		h.log.Err(err).Msg("failed to encode response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
