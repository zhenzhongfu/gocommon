package status

import (
	"sync/atomic"

	"github.com/zhenzhongfu/gocommon/logging"
)

var (
	// 处理的数据包数量
	RecvLen uint64
	SendLen uint64

	// 处理的数据包容量
	RecvNum uint64
	SendNum uint64

	// 平均单个数据包容量
	RecvAvgLen uint64
	SendAvgLen uint64

	// 处理的数据包数量
	DispatchNum uint64

	// 处理包的总时间
	DispatchTime uint64

	// 处理单个包的平均时间
	DispatchAvgTime uint64
)

func RecvOne(len uint64) {
	l := atomic.AddUint64(&RecvLen, len)
	n := atomic.AddUint64(&RecvNum, 1)
	atomic.StoreUint64(&RecvAvgLen, l/n)
}

func SendOne(len uint64) {
	l := atomic.AddUint64(&SendLen, len)
	n := atomic.AddUint64(&SendNum, 1)
	atomic.StoreUint64(&SendAvgLen, l/n)
}

func DispatchOne(deltaTime uint64) {
	n := atomic.AddUint64(&DispatchNum, 1)
	t := atomic.AddUint64(&DispatchTime, deltaTime)
	atomic.StoreUint64(&DispatchAvgTime, t/(n*1000)) //us
}

func Show() {
	logging.Warnln("STAT====\n",
		"recv len:", atomic.LoadUint64(&RecvLen), "\n",
		"send len:", atomic.LoadUint64(&SendLen), "\n",
		"recv num:", atomic.LoadUint64(&RecvNum), "\n",
		"send num:", atomic.LoadUint64(&SendNum), "\n",
		"RecvAvg len:", atomic.LoadUint64(&RecvAvgLen), "\n",
		"SendAvg len:", atomic.LoadUint64(&SendAvgLen), "\n",
		"DispatchNum:", atomic.LoadUint64(&DispatchNum), "\n",
		"DispatchTime:", atomic.LoadUint64(&DispatchTime), "us\n",
		"DispatchAvgTime:", atomic.LoadUint64(&DispatchAvgTime), "us\n",
		"====\n")
}
