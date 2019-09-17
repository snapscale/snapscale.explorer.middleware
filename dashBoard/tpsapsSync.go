package dashBoard

type tpsapsS struct {
	Tps int64
	Aps int64
}

var tpsaps chan *tpsapsS

var t1 int64
var a1 int64

func tpsapsSync() {
	for {
		select {
		case one := <-tpsaps:
			DataCenter.Performance.Aps = a1 + one.Aps
			DataCenter.Performance.Tps = t1 + one.Tps

			if len(DataCenter.Performance.TpsA) < 10 {
				DataCenter.Performance.TpsA = append(DataCenter.Performance.TpsA, DataCenter.Performance.Tps)
			} else {
				tmpA := append(DataCenter.Performance.TpsA, DataCenter.Performance.Tps)
				DataCenter.Performance.TpsA = tmpA[1:]
			}

			if len(DataCenter.Performance.ApsA) < 10 {
				DataCenter.Performance.ApsA = append(DataCenter.Performance.ApsA, DataCenter.Performance.Aps)
			} else {
				tmpB := append(DataCenter.Performance.ApsA, DataCenter.Performance.Aps)
				DataCenter.Performance.ApsA = tmpB[1:]
			}

			if DataCenter.Performance.Aps > DataCenter.Performance.ApsHigh {
				DataCenter.Performance.ApsHigh = DataCenter.Performance.Aps
			}
			if DataCenter.Performance.Tps > DataCenter.Performance.TpsHigh {
				DataCenter.Performance.TpsHigh = DataCenter.Performance.Tps
			}

			t1 = one.Tps
			a1 = one.Aps
		}
	}
}
