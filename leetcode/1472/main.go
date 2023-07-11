/* Idea: Store history as doubly-linked list.

*/
type SiteNode struct {
    URL string
    Prev *SiteNode
    Next *SiteNode
}

type BrowserHistory struct {
    Current *SiteNode
}


func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{
        Current: &SiteNode{
            URL: homepage,
        },
    }
}


func (this *BrowserHistory) Visit(url string)  {
    newSite := &SiteNode{
        URL: url,
        Prev: this.Current,
    }
    this.Current.Next = newSite
    this.Current = newSite
    return
}


func (this *BrowserHistory) Back(steps int) string {
    for i:=0; i< steps; i++{
        if this.Current.Prev == nil {
            break
        }
        this.Current = this.Current.Prev
    }
    return this.Current.URL
}


func (this *BrowserHistory) Forward(steps int) string {
    for i:=0; i< steps; i++{
        if this.Current.Next == nil {
            break
        }
        this.Current = this.Current.Next
    }
    return this.Current.URL
    
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
