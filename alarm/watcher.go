package alarm

// the watcher takes an alarm and subscribes the next occurance with the manager
// it actually is the cronjob service behind the a single next occurance of an alarm
// and its 'execution' in the player context

type AlarmWatcher struct {
	Alarm *Alarm
}
