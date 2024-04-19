package session

import "sync"

type Storage interface {
	save(data *SignerSessionData)
	retrieve() (*SignerSessionData, error)
}

type MemorySessionStorage struct {
	data *SignerSessionData
	sync.Mutex
}

func NewMemorySessionStorage(data *SignerSessionData) *MemorySessionStorage {
	return &MemorySessionStorage{
		data: data,
	}
}

func (mss *MemorySessionStorage) save(data *SignerSessionData) {
	mss.Lock()
	defer mss.Unlock()

	mss.data = data
}

func (mss *MemorySessionStorage) retrieve() (*SignerSessionData, error) {
	if mss.data == nil {
		return nil, NoDataErr
	}
	return mss.data, nil
}
