package filepath

import (
	"log"

	"github.com/howeyc/fsnotify"
)

// HandleFunc 用于自定义的文件事件处理函数
type HandleFunc func(event *fsnotify.FileEvent)

// Watch 监视文件系统中的文件变化并作出相应的处理。正常情况下，该函数会阻塞，直到从通道done中读取到一个元素未知。
func Watch(logger *log.Logger, done <-chan struct{}, path string, fn HandleFunc) (err error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Println(err)
		return err
	}

	go func() {
		for {
			select {
			case event := <-watcher.Event:
				fn(event)
			case err := <-watcher.Error:
				logger.Printf("watch event err: %s\n", err)
			}
		}
	}()

	err = watcher.Watch(path)
	if err != nil {
		logger.Println(err)
		return err
	}

	logger.Printf("Start watch path: %s\n", path)
	<-done
	logger.Printf("Finish watch path: %s\n", path)

	if err = watcher.Close(); err != nil {
		logger.Println(err)
		return err
	}
	return nil
}
