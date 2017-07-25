package manager

type Manager struct {
	// libs
	downloader *lib.Downloader
	plmanager  *lib.PlaylistManager

	// comm
	comm chan *Message
}

type Message struct {
	Status   int
	Receiver int
	Payload  []byte
}

const (
	STATUS_ENQUEUED = iota
	STATUS_DOWNLOADING
	STATUS_DOWNLOAD_COMPLETE
	STATUS_CONVERTING
	STATUS_CONVERTED
	STATUS_PROCESSING
	STATUS_AVAILABLE
)

var Manager *Manager

func Init() error {
	if Manager == nil {
		comchan := make(chan *Message, 10) // 10 simultaneous tasks? TODO: revise?
		Manager = &Manager{comm: comchan}
	}

	Manager.downloader = &lib.Downloader{comm}

	go Manager.read()
	return nil
}

func (m *Manager) read() {
	for {
		select {
		case v := <-m.comm:
			switch v.Status {
			case STATUS_ENQUEUED:
				// this means we have the payload enqueued (TODO(@krizz)maybe we should create a struct "job"?)
				// first download the shit \o\
				m.downloader.Download(v.Payload)
				break
			case STATUS_DOWNLOAD_COMPLETE:
				// who is responsible of the copying?
				break

			case STATUS_CONVERTING:
				m.plmanager.Handle(v.Payload)
				break

			default:
				break

			}

		default:
			break
			// anyway: do shit, someone wants to get notified
			// websocket.Send(v)

		}
	}
}
