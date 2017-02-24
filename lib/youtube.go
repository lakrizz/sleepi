package lib

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

var tmpFolder = "/tmp"
var destFolder = "/home/krizz/piclocker/"

func YoutubeDLInstalled() bool {
	return true
}

func DownloadVideo(link string) (*os.File, error) {
	tmpfile, err := ioutil.TempFile(tmpFolder, "ytdl_")
	cmd := exec.Command("youtube-dl", link, "-x", "--audio-format", "mp3", "--audio-quality", "192K", "-o", tmpfile.Name()+".%(ext)s")
	_, err = cmd.CombinedOutput()
	//	fmt.Println(string(out))
	if err != nil {
		return nil, err
	}
	fl, err := os.Open(tmpfile.Name() + ".mp3")
	return fl, err

}

func GetTitle(url string) (string, error) {

	cmd := exec.Command("youtube-dl", url, "-e")
	t, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(t), nil

}

func CpFile(fl *os.File, title string) (string, error) {
	destfname := path.Join(destFolder, title+".mp3")
	destFile, err := os.Create(destfname)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(destFile, fl)

	return destfname, err
}
