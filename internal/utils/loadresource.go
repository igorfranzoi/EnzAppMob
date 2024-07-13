package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Resource: é um único recurso binário, uma imagem ou fonte
// Um ​​recurso possui um nome identificador e o conteúdo de array de bytes
// O caminho serializado de um recurso pode ser obtido, o que pode resultar em um bloqueando a operação de gravação do sistema de arquivos
type Resource interface {
	Name() string
	Content() []byte
}

// StaticResource: é um recurso empacotado compilado no aplicativo.
// Esses recursos são normalmente gerados pelo comando fyne_bundle (incluído no kit de ferramentas Fyne)
type StaticResource struct {
	StaticName    string
	StaticContent []byte
}

// Name: retorna o nome exclusivo deste recurso, geralmente correspondendo ao arquivo a partir do qual foi gerado
func (r *StaticResource) Name() string {
	return r.StaticName
}

// Content returns the bytes of the bundled resource, no compression is applied
// but any compression on the resource is retained.
func (r *StaticResource) Content() []byte {
	return r.StaticContent
}

// NewStaticResource returns a new static resource object with the specified
// name and content. Creating a new static resource in memory results in
// sharable binary data that may be serialised to the location returned by
// CachePath().
func NewStaticResource(name string, content []byte) *StaticResource {
	return &StaticResource{
		StaticName:    name,
		StaticContent: content,
	}
}

// ResourcePathLoad creates a new StaticResource in memory using the contents of the specified file.
func ResourcePathLoad(strPath string) (Resource, error) {
	bytesFile, err := os.ReadFile(filepath.Clean(strPath))

	if err != nil {
		return nil, err
	}

	rscName := filepath.Base(strPath)

	return NewStaticResource(rscName, bytesFile), nil
}

// ResourceURLStringLoad creates a new StaticResource in memory using the body of the specified URL.
func ResourceURLStringLoad(strURL string) (Resource, error) {
	resGet, err := http.Get(strURL)

	if err != nil {
		return nil, err
	}

	defer resGet.Body.Close()

	bytes, err := io.ReadAll(resGet.Body)

	if err != nil {
		return nil, err
	}

	rscName := filepath.Base(strURL)

	return NewStaticResource(rscName, bytes), nil
}
