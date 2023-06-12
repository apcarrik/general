class Solution:
    def letterCasePermutation(self, s: str) -> List[str]:
        superset: list[str] = [""]
        for c in s:
            lowers: list[str] = [st+c.lower() for st in superset]
            uppers: list[str] = [st+c.upper() for st in superset]
            superset = lowers + uppers if c.isalpha() else lowers
        return superset
