class Solution:
    def findReplaceString(self, s: str, indices: List[int], sources: List[str], targets: List[str]) -> str:
        offset: int = 0
        for i,c,t in sorted(zip(indices,sources,targets)):
            i += offset
            if i >= 0 and i+len(c) <= len(s) and c == s[i:i+len(c)]:
                s = s[:i] + t + (s[i+len(c):] if i+len(c) < len(s) else "")
                offset += len(t) - len(c)
        return s
