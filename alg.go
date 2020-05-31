package main

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

// PickWinners
//  算法描述：
//  1. 选取指定时刻（即抽奖时间）后被挖出的第一个比特币区块的哈希值作为随机数种子，记作 S。
//  2. 用 SHA-256 算法计算 S 的的哈希值 H，然后把 H 作为16进制数字转换为长整数 L。
//  3. W = L % N 为中奖的奖券编号，其中 N 为总奖券数量，%为求余数。
//  4. 如需抽出 M 个中奖者，则设新种子为 S = H 并且重复 2、3 两步，直到抽出 M 个不重复的中奖者为止。
//  函数描述：
//  min 奖券编号中最小的（通常我们会把它设成0）
//  max 奖券编号中最大的（取决于参与抽奖的用户数）
//  num 中奖人数
//  key 指定的抽奖时间后被挖出的第一个比特币区块的哈希值
func PickWinners(min, max, num int, key string) []int {
	winners := make(map[int]bool)
	for res := key; len(winners) < num; {
		h := sha256.New()
		h.Write([]byte(res))
		res = hex.EncodeToString(h.Sum(nil))

		d := new(big.Int)
		d.SetString(hex.EncodeToString(h.Sum(nil)), 16)
		d = d.Mod(d, new(big.Int).SetInt64(int64(max-min+1)))

		winners[int(d.Int64())+min] = true
	}

	var ret []int
	for i := range winners {
		ret = append(ret, i)
	}

	return ret
}
