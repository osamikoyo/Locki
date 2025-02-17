package indexer

import (
	"strings"

	"github.com/osamikoyo/Locki/internal/models"
)

type Indexer struct{
	Word string
	Index map[string][]models.Document
}

func New(word string) *Indexer {
	return &Indexer{
		Word: word,
		Index: make(map[string][]models.Document),
	}
}

func removeStopWords(data string) string {
	replacer := strings.NewReplacer("a", "", "and", "", "the", "", "in", "", "on", "", "but", "", "or", "", "if", "")
	return replacer.Replace(data)
}

func (i *Indexer) RegisterDocument(doc models.Document) {
	doc.Data = removeStopWords(doc.Data)
}