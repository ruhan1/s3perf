package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	common "github.com/commonjava/indy-tests/pkg/common"
	"github.com/spf13/cobra"
)

func NewPrepareCmd() *cobra.Command {
	exec := &cobra.Command{
		Use:   "prepare $indy_url $folo_track_id",
		Short: "Download artifacts in folo record to local directory './test/${folo_track_id}'",
		Run: func(cmd *cobra.Command, args []string) {
			indyURL := args[0]
			foloTrackId := ""
			if len(args) >= 2 {
				foloTrackId = args[1]
			}
			if common.IsEmptyString(foloTrackId) {
				foloTrackId = DEFAULT_FOLO_TRACKING_ID
			}
			Run(indyURL, foloTrackId)
		},
	}
	return exec
}

const (
	TEST_DIR                 = "test/"
	DEFAULT_FOLO_TRACKING_ID = "build-A6RE4WO5CDYAA"
)

func Run(originIndy, foloId string) {
	foloRecord := GetFoloRecord(originIndy, foloId)

	dirLoc := TEST_DIR + foloId
	os.MkdirAll(dirLoc, os.FileMode(0755))

	fmt.Println("Start preparing artifacts.")
	fmt.Printf("==========================================\n\n")
	broken := false
	for index, down := range foloRecord.Downloads {
		repoPath := strings.ReplaceAll(down.StoreKey, ":", "/")
		downloadUrl := fmt.Sprintf("%s%s", originIndy, path.Join("/api/content", repoPath, down.Path))
		broken = DownloadFunc(dirLoc, down.Md5, downloadUrl, down.Path)
		if broken {
			fmt.Printf("Preparing artifacts failed (done: %d).\n\n", index)
			break
		}
	}
	if broken {
		os.Exit(1)
	} else {
		fmt.Printf("Preparing artifacts finished.\n\n")
	}
}

func GetFoloRecord(originIndy, foloId string) common.TrackedContent {
	fileLoc := TEST_DIR + foloId + "-report.json"
	if !common.FileOrDirExists(fileLoc) {
		data := common.GetFoloRecordAsString(originIndy, foloId)
		err := ioutil.WriteFile(fileLoc, []byte(data), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	return common.GetFoloRecordFromFile(fileLoc)
}

func DownloadFunc(dirLoc, md5str, downloadURL, localPath string) bool {
	fileLoc := path.Join(dirLoc, localPath)
	if common.FileOrDirExists(fileLoc) {
		return true // already exists
	}
	success, _ := common.DownloadFile(downloadURL, fileLoc)
	if success {
		return common.Md5Check(fileLoc, md5str)
	}
	return false
}
