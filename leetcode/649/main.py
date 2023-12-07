class Solution:
    def predictPartyVictory(self, senate: str) -> str:
        new_senate: string = ""
        r_count: int = 0
        d_count: int = 0
        for party in senate:
            if party == "R":
                r_count += 1
                d_count -= 1
                if r_count > 0:
                    new_senate += "R"
            else:
                d_count += 1
                r_count -= 1
                if d_count > 0:
                    new_senate += "D"
        # reconcile negative r or d count
        if r_count < 0:
            while r_count < 0:
                new_senate = new_senate.replace("R","",1)
                r_count += 1
        elif d_count < 0:
            while d_count < 0:
                new_senate = new_senate.replace("D","",1)
                d_count += 1

        if new_senate.count("R") > 0 and new_senate.count("D") == 0:
            return "Radiant"
        elif new_senate.count("D") > 0 and new_senate.count("R") == 0:
            return "Dire"

        return self.predictPartyVictory(new_senate)
