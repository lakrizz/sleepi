package youtube

import (
	"fmt"
	"log"
	"os"
	"path"

	"../../models"
	"github.com/k0kubun/pp"
	"github.com/kennygrant/sanitize"
	"github.com/rylio/ytdl"
)

var targetfolder string = "./downloads/"

type downloader struct {
}

func (d *downloader) Download(video *models.Video) error {
	info, err := ytdl.GetVideoInfoFromID(video.Id)
	if err != nil {
		return err
	}

	file, err := d.getpath(info)
	if err != nil {
		return err
	}

	pp.Println(info)

	defer file.Close()
	info.
	return info.Download(ytdl.FORMATS[140], file)

}

func (d *downloader) getpath(info *ytdl.VideoInfo) (*os.File, error) {
	if _, e := os.Stat(targetfolder); os.IsNotExist(e) {
		// create folder, right?
		os.Mkdir(targetfolder, os.ModeDir|0777)
	}

	// now check if the file probably already exists?
	targetfile := fmt.Sprintf("./%s.mp4", path.Join(targetfolder, sanitize.Path(info.Title)))
	if _, e := os.Stat(targetfile); !os.IsNotExist(e) {
		return nil, e
	}

	log.Println(targetfile)

	f, e := os.Create(targetfile)
	if e != nil {
		return nil, e
	}

	e = f.Chmod(0x777)
	if e != nil {
		return nil, e
	}

	return f, nil
}
