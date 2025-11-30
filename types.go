package chunkline

import (
	"fmt"
	"github.com/zeebo/xxh3"
	"time"
)

type Endpoint struct {
	Iterator string `json:"iterator"`
	Body     string `json:"body"`
}

type Manifest struct {
	Version    string   `json:"version"`
	ChunkSize  int64    `json:"chunk_size"`
	FirstChunk int      `json:"first_chunk"`
	LastChunk  int      `json:"last_chunk"`
	Ascending  Endpoint `json:"ascending"`
	Descending Endpoint `json:"descending"`
	Metadata   any      `json:"metadata"`
}

func (m Manifest) time2Chunk(t time.Time) int64 {
	return t.Unix() / m.ChunkSize
}

func (m Manifest) chunk2Time(c int64) time.Time {
	return time.Unix(c*m.ChunkSize, 0)
}

type ItrNode string

type BodyChunk struct {
	URI     string     `json:"uri"`
	ChunkID int64      `json:"chunk_id"`
	Items   []BodyItem `json:"items"`
}

type BodyItem struct {
	Timestamp   time.Time `json:"timestamp"`
	Content     string    `json:"content"`
	ContentType string    `json:"content_type"`
	Href        string    `json:"href"`
}

func (b BodyItem) ID() string {
	if b.Href != "" {
		return b.Href
	}
	hash := xxh3.HashString(b.Content)
	return fmt.Sprintf("urn:xxh3:%x", hash)
}
