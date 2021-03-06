package recorder

import (
	"github.com/luxcgo/lifesaver/dispatcher"
	"github.com/luxcgo/lifesaver/engine"
	"github.com/luxcgo/lifesaver/enum"
	l "github.com/luxcgo/lifesaver/log"
	"github.com/sirupsen/logrus"
)

var RecorderManager = NewManager()

func init() {
	dispatcher.SharedManager.Register(
		RecorderManager,
	)
}

func (m *Manager) Dispatch(event *dispatcher.Event) error {
	show := event.Object.(engine.Show)

	l.Logger.WithFields(logrus.Fields{
		"pf": show.GetPlatform(),
		"id": show.GetRoomID(),
	}).Info("dispatch ", event.Type)

	switch event.Type {
	case enum.EventType.AddRecorder:
		return m.addRecorder(show)
	case enum.EventType.RemoveRecorder:
		return m.removeRecorder(show)
	}
	return nil
}

func (m *Manager) DispatcherType() enum.DispatcherTypeID {
	return enum.DispatcherType.Recorder
}

func (m *Manager) DispatchTypes() []enum.EventTypeID {
	return []enum.EventTypeID{
		enum.EventType.AddRecorder,
		enum.EventType.RemoveRecorder,
	}
}
