package verinfo

import (
	"fmt"
	"runtime/debug"
	"time"
)

// CommitInfo contains information about the commit and tag a build was created with.
type CommitInfo struct {
	Version    string
	Revision   string
	LastCommit time.Time
	DirtyBuild bool
}

// Get fetches version and git commit data embedded in the current build.
func Get() (CommitInfo, error) {
	var commitInfo CommitInfo

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return CommitInfo{}, fmt.Errorf("no commit data found")
	}

	commitInfo.Version = info.Main.Version

	for _, kv := range info.Settings {
		switch kv.Key {
		case "vcs.revision":
			commitInfo.Revision = kv.Value
		case "vcs.time":
			commitInfo.LastCommit, _ = time.Parse(time.RFC3339, kv.Value)
		case "vcs.modified":
			commitInfo.DirtyBuild = kv.Value == "true"
		}
	}

	return commitInfo, nil
}
