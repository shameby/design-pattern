package main

import "fmt"

func main() {
	hd := NewHttpDownloader()
	fd := NewFtpDownloader()

	hd.Download("www.baidu.com")
	fd.Download("www.google.com")
}

func NewHttpDownloader() Downloader {
	downloader := &HttpDownloader{}
	temp := newTemplate(downloader)
	downloader.template = temp
	return downloader
}

func NewFtpDownloader() Downloader {
	downloader := &FtpDownloader{}
	tmp := newTemplate(downloader)
	downloader.template = tmp
	return downloader
}

type Downloader interface {
	Download(uri string)
}

type implement interface {
	download()
	save()
}

type template struct {
	implement
	uri string
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Println("prepare downloading")
	t.implement.download()
	t.implement.save()
	fmt.Println("finish downloading")
}

func (t *template) save() {
	fmt.Println("default save")
}

type HttpDownloader struct {
	*template
}

func (d *HttpDownloader) download() {
	fmt.Printf("download %s by http \n", d.uri)
}

func (*HttpDownloader) save() {
	fmt.Println("http save")
}

type FtpDownloader struct {
	*template
}

func (d *FtpDownloader) download() {
	fmt.Printf("download %s by ftp \n", d.uri)
}