package dashBoard

import (
	"snapscale-api/apiClient/chain"
)

func shortSync() {
	syncInfo()
}

func syncInfo() {
	t, e1, e2 := chain.GetInfo()
	if e1 == nil && e2 == nil && t != nil {
		DataCenter.Blocks.CurrentBlock = t.HeadBlockNum
		DataCenter.Blocks.LastIrreversibleBlock = t.LastIrreversibleBlockNum
		DataCenter.Producers.CurrentProducer = t.HeadBlockProducer
		DataCenter.IO.VirtualBlockCpuLimit = t.VirtualBlockCpuLimit
		DataCenter.IO.VirtualBlockNetLimit = t.VirtualBlockNetLimit
		DataCenter.IO.BlockCpuLimit = t.BlockCpuLimit
		DataCenter.IO.BlockNetLimit = t.BlockNetLimit
		BroadcastDashboard()

		if SynchronizedBlockId == 0 {
			SynchronizedBlockId = DataCenter.Blocks.CurrentBlock
		}

		if DataCenter.Blocks.CurrentBlock > SynchronizedBlockId {
			SyncFinish <- true
		}
	}
}
