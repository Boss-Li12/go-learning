package crawler

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"strconv"
	"strings"
)

var txHashBuf map[string]string

func GetBlockHashByBlockId(blockId int) string {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	var blockHash string
	c.OnHTML(".body .container-1280", func(e *colly.HTMLElement) {
		blockHash = e.ChildText("div.br-8 > div.page-block__hash-content > h2.mt-5")
	})
	c.Visit(fmt.Sprintf("https://blockchair.com/bitcoin/block/%d", blockId))
	return strings.TrimSpace(blockHash)
}

func GetTransactionHashByBlockHash(blockHash string) []string {
	c := colly.NewCollector(
		colly.MaxDepth(2),
	)
	totalTxHashs := make([]string, 0)
	first := true
	c.OnHTML("div.main-container > div:nth-of-type(2)", func(e *colly.HTMLElement) {
		txHashs := e.ChildTexts("div.well > h4.truncate > a[href]")
		for _, txHash := range txHashs {
			totalTxHashs = append(totalTxHashs, txHash)
			fmt.Println(txHash)
		}
		if first {
			first = false
			page := e.ChildText("div:nth-of-type(3) div.text-center")
			part := strings.Split(page, "of")
			totalPageStr := strings.TrimSpace(part[len(part) - 1])
			totalPage, _ := strconv.Atoi(totalPageStr)
			for i := 2; i <= totalPage; i++ {
				c.Visit(fmt.Sprintf("https://live.blockcypher.com/btc/block/%s?page=%d", blockHash, i))
			}
		}
 	})
	c.OnError(func(r *colly.Response, err error) {
		log.Fatalf("fail to crawl: %v", err)
	})
	c.Visit(fmt.Sprintf("https://live.blockcypher.com/btc/block/%s", blockHash))
	return totalTxHashs
}

func GetTransactionHashByBlockId(blockId int) []string {
	return GetTransactionHashByBlockHash(GetBlockHashByBlockId(blockId))
}
