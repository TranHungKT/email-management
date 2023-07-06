package utils

import (
	"bytes"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var (
	regexpSpaces = regexp.MustCompile(`[\s]+`)
)

func BindJSON(ctx *gin.Context, model interface{}) {
	if err := ctx.BindJSON(model); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
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
