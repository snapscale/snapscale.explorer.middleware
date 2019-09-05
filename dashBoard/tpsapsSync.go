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
