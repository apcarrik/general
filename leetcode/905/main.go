
// import "fmt"

type Basket struct {
    first int // first index seen
    last int // last index seen
    fruit int
}

func totalFruit(fruits []int) int {
    // Idea: iterate through each fruit tree. Keep track of "baskets" 1 and 2, storing the type of fruit in that basket and the index of the tree you first saw the fruit. Once you see a new type of fruit, check if this is the maximum number of fruits picked so far and record if so, then replace basket 1 with basket 2's contents, and replace basket 2 with this fruit type and index. Basically, it's a queue of length 2.
    if len(fruits) < 3 {
        return len(fruits)
    }
    var b1, b2 Basket
    b1.first, b1.last = 0,0
    b1.fruit = fruits[0]
    b2.first, b2.last = 1,1
    b2.fruit = fruits[1]
    if b1.fruit == b2.fruit {
        b1.last = b2.last
    }
    max := 2
    for i:=2; i<len(fruits); i++ {
        if fruits[i] != b1.fruit && fruits[i] != b2.fruit {
            if b1.fruit == b2.fruit {
                b2.first, b2.last = i,i
                b2.fruit = fruits[i]
            } else {
                //fmt.Println("i:",i,"f:",fruits[i],"b1:",b1,"b2:",b2)
                if i - b1.first > max {
                    //fmt.Println("i-b1.first:", i - b1.first)
                    max = i - b1.first
                }
                if fruits[i-1] == b1.fruit{
                    b1.first = b2.last+1
                    b1.last = i-1
                } else {
                    b1.fruit = b2.fruit
                    b1.first = b1.last +1
                    b1.last = i-1
                }
                b2.fruit = fruits[i]
                b2.first, b2.last = i,i
                //fmt.Println("i:",i,"f:",fruits[i],"b1:",b1,"b2:",b2)
            }
        } 
        if fruits[i] == b1.fruit {
            //fmt.Println("i=",i,"b1.last=i")
            b1.last = i
        } 
        if fruits[i] == b2.fruit {
            //fmt.Println("i=",i,"b1=", b1, "b2=", b2, "b2.last=i")
            b2.last = i
        }
    }
    if len(fruits) - b1.first > max {
        //fmt.Println("i:",len(fruits),"f:",fruits[len(fruits)-1],"b1:",b1,"b2:",b2,"i-b1.first:", len(fruits) - b1.first)
        max = len(fruits) - b1.first
    }
    return max
}
