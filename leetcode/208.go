package main

import "fmt"

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{
		[26]*Trie{},
		false,
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, c := range word {
		//找得到当前字符，当前指针指向下一个子节点
		cindex := c - 'a'
		//如果找不到，就插入
		if cur.children[cindex] == nil {
			cur.children[cindex] = &Trie{}
		}
		cur = cur.children[cindex]
	}
	//最后设置标志位
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, c := range word {
		cindex := c - 'a'
		//如果字母不在字典中，说明没有存
		if cur.children[cindex] == nil {
			return false
		}
		//没遍历完，则查找下一个字母
		cur = cur.children[cindex]
	}
	//注意判空
	return cur != nil && cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, c := range prefix {
		cindex := c - 'a'
		//如果字母不在字典中，说明没有存
		if cur.children[cindex] == nil {
			return false
		}
		//没遍历完，则查找下一个字母
		cur = cur.children[cindex]
	}
	//遍历完prefix，说明都找到了
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // 返回 True
	fmt.Println(trie.Search("app"))     // 返回 False
	fmt.Println(trie.StartsWith("app")) // 返回 True
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}
