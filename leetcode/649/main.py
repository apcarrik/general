class Solution:
    def predictPartyVictory(self, senate: str) -> str:
        new_senate: string = ""
        r_count: int = 0
        r_found: bool = False
        d_count: int = 0
        d_found: bool = False
        for party in senate:
            if party == "R":
                r_count += 1
                d_count -= 1
                if r_count > 0:
                    new_senate += "R"
                    r_found= True
            else:
                d_count += 1
                r_count -= 1
                if d_count > 0:
                    new_senate += "D"
                    d_found = True
        # reconcile negative r or d count
        if r_count < 0:
            while r_count < 0:
                new_senate = new_senate.replace("R","",1)
                r_count += 1
        elif d_count < 0:
            while d_count < 0:
                new_senate = new_senate.replace("D","",1)
                d_count += 1
                
        if r_found and not d_found:
            return "Radiant"
        elif d_found and not r_found:
            return "Dire"

        return self.predictPartyVictory(new_senate)
