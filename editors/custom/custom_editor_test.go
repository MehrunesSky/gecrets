package custom

import (
	"github.com/MehrunesSky/gecrets/common"
	"github.com/MehrunesSky/gecrets/keyvaults/azure"
	"github.com/MehrunesSky/gecrets/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"strings"
	"testing"
)

type TestExecutor struct {
	mock.Mock
}

type File struct {
	mock.Mock
	*strings.Reader
}

func (f *File) WriteString(s string) (int, error) {
	args := f.Called(s)
	return args.Int(0), args.Error(1)
}

func (f *File) Close() error {
	return f.Called().Error(0)
}

func (f *File) Name() string {
	return f.Called().String(0)
}

func (f *File) Sync() error {
	return f.Called().Error(0)
}

func (t *TestExecutor) Execute(arg string, args ...string) error {
	mockArgs := t.Called(arg, args)
	return mockArgs.Error(0)
}

type TestFileService struct {
	mock.Mock
}

func (o *TestFileService) OpenFile(path string) (utils.File, error) {
	args := o.Called(path)
	return args.Get(0).(utils.File), args.Error(1)
}

func (o *TestFileService) CreateTempFile() (utils.File, error) {
	args := o.Called()
	return args.Get(0).(utils.File), args.Error(1)
}

func NewCustomTest(cmd string, model common.SecretI, executor *TestExecutor, fileService *TestFileService) *Custom {
	return &Custom{
		cmd:               cmd,
		model:             model,
		executor:          executor,
		fileService:       fileService,
		fileOpenerService: fileService,
	}
}

func TestCustom_ReadSecrets(t *testing.T) {
	executor := new(TestExecutor)
	fileService := new(TestFileService)

	cut := NewCustomTest(
		"",
		&azure.AzureSecret{},
		executor,
		fileService,
	)

	content := `##IGNORE
#{"key":"secretIgnore", "value" : "valueOfSecret", "contentType" : "ContentTypeOfSecret"}
{"key":"secret", "value" : "valueOfSecret", "contentType" : "ContentTypeOfSecret"}`
	reader := strings.NewReader(content)
	fileMock := &File{Reader: reader}

	fileService.On("OpenFile", "path").Return(fileMock, nil)

	secrets := cut.ReadSecrets("path")

	fileService.AssertExpectations(t)

	assert.EqualValuesf(
		t,
		common.SecretIs{azure.AzureSecret{Key: "secret", Value: "valueOfSecret", ContentType: "ContentTypeOfSecret"}},
		secrets,
		"",
	)

}

func TestCustom_Write(t *testing.T) {

	executor := new(TestExecutor)
	fileService := new(TestFileService)

	file := new(File)

	fileService.On("CreateTempFile").Return(file, nil)

	file.On("WriteString", mock.Anything).Return(0, nil)
	file.On("Name").Return("Path")
	file.On("Sync").Return(nil)
	file.On("Close").Return(nil)

	cut := NewCustomTest(
		"",
		&azure.AzureSecret{},
		executor,
		fileService,
	)

	secrets := common.SecretIs{
		azure.AzureSecret{Key: "secret", Value: "valueOfSecret", ContentType: "ContentTypeOfSecret"},
	}
	path := cut.Write(secrets)

	assert.Equal(t, "Paths", path)
	fileService.AssertExpectations(t)
	file.AssertExpectations(t)

}
