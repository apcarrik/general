# class SnapshotArray:

#     def __init__(self, length: int):
        

#     def set(self, index: int, val: int) -> None:
        

#     def snap(self) -> int:
        

#     def get(self, index: int, snap_id: int) -> int:

class SnapshotArray:
    def __init__(self, length: int):
        self.arraylen: int = length
        self.changelog: dict[int, int] = {}
        self.snapshots: list[dict[int, int]] = []

    def set(self, index: int, val: int):
        if index < self.arraylen:             
            self.changelog[index] = val

    def snap(self) -> int:
        self.snapshots.append(self.changelog)
        self.changelog = {}
        return len(self.snapshots) - 1
    
    def get(self, index: int, snap_id: int) -> int:
        if snap_id >= len(self.snapshots) or index >= self.arraylen:
            return None
        for cl_idx in range(snap_id, -1, -1):
            if index in self.snapshots[cl_idx]:
                return self.snapshots[cl_idx][index]
        return 0
        


# Your SnapshotArray object will be instantiated and called as such:
# obj = SnapshotArray(length)
# obj.set(index,val)
# param_2 = obj.snap()
# param_3 = obj.get(index,snap_id)
