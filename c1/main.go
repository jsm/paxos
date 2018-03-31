package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jsm/paxos/c1/support"
)

type messageRequest struct {
	Message string `json:"message"`
}

var digestToMessageMap map[string]string = map[string]string{}

func main() {
	// Initialize
	r := chi.NewRouter()

	// Setup Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	r.Use(setContentJSON)

	// Setup routes
	r.Get("/messages/{digest}", getMessagesHandler)
	r.Post("/messages", postMessagesHandler)

	http.ListenAndServe(":3000", r)
}

// Custom middleware
func setContentJSON(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func getMessagesHandler(w http.ResponseWriter, r *http.Request) {
	digest := chi.URLParam(r, "digest")

	message, err := getMessageFromDigest(digest)
	if err != nil {
		support.HandleError(err, w)
		return
	}

	respJ := support.CreateResponseJSON(map[string]string{
		"message": message,
	}, nil)

	w.WriteHeader(http.StatusOK)
	w.Write(respJ)

	return
}

func postMessagesHandler(w http.ResponseWriter, r *http.Request) {
	var request messageRequest
	if !support.HandleJSONDecode(&request, w, r) {
		return
	}

	hash, err := generateMessageDigest(request.Message)
	if err != nil {
		support.HandleError(err, w)
		return
	}

	respJ := support.CreateResponseJSON(map[string]string{
		"digest": hash,
	}, nil)

	w.WriteHeader(http.StatusOK)
	w.Write(respJ)

	return
}

func generateMessageDigest(message string) (string, error) {
	hasher := sha256.New()
	if _, err := hasher.Write([]byte(message)); err != nil {
		return "", err
	}

	digest := hex.EncodeToString(hasher.Sum(nil))
	digestToMessageMap[digest] = message

	return digest, nil
}

func getMessageFromDigest(digest string) (string, error) {
	message, ok := digestToMessageMap[digest]
	if !ok {
		return "", support.ErrDigestNotFound(digest)
	}

	return message, nil
}
