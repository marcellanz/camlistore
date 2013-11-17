package index

import (
	"errors"
	"sync"

	"camlistore.org/pkg/blob"
	"camlistore.org/pkg/types/camtypes"
)

// Corpus is an in-memory summary of all of a user's blobs' metadata.
type Corpus struct {
	mu    sync.RWMutex
	strs  map[string]string // interned strings
	blobs map[blob.Ref]camtypes.BlobMeta
	// TODO: add GoLLRB to third_party; keep sorted BlobMeta
	files      map[blob.Ref]*FileMeta
	permanodes map[blob.Ref]*PermanodeMeta

	// TOOD: use deletedCache instead?
	deletedBy map[blob.Ref]blob.Ref // key is deleted by value
}

type FileMeta struct {
	size      int64
	mimeType  string
	wholeRefs []blob.Ref
}

type PermanodeMeta struct {
	OwnerKeyId string
	// Claims     ClaimList
}

func newCorpus() *Corpus {
	return &Corpus{
		blobs:      make(map[blob.Ref]camtypes.BlobMeta),
		files:      make(map[blob.Ref]*FileMeta),
		permanodes: make(map[blob.Ref]*PermanodeMeta),
		deletedBy:  make(map[blob.Ref]blob.Ref),
	}
}

func NewCorpusFromStorage(s Storage) (*Corpus, error) {
	if s == nil {
		return nil, errors.New("storage is nil")
	}
	c := newCorpus()
	err := enumerateBlobMeta(s, func(bm camtypes.BlobMeta) error {
		c.blobs[bm.Ref] = bm
		// TODO: populate blobref intern table
		return nil
	})
	if err != nil {
		return nil, err
	}
	// TODO: slurp more from storage
	return c, nil
}

func (x *Index) KeepInMemory() (*Corpus, error) {
	var err error
	x.corpus, err = NewCorpusFromStorage(x.s)
	return x.corpus, err
}

func (c *Corpus) EnumerateBlobMeta(ch chan<- camtypes.BlobMeta) error {
	defer close(ch)
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, bm := range c.blobs {
		ch <- bm
	}
	return nil
}