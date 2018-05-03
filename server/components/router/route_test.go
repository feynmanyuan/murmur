package router

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStaticMapping(t *testing.T) {
	StaticMapping("/static/", "/static")
	StaticMapping("static1/", "static/")
	StaticMapping("/static2", "static")

	m := routerServiceInstance.getStaticPath()

	v, exists := m["/static/"]

	assert.True(t, exists)
	assert.Equal(t, "/static/", v)

	v, exists = m["/static1/"]

	assert.True(t, exists)
	assert.Equal(t, "/static/", v)

	v, exists = m["/static2/"]

	assert.True(t, exists)
	assert.Equal(t, "/static/", v)
}

func TestRegister(t *testing.T) {

}
