package cmd

import (
	"fmt"
	"os"
	"path"

	common "github.com/commonjava/indy-tests/pkg/common"
	"github.com/spf13/cobra"
)

func NewExecuteCmd() *cobra.Command {
	exec := &cobra.Command{
		Use:   "execute $target $folo_track_id",
		Short: "Upload and download artifacts in folo record ${folo_track_id}",
		Run: func(cmd *cobra.Command, args []string) {
			target := args[0]
			foloTrackId := ""
			if len(args) >= 2 {
				foloTrackId = args[1]
			}
			if common.IsEmptyString(foloTrackId) {
				foloTrackId = DEFAULT_FOLO_TRACKING_ID
			}
			Exeute(target, foloTrackId)
		},
	}

	return exec
}

func Exeute(target, foloId string) {
	fileLoc := TEST_DIR + foloId + "-report.json"
	dirLoc := TEST_DIR + foloId

	foloRecord := common.GetFoloRecordFromFile(fileLoc)
	testEntries := foloRecord.Downloads[:1] // test 1 file

	fmt.Printf("\nStart uploading artifacts.\n")
	fmt.Println("==========================================")
	broken := false
	for index, entry := range testEntries {
		storageUrl := fmt.Sprintf("%s%s", target, path.Join("/api/storage/content", entry.StoreKey, entry.Path))
		localFile := path.Join(dirLoc, entry.Path)
		broken = !common.UploadFile(storageUrl, localFile)
		if broken {
			fmt.Printf("Uploading artifacts failed (done: %d).\n\n", index)
			break
		}
	}
	if broken {
		os.Exit(1)
	} else {
		fmt.Printf("Uploading artifacts finished.\n\n")
	}

	fmt.Printf("\nStart downloading artifacts.\n")
	fmt.Println("==========================================")
	tmpDir := "/tmp"
	for index, entry := range testEntries {
		storageUrl := fmt.Sprintf("%s%s", target, path.Join("/api/storage/content", entry.StoreKey, entry.Path))
		broken = !DownloadFunc(tmpDir, entry.Md5, storageUrl, entry.Path)
		if broken {
			fmt.Printf("Downloading artifacts failed (done: %d).\n\n", index)
			break
		}
	}
	if broken {
		os.Exit(1)
	} else {
		fmt.Printf("Uploading artifacts finished.\n\n")
	}
}
