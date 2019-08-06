package youtube

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"../../models"
	"github.com/kennygrant/sanitize"
)

var targetfolder string = "./downloads/"

type downloader struct {
	messages chan *models.Message
}

func (d *downloader) Download(video *models.Video) error {
	file, err := d.getpath(video)
	if err != nil {
		return err
	}

	return d.download(video, file)
}

func (d *downloader) getpath(video *models.Video) (string, error) {
	if _, e := os.Stat(targetfolder); os.IsNotExist(e) {
		// create folder, right?
		os.Mkdir(targetfolder, os.ModeDir|0777)
	}

	// now check if the file probably already exists?
	targetfile := fmt.Sprintf("./%s.mp4", path.Join(targetfolder, sanitize.Path(video.Title)))

	return targetfile, nil
}

func (d *downloader) download(video *models.Video, file string) error {
	// spawn the process, i fucking hate relying on third party programs, but at least they fucking work...
	d.messages <- models.CreateMessage(video.Title, models.MSG_DOWNLOAD_STARTING)
	cmd := exec.Command("youtube-dl", "-x", "-f", "140", "-o", file, video.GetUrl())
	_, err := cmd.CombinedOutput()
	if err != nil {
		d.messages <- models.CreateMessage(video.Title, models.MSG_DOWNLOAD_ERROR)
		return err
	}
	d.messages <- models.CreateMessage(video.Title, models.MSG_DOWNLOAD_FINISHED)
	return nil
}
