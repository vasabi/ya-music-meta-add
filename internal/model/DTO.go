package model

type Track struct {
	Id                string
	RealId            string
	Title             string
	DurationMillis    string
	Available         string
	FileSize          string
	Token             string
	IsOffline         string
	CoverUri          string
	ContentWarning    string
	IsLyricsAvailable string
	Type              string
	TrackOptions      string
	PubDate           string
}

type TrackAlbum struct {
	AutoId        string
	TrackId       string
	AlbumId       string
	TrackPosition string
	AlbumVolume   string
}

type TrackLyrics struct {
	TrackId    string
	Lyrics     string
	FullLyrics string
	Url        string
	HasRights  string
}

type Album struct {
	Id            string
	Title         string
	ArtistsString string
	AlbumVersion  string
	Year          string
	GenreId       string
	GenreTitle    string
	CoverUri      string
	TrackCount    string
	AlbumOptions  string
}
