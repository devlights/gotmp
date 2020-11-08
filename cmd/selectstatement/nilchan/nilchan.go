package main

import (
	"log"
	"os"
	"time"
)

var (
	tickLog = log.New(os.Stdout, "[tick  ] ", 0)
	dataLog = log.New(os.Stdout, "[dataCh] ", 0)
)

func main() {
	os.Exit(run())
}

func run() int {
	var (
		// データを流すチャネル
		//
		// キーボードよりエンター入力で活性状態になる
		dataCh chan int64
	)

	// <<非同期>>
	//
	// 終了を通知するチャネル
	theEnd := func() <-chan struct{} {
		ch := make(chan struct{})
		go func() {
			time.Sleep(10 * time.Second)
			close(ch)
		}()

		return ch
	}()

	// <<非同期>>
	//
	// キーボードから入力を読み取り dataCh の有効開始を通知するチャネル
	enableDataCh := func() <-chan struct{} {
		ch := make(chan struct{})
		go func() {
			os.Stdin.Read(make([]byte, 1))
			close(ch)
		}()

		return ch
	}()

	// <<非同期>>
	//
	// dataCh の有効通知が来たら、チャネルの実体を作り
	// インターバルでデータを投入する
	go func() {
		<-enableDataCh

		dataCh = make(chan int64)

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

	LOOP:
		for {
			select {
			case t := <-ticker.C:
				dataCh <- t.Local().Unix()
			case <-theEnd:
				break LOOP
			}
		}
	}()

	// <<メインゴルーチン>>
	// 
	// dataCh が有効になっていない間は time.Ticker を出力。
	// 有効になった後は dataCh の出力に切り替える。
	ticker := time.NewTicker(1 * time.Second)
	count := 0

LOOP:
	for {
		select {
		case <-enableDataCh:
			ticker.Stop()

			count++
			tickLog.Println(count, "stop")

			// 次の select で選択に入らないように nil を設定
			enableDataCh = nil
		case <-ticker.C:
			count++
			tickLog.Println(count)
		case x := <-dataCh:
			count++
			dataLog.Println(count, x)
		case <-theEnd:
			break LOOP
		}
	}

	return 0
}
