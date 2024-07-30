package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-webdav/webdav"
)

const (
	webDavURL   = "http://11.0.1.132:8000/webdav/user/admin"
	username    = "admin"
	password    = "admin123"
	watchDir    = "/home/admin/workspace/testwebdav" // 监听的本地目录
	bufferSize  = 10                                 // 文件事件缓冲区大小
	pollTimeout = 1 * time.Second                    // 轮询间隔
)

func uploadFile(filePath string, client *webdav.Client) error {
	// 将文件路径转换为WebDAV路径
	remotePath := filePath[len(watchDir):]

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// 上传文件
	err = client.PutFile(context.Background(), remotePath, file)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	return nil
}

func watchDirectory(client *webdav.Client) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("failed to create watcher: %v", err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					err := uploadFile(event.Name, client)
					if err != nil {
						log.Printf("Error uploading file: %v", err)
					} else {
						log.Printf("Uploaded: %s", event.Name)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("Error: %v", err)
			}
		}
	}()

	err = watcher.Add(watchDir)
	if err != nil {
		log.Fatalf("failed to add directory: %v", err)
	}

	<-done
}

func main() {
	log.Println("Starting WebDAV Client...")

	// 创建WebDAV客户端
	client, err := webdav.Dial(webDavURL, &webdav.Config{
		BasicAuth: &webdav.BasicAuth{
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer client.Close()

	// 开始监听目录
	watchDirectory(client)
}
