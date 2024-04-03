package cmd

import (
	"fmt"
	"log"
	"path"
	"strings"

	common "github.com/commonjava/indy-tests/pkg/common"
	"github.com/spf13/cobra"
)

func NewExecuteCmd() *cobra.Command {

	exec := &cobra.Command{
		Use:   "execute $target $folo_track_id",
		Short: "upload and download artifacts in folo record ${folo_track_id}",
		Run: func(cmd *cobra.Command, args []string) {
			indyURL := args[0]
			foloTrackId := ""
			if len(args) >= 2 {
				foloTrackId = args[1]
			}
			if common.IsEmptyString(foloTrackId) {
				foloTrackId = DEFAULT_FOLO_TRACKING_ID
			}
			Exeute(indyURL, foloTrackId)
		},
	}

	return exec
}

func Exeute(originIndy, foloId string) {
	fileLoc := TEST_DIR + foloId + "-report.json"
	//dirLoc := TEST_DIR + foloId

	foloRecord := common.GetFoloRecordFromFile(fileLoc)
	for _, down := range foloRecord.Downloads {
		repoPath := strings.ReplaceAll(down.StoreKey, ":", "/")
		downUrl := fmt.Sprintf("%s%s", originIndy, path.Join("/api/content", repoPath, down.Path))
		log.Println("[Up/Download] " + downUrl)
	}
}
