package log

type LamportClock struct {
	ID   string
	Time int64
}

func (l LamportClock) Tick() LamportClock {
	return LamportClock{
		ID:   l.ID,
		Time: l.Time + 1,
	}
}

func (l LamportClock) Merge(clock LamportClock) LamportClock {
	if l.Time < clock.Time {
		l.Time = clock.Time
	}

	return LamportClock{
		ID:   l.ID,
		Time: l.Time,
	}
}

func (l LamportClock) Clone() LamportClock {
	return LamportClock{
		ID:   l.ID,
		Time: l.Time,
	}
}

func CompareClocks(a, b LamportClock) int64 {
	dist := a.Time - b.Time

	if dist == 0 && a.ID != b.ID {
		if a.ID < b.ID {
			return -1
		}
		return 1
	}

	return dist
}
