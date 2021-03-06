# lottery-alg

![Alt text](https://github.com/BetaCat0/lottery-alg/raw/master/dice.jpg)

朋友圈里六一小礼物抽奖算法的实现 :)

#### 算法描述：
在这里假设奖券编号是连续发放的整数。抽奖算法描述如下：

1. 选取指定时刻（即抽奖时间）后被挖出的第一个比特币区块的哈希值作为随机数种子，记作 S。
2. 用 SHA-256 算法计算 S 的的哈希值 H，然后把 H 作为16进制数字转换为长整数 L。
3. W = L % N 为中奖的奖券编号，其中 N 为总奖券数量，%为求余数。
4. 如需抽出 M 个中奖者，则设新种子为 S = H 并且重复 2、3 两步，直到抽出 M 个不重复的中奖者为止。

上述抽奖步骤实际上是用完全公开可验证的方法生成了一个或多个不可控的随机数，其中最重要的随机数种子由比特币区块的性质来保证它满足所有的要求。
只要知道了公布的抽奖时间和发放的奖券总数，任何人都可以在奖券停止发放后计算出一样的伪随机数，从而实现了可验证的公平抽奖结果。

参考文章：<a href="https://www.uscreditcardguide.com/how-do-we-make-our-giveaway-fair/">《我们是如何用科学的方法保证抽奖的公平性的》</a>
