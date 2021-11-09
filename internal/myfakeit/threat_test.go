package myfakeit

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestAppExecBlockedEvents(t *testing.T) {
	gofakeit.Seed(11)

	threats := AppExecBlockedEvents(2, 2)
	assert.Equal(t, 2, len(threats))
	assert.Equal(t, threats[0].Guid, threats[1].Guid)
	assert.Equal(t, "C:\\Program Files\\ASDA_Soft_V5\\ASDA_Soft_V5.exe", threats[0].File)
	assert.Equal(t, "032a4e38b86d929ca23b3755cff582df406d4cbf758e3c91715e774dc270ede2", threats[0].Hash)
	assert.Equal(t, "Virus", threats[0].Type)
	assert.Equal(t, "PE_TEST_VIRUS", threats[0].Name)
	assert.Equal(t, "C:\\Windows\\SysWOW64\\LRXSW 3.45\\LRXSW 3.45.exe", threats[1].File)
	assert.Equal(t, "2ec1f96e1a3b197dcb71baacb9abe6a1f688a2642b8d0d6e011701d0c544810b", threats[1].Hash)
	assert.Equal(t, "Virus", threats[1].Type)
	assert.Equal(t, "PE_TEST_VIRUS", threats[1].Name)
}
