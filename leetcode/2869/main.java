class Solution {
    public int minOperations(List<Integer> nums, int k) {
        // idea: keep a set of ints from 1:k and working backwards through the array, pop the end of nums and check if it matches something in the set. If so, remove it from the set and return the iteration number if the set is empty.
        // initialize set
        Set<Integer> set = new HashSet<Integer> (); 
        for (int i = 1; i <= k; i++) {
            set.add(i);
        }

        // iterate through nums from tail to head
        int iter = 1;
        for (; iter <= nums.size(); iter++) {
            if (set.remove(nums.get(nums.size() - iter)) && set.size() == 0) {
                break;
            }
        }

        // return the number of iterations it took to collect all values from 1:k
        return iter;
    }
}
