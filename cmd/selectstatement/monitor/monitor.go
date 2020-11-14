package main

import (
	"log"
	"os"
	"time"
)

func main() {
	var (
		appLog    = log.New(os.Stdout, "[main] ", 0)
		g1Log     = log.New(os.Stdout, "[G1  ] ", 0)
		g2Log     = log.New(os.Stdout, "[G2  ] ", 0)
		closerLog = log.New(os.Stdout, "[closer] ", 0)
	)

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
	startGoroutine(begin, g1End, g1Log)
	startGoroutine(g1End, g2End, g2Log)

	// 終わりを検知するためのゴルーチンを起動
	startGoroutine(g2End, end, closerLog)

	// 現在状況をお知らせするゴルーチンを起動
	status := startMonitor(begin, g1End, g2End, end)

	// 処理の開始を告げる
	go func(begin chan<- struct{}) {
		time.Sleep(3 * time.Second)
		close(begin)
	}(begin)

	// ------------------------------------------------
	// メインゴルーチン
	// ------------------------------------------------
	// 経過出力
	for s := range status {
		appLog.Println(s, time.Now().Unix())
	}

	appLog.Println("done")
}

func startGoroutine(begin <-chan struct{}, end chan<- struct{}, l *log.Logger) {
	go func(begin <-chan struct{}, end chan<- struct{}, l *log.Logger) {
		<-begin

		l.Println("started", time.Now().Unix())
		time.Sleep(3 * time.Second)
		close(end)
		l.Println("end", time.Now().Unix())
	}(begin, end, l)
}

func startMonitor(begin, g1, g2, end <-chan struct{}) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

	LOOP:
		for {
			select {
			case <-ticker.C:
				switch {
				case !isDone(begin):
					out <- "start pending"
					continue LOOP
				case !isDone(g1):
					out <- "g1 running"
					continue LOOP
				case !isDone(g2):
					out <- "g2 running"
					continue LOOP
				case !isDone(end):
					out <- "closer running"
					continue LOOP
				default:
					break LOOP
				}
			}
		}
	}()
	return out
}

func isDone(ch <-chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
