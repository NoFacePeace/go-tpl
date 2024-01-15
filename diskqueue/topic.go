package diskqueue

import (
	"bytes"
	"encoding/binary"
	"os"
	"sort"
	"strconv"
	"sync"

	"github.com/pkg/errors"
)

type Topic struct {
	sync.Mutex
	writeFile   *os.File
	writePos    int64
	writeBuf    bytes.Buffer
	path        string
	files       []int
	name        string
	maxFileSize int64
}

func NewTopic(topic string) (*Topic, error) {
	t := &Topic{}
	t.name = topic
	t.path = topic
	if err := os.MkdirAll(t.name, os.ModePerm); err != nil {
		return nil, errors.Wrap(err, "mkdir all error")
	}
	dir, err := os.ReadDir(t.path)
	if err != nil {
		return nil, errors.Wrap(err, "read dir error")
	}
	t.files = []int{}
	for _, d := range dir {
		num, err := strconv.Atoi(d.Name())
		if err != nil {
			return nil, errors.Wrap(err, "str convert int error")
		}
		t.files = append(t.files, num)
	}
	sort.Ints(t.files)
	l := len(t.files)
	if l == 0 {
		file, err := os.OpenFile(t.path+"/"+"0", os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return nil, errors.Wrap(err, "open file error")
		}
		t.files = append(t.files, 0)
		t.writeFile = file
		t.writePos = 0
	} else {
		num := t.files[l-1]
		file, err := os.OpenFile(t.path+"/"+strconv.Itoa(num), os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return nil, errors.Wrap(err, "open file error")
		}
		t.writeFile = file
		fi, err := file.Stat()
		if err != nil {
			file.Close()
			return nil, errors.Wrap(err, "file stat error")
		}
		t.writePos = int64(fi.Size())
	}
	return t, nil
}

func (t *Topic) Write(msg []byte) (int, error) {
	t.Lock()
	defer t.Unlock()
	t.writeBuf.Reset()

	if err := binary.Write(&t.writeBuf, binary.BigEndian, int32(len(msg))); err != nil {
		return 0, errors.Wrap(err, "binary write error")
	}
	_, err := t.writeBuf.Write(msg)
	if err != nil {
		return 0, errors.Wrap(err, "writebuf write error")
	}
	_, err = t.writeFile.Write(t.writeBuf.Bytes())
	if err != nil {
		return 0, errors.Wrap(err, "write file error")
	}
	t.writePos += int64(8 + len(msg))
	if t.writePos > t.maxFileSize {
		t.writeFile.Sync()
		t.writeFile.Close()

	}
	return 0, nil
}

func (t *Topic) Read(group string, offset int) ([]byte, int, error) {
	return nil, 0, nil
}

func (t *Topic) openFile(name string) (*os.File, error) {
	file, err := os.OpenFile(t.path+"/"+"0", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, errors.Wrap(err, "open file error")
	}
	return file, nil
}

func (t *Topic) nextFile() string {
	return ""
}
