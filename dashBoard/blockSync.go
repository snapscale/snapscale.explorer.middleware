package dashBoard

import (
	"fmt"
	"snapscale-api/apiClient/chain"
)

var SynchronizedBlockId int64
var SyncNow chan int64
var SyncFinish chan bool

func blockSync() {
	for {
		select {
		case id := <-SyncNow:
			t, e1, e2 := chain.GetBlock(fmt.Sprintf("%d", id))
			if e1 == nil && e2 == nil && t != nil {
				var result = new(tpsapsS)
				result.Tps = int64(len(t.Transactions))
				for _, one := range t.Transactions {
					result.Aps += int64(len(one.Trx.Transaction.Actions))
				}
				tpsaps <- result
				BroadcastBlock(t)
				SynchronizedBlockId += 1
				go syncFinish()
			}
		case <-SyncFinish:
			if DataCenter.Blocks.CurrentBlock > SynchronizedBlockId {
				go nextBlock()
			}
		}
	}
}

func nextBlock() {
	SyncNow <- SynchronizedBlockId + 1
}

func syncFinish() {
	SyncFinish <- true
}
