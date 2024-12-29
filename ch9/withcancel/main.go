package main

import (
	"context"
	"crypto/sha256"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"
)

func main() {
	targetBits, _ := strconv.Atoi(os.Args[1])
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// ctx, cancel = context.WithTimeout(ctx, 1*time.Second)
	// defer cancel()

	resultCh := make(chan string, 1)

	go pow(ctx, targetBits, resultCh)

	select {
	case <-time.After(2 * time.Second):
		log.Println("找不到比目標值小的數")
		return
	case result := <-resultCh:
		log.Println("找到一個比目標值小的數", result)
		return
	}
}

func pow(ctx context.Context, targetBits int, result chan string) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	log.Println("開始尋找一個數, 使得 hash 小於 target")
	for {
		select {
		case <-ctx.Done():
			log.Println("pow:", "done")
			return
		default:
			data := "hello world " + strconv.Itoa(nonce)
			hash = sha256.Sum256([]byte(data)) // 计算hash值
			hashInt.SetBytes(hash[:])          // 将hash值转换为big.Int

			if hashInt.Cmp(target) < 1 { // hashInt < target, 找到一个超过目标数值的数，也就是至少前targetBits位为0
				result <- data
				return
			} else { // 没找到，继续找
				nonce++
			}
		}
	}
}
