package diskqueue

import "sync"

type FileStorage struct {
	topicMap sync.Map
	sync.Mutex
}

func (f *FileStorage) Write(topic string, msg []byte) (int, error) {
	t, err := f.GetTopic(topic)
	if err != nil {
		return 0, err
	}
	return t.Write(msg)
}

func (f *FileStorage) Read(topic string, group string, offset int) ([]byte, int, error) {
	t, err := f.GetTopic(topic)
	if err != nil {
		return nil, 0, err
	}
	return t.Read(group, offset)
}

func (f *FileStorage) GetTopic(topic string) (*Topic, error) {
	if t, ok := f.topicMap.Load(topic); ok {
		return t.(*Topic), nil
	}
	f.Lock()
	defer f.Unlock()
	if t, ok := f.topicMap.Load(topic); ok {
		return t.(*Topic), nil
	}
	t, err := NewTopic(topic)
	if err != nil {
		return nil, err
	}
	f.topicMap.Store(topic, t)
	return t, nil
}
