package handlers

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// remember to add -t 0 flag  like below
// go test -timeout 0 -run TestAlbumsRouter
// go test -run TestAlbumsRouter
func TestAlbumsRouter(t *testing.T) {
	router := gin.Default()
	Register(router)
	req := httptest.NewRequest("GET", "/api/albums", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)

	if body == nil || err != nil {
			t.Fatalf("%v", err)
	}
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, body != nil)
}