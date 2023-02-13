func countOdds(low int, high int) int {
    if high-low == 0 {
        return 0+(high%2 | low%2)
    }
    return (high-low)/2+(high%2 | low%2)
}
