package post

import (
	"time"
)

//go:generate stringer -type=MoodState
type MoodState int

// All possible mood states.
const (
	MoodStateNeutral MoodState = iota
	MoodStateHappy
	MoodStateSad
	MoodStateAngry
	MoodStateHopeful
	MoodStateThrilled
	MoodStateBored
	MoodStateShy
	MoodStateComical
	MoodStateOnCloudNine
	MoodStateSmall
)

// AuditableContent types are meant to be embedded into types we want to keep a
// check on for auditing purposes
type AuditableContent struct {
	TimeCreated  time.Time
	TimeModified time.Time
	CreatedBy    string
	ModifiedBy   string
}

// Post represents a Social Media Post type.
type Post struct {
	AuditableContent // Embedded type
	Caption          string
	MessageBody      string
	URL              string
	ImageURI         string
	ThumbnailURI     string
	Keywords         []string
	Likers           []string
	AuthorMood       MoodState
}

// NewPostInput defines structured input for NewPost
type NewPostInput struct {
	Username     string
	Mood         MoodState
	Caption      string
	MessageBody  string
	Url          string
	ImageURI     string
	ThumbnailURI string
	Keywords     []string
}

// Map that holds the various mood states with keys to serve as
// aliases to their respective mood states.
var Moods map[string]MoodState

// The init() function is responsible for initializing the mood state
func init() {
	Moods = map[string]MoodState{"neutral": MoodStateNeutral, "happy": MoodStateHappy, "sad": MoodStateSad, "angry": MoodStateAngry, "hopeful": MoodStateHopeful, "thrilled": MoodStateThrilled, "bored": MoodStateBored, "shy": MoodStateShy, "comical": MoodStateComical, "cloudnine": MoodStateOnCloudNine, "small": MoodStateSmall}
}

// NewPost is responsible for creating an instance of the Post type.
func NewPost(p *NewPostInput) *Post {

	auditableContent := AuditableContent{CreatedBy: p.Username, TimeCreated: time.Now()}
	return &Post{Caption: p.Caption, MessageBody: p.MessageBody, URL: p.Url, ImageURI: p.ImageURI, ThumbnailURI: p.ThumbnailURI, AuthorMood: p.Mood, Keywords: p.Keywords, AuditableContent: auditableContent}
}
