package main

import (
	"log"
	"os"
	"time"
)

var (
	appLog = log.New(os.Stderr, "[main] ", 0)
	g1Log  = log.New(os.Stderr, "[G1  ] ", 0)
	g2Log  = log.New(os.Stderr, "[G2  ] ", 0)
)

func main() {
	// 2つのゴルーチンを順に実行していき
	// その都度メインゴルーチンで状況を出力

	// 始まりと終わりを検知するためのチャネル
	var (
		begin = make(chan struct{})
		end   = make(chan struct{})
	)

	// 2つのゴルーチンの完了検知のためのチャネル
	var (
		g1End = make(chan struct{})
		g2End = make(chan struct{})
	)

	// 処理対象のゴルーチンを２つ起動
	go func(begin <-chan struct{}, end chan<- struct{}) {
		<-begin

		g1Log.Println("started", time.Now().Unix())
		time.Sleep(3 * time.Second)
		close(end)
		g1Log.Println("end", time.Now().Unix())
	}(begin, g1End)

	go func(begin <-chan struct{}, end chan<- struct{}) {
		<-begin

		g2Log.Println("started", time.Now().Unix())
		time.Sleep(3 * time.Second)
		close(end)
		g2Log.Println("end", time.Now().Unix())
	}(g1End, g2End)

	// 終わりを検知するためのゴルーチンを起動
	go func(begin <-chan struct{}, end chan<- struct{}) {
		<-begin

		time.Sleep(3 * time.Second)
		close(end)
	}(g2End, end)

	// 現在状況をお知らせするゴルーチンを起動
	status := func(begin, g1, g2, end <-chan struct{}) <-chan string {
		out := make(chan string)
		go func() {
			defer close(out)

			ticker := time.NewTicker(1 * time.Second)
			defer ticker.Stop()

		LOOP:
			for {
				select {
				case <-ticker.C:

					select {
					case <-begin:
					default:
						out <- "start pending"
						continue LOOP
					}

					select {
					case <-g1End:
					default:
						out <- "g1 running"
						continue LOOP
					}

					select {
					case <-g2End:
					default:
						out <- "g2 running"
						continue LOOP
					}

					select {
					case <-end:
						out <- "done"
						break LOOP
					default:
						out <- "g1, g2 ended, wait finishing"
						continue LOOP
					}
				}
			}
		}()
		return out
	}(begin, g1End, g2End, end)

	// ------------------------------------------------
	// メインゴルーチン
	// ------------------------------------------------
	// 処理の開始を告げる
	go func(begin chan<- struct{}) {
		time.Sleep(3 * time.Second)
		close(begin)
	}(begin)

	// 経過出力
	for s := range status {
		appLog.Println(s, time.Now().Unix())
	}
}
