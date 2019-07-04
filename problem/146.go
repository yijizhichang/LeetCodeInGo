/**
 * Created by GoLand.
 * User: zghua
 * Email: xzghua@gmail.com
 * Date: 2019-07-04
 * Time: 11:11
 */
package problem

//
//### 题目:
//> 运用你所掌握的数据结构，设计和实现一个`LRU (最近最少使用) 缓存机制`。它应该支持以下操作： 获取数据 `get` 和 写入数据 `put` 。
//>
//> 获取数据 `get(key)`
//> - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//>
//> 写入数据 `put(key, value)`
//> - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
//>
//>
//> 你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
//
//### 示例:
//
//```
//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4
//
//```
//
//### 解析题目:
//
//- 首先得了解一下LRU,LRU举例来说,就是一个固定盒子只能放n(n固定)本书,每放一本书,后放的肯定放在前面一本的上面,然后如果想要拿书,如果里面有,就把书拿出来,放在底下所有书的最上面!当书本数超出n,就把最底下那本书给拿出去,最后放的依然放在最上面!
//- 题目示例里,`key`是等于`value`的,实际情况并不是!`key`跟`value`只是对应关系,并无对等关系!
//- 常见的LRU例子都是对`value`进行整理,但此题得注意,常操作的是是put里的`key`,`value`只是根据`key`返回的结果!所以,LRU里操作的是本题里的`key`!
//- 超出得去掉不常用的那个数
//- get操作不仅仅是为了获得值,当获取到值了以后,那么这个值就会被放到最上面,但当前整体数量不会发生改变
//- 进阶里的`O(1)` 那就别循环了,循环开始就 `O(n)`了
//- 如果相同`key`对应不同的`value`,会覆盖以前的value,而并不是两个值,比如高度为3,先放一个 `put(1,1)`,在放一个`put(2,4)`,再放一个`put(1,2)`,这时候,链表里应该是两个值,分别是(高到低) `[1,2]` ,而不是`[1,2,1]`
//
//### 思路分析:
//
//- 需要知道最上面是谁,也需要知道最下面是谁,一旦变动,可能要把自己放到最上面,然后把自己上面的值和自己下面的值进行相连,所以,最好的方式就是用个双向链表的方式来表示,既知道自己上面是谁,也知道自己下面是谁!
//- 因为每次`新放一个数`或者`获得到一个数`的时候,都会把这个数放到最顶上,那么可以设计一个 `Head` 代表 链表的头节点,
//- 因为可能会超出长度,超出的要删掉最后一个,而且所有的数都是一个一个的添加进去的,所以可以保存最后一个节点,方便删除最后一个节点! 可以设计一个`Last` 代表链表的最后一个节点
//- 因为数是一个一个添加进去的,所以得知道`总的高度`是多少,`当前高度`是多少,可以设计一个 `Lens`代表总高度,`NowLens`代表当前高度
//

type ListNode struct {
	Val int
	Key int
	Next *ListNode
	Pre *ListNode
}

type LRUCache struct {
	MapKeyValueAddr map[int]*ListNode
	Lens 		int
	NowLens 	int
	Head 		*ListNode
	Last 		*ListNode
}


func Constructor(capacity int) LRUCache {
	r := make(map[int]*ListNode,capacity)
	lru := LRUCache{
		MapKeyValueAddr: r,
		Lens: capacity,
		NowLens: 0,
		Head: nil,
		Last: nil,
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	// 如果当前的高度为0,说明一个数也没有,返回 -1
	if this.NowLens == 0 {
		return -1
	}

	// 看是否能找到这个值
	// 找不到直接返回-1
	res,ok := this.MapKeyValueAddr[key]
	if !ok {
		return -1
	}

	// 如果能找到 把这个挪到第一位

	// 如果在头,直接返回其值,顺序啥的都不变
	if this.Head == res {
		return res.Val
	}

	// 如果不在头, 先定义一个 头节点
	newHead := &ListNode{
		Val: res.Val,
		Key: key,
		Pre: nil,
	}

	// 进行分析和挪动
	this.move(res,newHead,key)
	return res.Val
}


func (this *LRUCache) move(res *ListNode,newHead *ListNode,key int) {

	// 如果在末尾
	if this.Last == res {
		// 把数放到第一位,下一个指针指向之前的head节点,长度因为没有超出,所以把该数的上一位设置成最后一个节点,
		oldHead := this.Head
		oldHead.Pre = newHead
		newHead.Next = oldHead
		res.Pre.Next = nil
		this.MapKeyValueAddr[res.Pre.Key] = res.Pre
		this.Last = res.Pre
		this.MapKeyValueAddr[key] = newHead
		this.Head = newHead
		return
	}

	// 如果在中间, 把该数放到头,它的next指针指向以前的head节点,然后把它之前的上一个节点和它的下一个节点进行相连!
	oldHead := this.Head
	oldHead.Pre = newHead
	newHead.Next = oldHead
	this.MapKeyValueAddr[key].Pre.Next = this.MapKeyValueAddr[key].Next
	this.MapKeyValueAddr[key].Next.Pre = this.MapKeyValueAddr[key].Pre
	this.MapKeyValueAddr[this.MapKeyValueAddr[key].Next.Key] = this.MapKeyValueAddr[key].Next
	this.MapKeyValueAddr[oldHead.Key] = oldHead
	this.MapKeyValueAddr[this.MapKeyValueAddr[key].Pre.Key] = this.MapKeyValueAddr[key].Pre
	this.MapKeyValueAddr[key] = newHead
	this.Head = newHead
	return
}


func (this *LRUCache) Put(key int, value int)  {
	// 如果总高度小于等于0,直接返回,啥也放不了
	if this.Lens <= 0 {
		return
	}

	// 如果当前高度为零,直接放进去,自己既是头节点,也是最后一个节点
	if this.NowLens == 0 {
		newHead := &ListNode{
			Val: value,
			Key: key,
			Pre: nil,
			Next: nil,
		}
		this.Head = newHead
		this.Last = newHead
		this.MapKeyValueAddr[key] = newHead
		this.NowLens++
		return
	}

	// 新放入的key==>value  看是否key存在
	if res,ok := this.MapKeyValueAddr[key];ok {
		// 如果存在

		// 如果在头,替换其值,其他不变
		if this.Head == res {
			res.Val = value
			this.MapKeyValueAddr[key] = res
			return
		}
		newHead := &ListNode{
			Val: value,
			Key: key,
			Pre: nil,
		}

		// 进行判断和挪动
		this.move(res,newHead,key)
		return
	}

	// 如果当前的key没有找到,表示以前没有该值
	// 那么将此值放到链表头位置
	// 判断链表总长度,如果超出了,就删掉最后一个

	this.NowLens++

	newHead := &ListNode{
		Val: value,
		Key: key,
		Pre: nil,
	}
	// 将新值放到头位置
	oldHead := this.Head
	oldHead.Pre = newHead
	newHead.Next = oldHead
	this.Head = newHead
	this.MapKeyValueAddr[oldHead.Key] = oldHead
	this.MapKeyValueAddr[key] = newHead

	// 如果高度超出了,就删掉最后一个节点,之前最后一个节点的上一个节点设置成最后一个节点
	if this.NowLens > this.Lens {
		this.Last.Pre.Next = nil
		delete(this.MapKeyValueAddr,this.Last.Key)
		this.Last = this.Last.Pre
		this.NowLens--
	}
	return
}

