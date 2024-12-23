package inference

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func getResponse(resp *genai.GenerateContentResponse) (result string) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				result = fmt.Sprintf("%s", part)
			}
		}
	}
	return result
}

func ImageInference(image *multipart.FileHeader) (string, error) {
	key := os.Getenv("GEMINI_KEY")
	if key == "" {
		return "", errors.New("GEMINI_KEY is not set")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(key))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	prompt, err := os.ReadFile("./src/inference/prompt.txt")
	if err != nil {
		return "", err
	}

	src, err := image.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	imageBytes, err := io.ReadAll(src)
	if err != nil {
		return "", err
	}

	resp, err := model.GenerateContent(
		ctx,
		genai.Text(prompt),
		genai.ImageData("jpeg", imageBytes))
	if err != nil {
		return "", err
	}

	return getResponse(resp), nil
}

func InferenceService(lat string, lon string, image *multipart.FileHeader) error {
	result, err := ImageInference(image)
	if err != nil {
		return err
	}

	var analysis ImageAnalysis
	if err := json.Unmarshal([]byte(result), &analysis); err != nil {
		return err
	}

	return nil
}
