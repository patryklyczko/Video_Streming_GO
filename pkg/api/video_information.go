package api

import "github.com/patryklyczko/Video_Streming_GO/pkg/db"

func (a *InstanceAPI) VideoInfo(name string) (*db.VideoInformation, error) {
	return a.dbController.VideoInfo(name)
}
