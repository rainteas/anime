package dto

type TorrentStatus int

const (
	Stopped TorrentStatus = iota
	CheckPending
	Checking
	DownloadPending
	Downloading
	SeedPending
	Seeding
)

func (ts TorrentStatus) String() string {
	switch ts {
	case Stopped:
		return "Stopped"
	case CheckPending:
		return "Check Pending"
	case Checking:
		return "Checking"
	case DownloadPending:
		return "Download Pending"
	case Downloading:
		return "Downloading"
	case SeedPending:
		return "Seed Pending"
	case Seeding:
		return "Seeding"
	default:
		return "Unknown"
	}
}
