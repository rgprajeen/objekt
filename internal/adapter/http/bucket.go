package http // import "go.prajeen.com/objekt/internal/adapter/http"

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.prajeen.com/objekt/internal/core/domain"
	"go.prajeen.com/objekt/internal/core/port"
)

type BucketHandler struct {
	svc port.BucketService
}

func NewBucketHandler(svc port.BucketService) *BucketHandler {
	return &BucketHandler{svc: svc}
}

type createBucketRequest struct {
	Name   string            `json:"name"`
	Type   domain.BucketType `json:"type"`
	Region string            `json:"region"`
}

func (h *BucketHandler) CreateBucket(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var requestBody createBucketRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bucket := &domain.Bucket{
		Name:   requestBody.Name,
		Type:   requestBody.Type,
		Region: requestBody.Region,
	}

	bucket, err = h.svc.CreateBucket(r.Context(), bucket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bucket)
}

func (h *BucketHandler) GetBucket(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bucketID := p.ByName("id")
	bucket, err := h.svc.GetBucket(r.Context(), bucketID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(bucket)
}

func (h *BucketHandler) ListBuckets(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	buckets, err := h.svc.ListBuckets(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(buckets)
}

func (h *BucketHandler) DeleteBucket(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bucketID := p.ByName("id")
	err := h.svc.DeleteBucket(r.Context(), bucketID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
