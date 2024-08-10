package controllers

import (
	"diving-log-book-service/internal/pkg/apihelper"
	"diving-log-book-service/internal/services"
	"diving-log-book-service/internal/types"
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
)

type StorageController struct {
	storageService *services.StorageService
}

func NewStorageController(storageService *services.StorageService) *StorageController {
	return &StorageController{
		storageService,
	}
}

func (s StorageController) Get(w http.ResponseWriter, r *http.Request) {
	input := &types.GetUrl{
		Bucket: r.URL.Query().Get("bucket"),
		Key:    r.URL.Query().Get("key"),
	}

	err := apihelper.Validate(input)
	if err != nil {
		apihelper.ValidationError(w, err.(validator.ValidationErrors))
		return
	}

	url, err := s.storageService.GetUrl(input)
	if err != nil {
		apihelper.Error(w, err)
		return
	}

	apihelper.Response(w, url)
}

func (s StorageController) Upload(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	file, _, err := r.FormFile("file")
	if err != nil {
		apihelper.Error(w, err)
		return
	}

	payload := &types.UploadPayload{
		Bucket: os.Getenv("STORAGE_BUCKET"),
		Name:   name,
		Body:   file,
	}

	obj, err := s.storageService.Upaload(payload)
	if err != nil {
		apihelper.Error(w, err)
		return
	}

	apihelper.Response(w, obj)
}

func (s StorageController) Delete(w http.ResponseWriter, r *http.Request) {
	var body *types.DeleteObject

	json.NewDecoder(r.Body).Decode(&body)
	validationErr := apihelper.Validate(body)
	if validationErr != nil {
		apihelper.ValidationError(w, validationErr.(validator.ValidationErrors))
		return
	}

	err := s.storageService.Delete(body)
	if err != nil {
		apihelper.Error(w, err)
	}

	apihelper.Response(w, "Media deleted successfully")
}
