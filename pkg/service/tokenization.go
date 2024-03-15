package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/pandishpancheta/tokenization-service/pkg/config"
	tokenization "github.com/pandishpancheta/tokenization-service/pkg/pb"
	httpretry "github.com/wabarc/ipfs-pinner/http"
)

type TokenizationService interface {
	Tokenize(ctx context.Context, req *tokenization.TokenizationRequest) (*tokenization.TokenizationResponse, error)
}

type tokenizationService struct {
	cfg *config.Config
}

func NewTokenizationService(cfg config.Config) TokenizationService {
	log.Println("Creating new tokenization service...")

	return &tokenizationService{
		cfg: &cfg,
	}
}

func (s *tokenizationService) Tokenize(ctx context.Context, req *tokenization.TokenizationRequest) (*tokenization.TokenizationResponse, error) {
	chunks := req.GetChunk()

	url := "https://api.pinata.cloud/pinning/pinFileToIPFS"

	payload := bytes.Buffer{}
	writer := multipart.NewWriter(&payload)
	part, _ := writer.CreateFormFile("file", req.GetTokenId()+".jpeg")

	// save file locally for testing
	err := os.WriteFile("temp/test.jpeg", chunks, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	// open file for testing
	file, err := os.Open("temp/test.jpeg")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	if _, err := io.Copy(part, file); err != nil {
		return nil, err
	}

	writer.Close()

	request, _ := http.NewRequest("POST", url, &payload)

	request.Header.Add("pinata_api_key", s.cfg.PinataApiKey)
	request.Header.Add("pinata_secret_api_key", s.cfg.PinataSecretApiKey)
	request.Header.Add("Content-Type", writer.FormDataContentType())

	client := httpretry.NewClient(nil)
	res, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(data, &dat); err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			return &tokenization.TokenizationResponse{
				TokenURI: "",
			}, fmt.Errorf("syntax error at byte offset %d: %s", e.Offset, e)
		}
		return &tokenization.TokenizationResponse{
			TokenURI: "",
		}, fmt.Errorf("failed to parse json: %s", err)

	}

	if out, err := dat["error"].(string); err {
		return &tokenization.TokenizationResponse{
			TokenURI: "",
		}, fmt.Errorf("pin file to Pinata failure: %s", out)

	}
	if hash, ok := dat["IpfsHash"].(string); ok {
		log.Println("IPFS hash: ", hash)
		return &tokenization.TokenizationResponse{
			TokenURI: hash,
		}, nil

	}

	return &tokenization.TokenizationResponse{
		TokenURI: "",
	}, fmt.Errorf("failed to parse json: %s", err)

}
