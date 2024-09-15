package util

import (
	"math"
	"math/rand/v2"
	"time"
)

var (
	rng          = rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 114514))
	mean float64 = 5.3
	// 标准差
	stdDev float64 = 3.0
)

func generateScore() float64 {
	// 生成标准正态分布的随机数
	standardNormal := rng.NormFloat64()
	// 调整为指定均值和标准差的正态分布
	score := mean + stdDev*standardNormal

	return score
}

func GenerateScore() int64 {
	score := generateScore()
	//for score < 0.0 {
	//	score = generateScore()
	//}
	return int64(math.Round(score))
}
