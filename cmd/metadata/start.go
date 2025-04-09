package metadata

import (
	"github.com/alicebob/sqlittle"
	"github.com/bogem/id3v2/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
	"ya-music-meta-add/internal"
	"ya-music-meta-add/internal/model"
)

var start = &cobra.Command{
	Use:   "start",
	Short: "start decomposition",
	Long:  "start decomposition",
	Args:  cobra.MatchAll(cobra.ExactArgs(0)),
	Run: func(cmd *cobra.Command, args []string) {
		const (
			trackT       = "T_Track"
			trackAlbumT  = "T_TrackAlbum"
			trackLyricsT = "T_TrackLyrics"
			albumT       = "T_Album"
		)

		db, err := sqlittle.Open(internal.GlobalConfig.DBPath)
		if err != nil {
			logrus.Fatalf("Cannot read db file: %s, err: %s", internal.GlobalConfig.DBPath, err.Error())
		}

		defer func() {
			_ = db.Close()
		}()

		tracks := getTracks(db, trackT)
		tAlbums := getTrackAlbums(db, trackAlbumT)
		tLyrics := getTrackLyrics(db, trackLyricsT)
		albums := getAlbums(db, albumT)

		logrus.Info(tracks, tAlbums, tLyrics, albums)

		err = filepath.Walk(internal.GlobalConfig.MusicPath, func(wPath string, info os.FileInfo, err error) error {
			// Выводится название файла
			if wPath != internal.GlobalConfig.MusicPath && !info.IsDir() && info.Name() != ".DS_Store" {
				logrus.Infof("path: %s", wPath)

				tags, err := id3v2.Open(wPath, id3v2.Options{Parse: true})
				if err != nil {
					logrus.Warnf("Cannot get tags: %s, err: %s", wPath, err.Error())
					return nil
				}

				logrus.Infof("%+v", tags)

				fileNameSplit := strings.Split(info.Name(), ".")
				fileNameSplit = fileNameSplit[:len(fileNameSplit)-1]
				trackID := strings.Join(fileNameSplit, "")

				logrus.Infof("%+v", trackID)

				tags.SetArtist(albums[tAlbums[trackID].AlbumId].ArtistsString)
				tags.SetAlbum(albums[tAlbums[trackID].AlbumId].Title)
				tags.SetGenre(albums[tAlbums[trackID].AlbumId].GenreId)
				tags.SetTitle(tracks[trackID].Title)
				tags.SetYear(albums[tAlbums[trackID].AlbumId].Year)

				lyrics := id3v2.UnsynchronisedLyricsFrame{
					Encoding:          id3v2.EncodingUTF8,
					Language:          "eng",
					ContentDescriptor: tLyrics[trackID].Lyrics,
					Lyrics:            tLyrics[trackID].FullLyrics,
				}
				tags.AddUnsynchronisedLyricsFrame(lyrics)

				err = tags.Save()
				if err != nil {
					logrus.Warnf("Cannot save tags: %s, tags: %+v, err: %s", wPath, tags, err.Error())
					return nil
				}

				_ = tags.Close()
			}

			return nil
		})
		if err != nil {
			logrus.Fatalf("Wrong music path: %s, err: %s", internal.GlobalConfig.MusicPath, err.Error())
		}
	},
}

func getTracks(db *sqlittle.DB, currentTable string) (list map[string]model.Track) {
	list = make(map[string]model.Track)
	var selectAllFunc = func(r sqlittle.Row) {
		var obj model.Track
		err := r.Scan(
			&obj.Id,
			&obj.RealId,
			&obj.Title,
			&obj.DurationMillis,
			&obj.Available,
			&obj.FileSize,
			&obj.Token,
			&obj.IsOffline,
			&obj.CoverUri,
			&obj.ContentWarning,
			&obj.IsLyricsAvailable,
			&obj.Type,
			&obj.TrackOptions,
			&obj.PubDate,
		)
		if err != nil {
			logrus.Fatalf("Cannot scan rows from table: %s, err: %s", currentTable, err.Error())
		}

		list[obj.Id] = obj
	}

	col, err := db.Columns(currentTable)
	if err != nil {
		logrus.Fatalf("Cannot get columns from table: %s, err: %s", currentTable, err.Error())
	}

	err = db.Select(currentTable, selectAllFunc, col...)
	if err != nil {
		logrus.Fatalf("Cannot select from table: %s, err: %s", currentTable, err.Error())
	}

	return
}

func getTrackAlbums(db *sqlittle.DB, currentTable string) (list map[string]model.TrackAlbum) {
	list = make(map[string]model.TrackAlbum)
	var selectAllFunc = func(r sqlittle.Row) {
		var obj model.TrackAlbum
		err := r.Scan(
			&obj.AutoId,
			&obj.TrackId,
			&obj.AlbumId,
			&obj.TrackPosition,
			&obj.AlbumVolume,
		)
		if err != nil {
			logrus.Fatalf("Cannot scan rows from table: %s, err: %s", currentTable, err.Error())
		}

		list[obj.TrackId] = obj
	}

	col, err := db.Columns(currentTable)
	if err != nil {
		logrus.Fatalf("Cannot get columns from table: %s, err: %s", currentTable, err.Error())
	}

	err = db.Select(currentTable, selectAllFunc, col...)
	if err != nil {
		logrus.Fatalf("Cannot select from table: %s, err: %s", currentTable, err.Error())
	}

	return
}

func getTrackLyrics(db *sqlittle.DB, currentTable string) (list map[string]model.TrackLyrics) {
	list = make(map[string]model.TrackLyrics)
	var selectAllFunc = func(r sqlittle.Row) {
		var obj model.TrackLyrics
		err := r.Scan(
			&obj.TrackId,
			&obj.Lyrics,
			&obj.FullLyrics,
			&obj.Url,
			&obj.HasRights,
		)
		if err != nil {
			logrus.Fatalf("Cannot scan rows from table: %s, err: %s", currentTable, err.Error())
		}

		list[obj.TrackId] = obj
	}

	col, err := db.Columns(currentTable)
	if err != nil {
		logrus.Fatalf("Cannot get columns from table: %s, err: %s", currentTable, err.Error())
	}

	err = db.Select(currentTable, selectAllFunc, col...)
	if err != nil {
		logrus.Fatalf("Cannot select from table: %s, err: %s", currentTable, err.Error())
	}

	return
}

func getAlbums(db *sqlittle.DB, currentTable string) (list map[string]model.Album) {
	list = make(map[string]model.Album)
	var selectAllFunc = func(r sqlittle.Row) {
		var obj model.Album
		err := r.Scan(
			&obj.Id,
			&obj.Title,
			&obj.ArtistsString,
			&obj.AlbumVersion,
			&obj.Year,
			&obj.GenreId,
			&obj.GenreTitle,
			&obj.CoverUri,
			&obj.TrackCount,
			&obj.AlbumOptions,
		)
		if err != nil {
			logrus.Fatalf("Cannot scan rows from table: %s, err: %s", currentTable, err.Error())
		}

		list[obj.Id] = obj
	}

	col, err := db.Columns(currentTable)
	if err != nil {
		logrus.Fatalf("Cannot get columns from table: %s, err: %s", currentTable, err.Error())
	}

	err = db.Select(currentTable, selectAllFunc, col...)
	if err != nil {
		logrus.Fatalf("Cannot select from table: %s, err: %s", currentTable, err.Error())
	}

	return
}
