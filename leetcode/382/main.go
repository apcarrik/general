/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import (
    "math/rand"
    "time"
    //"fmt"
)

 // Idea: store length of LL and a slice list of runners with pointers to every 100th node of LL
type Solution struct {
    Head *ListNode
    Runners []*ListNode
    RunnerSpacing int
    Len int
}


func Constructor(head *ListNode) Solution {
    // Find length and store it in Len
    listLen:=0
    runners:= []*ListNode{}
    iter := head
    runnerSpacing := 25
    for iter != nil {
        if listLen % runnerSpacing == 0{
            runners = append(runners,iter)
        }
        listLen++
        iter=iter.Next
    }
    //fmt.Println("len(runners):", len(runners))
    return Solution{Head:head, Runners:runners, RunnerSpacing: runnerSpacing, Len:listLen}
}


func (this *Solution) GetRandom() int {
    // seed random number
    rand.Seed(time.Now().UnixNano())
    // find random number between 0,LL.length. return that node index
    rando := rand.Intn(this.Len)    
    runnerIdx := rando / this.RunnerSpacing
    iter := this.Runners[runnerIdx]
    //fmt.Println(rando, rando / this.RunnerSpacing)
    for i:=0; i < rando - (runnerIdx*this.RunnerSpacing); i++{
        if iter.Next == nil {
            //fmt.Println(">",i)
        }
        iter=iter.Next
    }

    return iter.Val
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
