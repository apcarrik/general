// type MyHashSet struct {
    
// }


// func Constructor() MyHashSet {
    
// }


// func (this *MyHashSet) Add(key int)  {
    
// }


// func (this *MyHashSet) Remove(key int)  {
    
// }


// func (this *MyHashSet) Contains(key int) bool {
    
// }

type MyHashSet struct {
	n int
	table [][]int
}


func Constructor() MyHashSet {
	mhs := new(MyHashSet)
	mhs.n = 50000
	mhs.table = [][]int{}
	for i := 0; i < mhs.n; i++ {
		mhs.table = append(mhs.table, []int{})
	}
	return *mhs
}

func (this *MyHashSet) hashFunc(key int) int {
	return key % this.n
}

func (this *MyHashSet) Add(key int) {
	items := this.table[this.hashFunc(key)]
	for _, item := range items {
		if item == key {
			return 
		}
	}
	this.table[this.hashFunc(key)] = append(this.table[this.hashFunc(key)], key)	
}

func (this *MyHashSet) Remove(key int) {
	items := this.table[this.hashFunc(key)]
	for i, item := range items {
		if item == key {
			if len(items) == 1 {
				this.table[this.hashFunc(key)] = []int{}
			} else {
				this.table[this.hashFunc(key)][i] = items[len(items)-1]
				this.table[this.hashFunc(key)] = items[:len(items)-1]
			}
		}
	}

}

func (this *MyHashSet) Contains(key int) bool {
	items := this.table[this.hashFunc(key)]
	for _, item := range items {
		if item == key {
			return true
		}
	}
	return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
