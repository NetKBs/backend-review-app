package inference

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"image/jpeg"
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/NetKBs/backend-reviewapp/geoapify"
	"github.com/NetKBs/backend-reviewapp/src/social/place"
	"github.com/disintegration/imaging"
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

func OptimizeImage(image *multipart.FileHeader) ([]byte, error) {
	src, err := image.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	img, err := imaging.Decode(src)
	if err != nil {
		return nil, err
	}

	resizedImg := imaging.Resize(img, 500, 0, imaging.Lanczos)

	var buf bytes.Buffer
	err = jpeg.Encode(&buf, resizedImg, &jpeg.Options{Quality: 85})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
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

	opmtimizedImage, err := OptimizeImage(image)
	if err != nil {
		return "", err
	}

	resp, err := model.GenerateContent(
		ctx,
		genai.Text(prompt),
		genai.ImageData("jpeg", opmtimizedImage))
	if err != nil {
		return "", err
	}

	return getResponse(resp), nil
}

func InferenceService(lat string, lon string, image *multipart.FileHeader) (geoapify.Places, error) {
	result, err := ImageInference(image)
	if err != nil {
		return nil, err
	}

	var analysis ImageAnalysis
	if err := json.Unmarshal([]byte(result), &analysis); err != nil {
		return nil, err
	}
	log.Println(analysis)

	ctx := context.TODO()
	placesDTO, err := place.GetPlacesByCoordsService(ctx, analysis.Categories, lat, lon)
	if err != nil {
		return nil, err
	}
	placesData := placesDTO.Data

	if placesData == nil {
		return []geoapify.Place{}, nil
	}

	if len(analysis.VisibleText) > 0 {
		var matchedPlaces []geoapify.Place
		var otherPlaces []geoapify.Place

		for _, place := range placesData {
			matched := false
			for _, text := range analysis.VisibleText {
				if strings.Contains(strings.ToLower(place.Name), strings.ToLower(text)) {
					matched = true
					break
				}
			}
			if matched {
				matchedPlaces = append(matchedPlaces, place)
			} else {
				otherPlaces = append(otherPlaces, place)
			}
		}

		placesData = append(matchedPlaces, otherPlaces...)
	}

	return placesData, nil
}
