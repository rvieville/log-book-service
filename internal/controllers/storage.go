package controllers

import (
	"diving-log-book-service/internal/pkg/apihelper"
	"diving-log-book-service/internal/services"
	"diving-log-book-service/internal/types"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"golang.org/x/sync/errgroup"
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

	fmt.Println(input)

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
	err := r.ParseMultipartForm(50 << 20)
	if err != nil {
		apihelper.Error(w, err)
		return
	}

	diveIDStr := r.MultipartForm.Value["diveId"][0] // form data fields
	files := r.MultipartForm.File["files"]

	diveID, err := strconv.ParseUint(diveIDStr, 10, 32)
	if err != nil {
		apihelper.ValidationError(w, errors.New("dive id is not a unit"))
	}

	results := make(chan *types.UploadedFile, len(files))
	g, _ := errgroup.WithContext(r.Context())

	for _, file := range files {
		g.Go(func() error {
			body, _ := file.Open()
			payload := &types.UploadPayload{
				Bucket: os.Getenv("STORAGE_BUCKET"),
				Body:   body,
				Key:    fmt.Sprintf("%d/%s", diveID, url.QueryEscape(file.Filename)),
			}

			res, err := s.storageService.Upaload(payload)
			if err != nil {
				return err
			}

			results <- res

			return nil
		})
	}

	if err = g.Wait(); err != nil {
		apihelper.Error(w, err)
		return
	}
	close(results)

	res := make([]*types.UploadedFile, 0)
	for data := range results {
		res = append(res, data)
	}

	// apihelper.Response(w, obj)
	apihelper.Response(w, res)
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
