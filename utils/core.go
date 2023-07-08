package utils

import (
	"bytes"
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	regexpSpaces = regexp.MustCompile(`[\s]+`)
)

func BindJSON(ctx *gin.Context, model interface{}) error {
	if err := ctx.BindJSON(model); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return errors.New("BIND_DATA_ERROR")
	}
	return nil
}

func ValidateByStruct(ctx *gin.Context, payload interface{}) error {
	validationErr := validator.New().Struct(payload)

	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		ctx.Abort()
		return errors.New("VALIDATION_ERROR")
	}
	return nil
}

func BindJSONAndValidateByStruct(ctx *gin.Context, payload interface{}) error {
	err := BindJSON(ctx, payload)
	if err != nil {
		return err
	}

	err = ValidateByStruct(ctx, payload)
	if err != nil {
		return err
	}
	return nil

}

func NormalizeTags(tags []string) []string {
	var (
		out  []string
		dash = []byte("-")
	)

	for _, t := range tags {
		rep := regexpSpaces.ReplaceAll(bytes.TrimSpace([]byte(t)), dash)

		if len(rep) > 0 {
			out = append(out, string(rep))
		}
	}

	return out
}
