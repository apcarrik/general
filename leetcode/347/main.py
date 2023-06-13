class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        # create frequency map
        fmap: dict[int,int] = {}
        for num in nums:
            if num in fmap:
                fmap[num] += 1
            else:
                fmap[num] = 1

        # find top k elements by min heap
        topk: list[int] = []
        heapq.heapify(topk)
        for (elem, freq) in fmap.items():
            if len(topk) < k:
                heapq.heappush(topk, (freq,elem))
            else:
                if freq > topk[0][0]:
                    heapq.heappop(topk)
                    heapq.heappush(topk, (freq,elem))
        return [elem for _,elem in topk]
