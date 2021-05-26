'''
Solution for problem 9.2 in Cracking the Coding Interview

'''

# Database function placeholders


# Adds an entry to friend database, returns a unique id for new friend
def addToDatabase():
    return 0

# Returns the name associated with a profile ID
def searchProfileName(profileID):
    return ""

# Data structures for people & social network

class Profile:
    def __init__(self, name):
        self.name = name
        self.id = addToDatabase()
        self.friends = []

    # Handles editing profile info
    def editProfileInfo(self, school, job, orginizations):
        if(school != None):
            self.school = school
        if(job != None):
            self.job = job
        if(orginizations != None):
            self.orginizations = orginizations

    # Adds a new friend given friend unique id
    def addFriend(self,friendId):
        self.friends.append(friendId)

    # Returns list of names of friends
    def getFriendNames(self):
        friendNames = []
        for friendId in self.Friends:
                friendNames.append( searchProfileName(friendId))
        return friendNames

    # Returns list of friend unique IDs
    def getFriendIds(self):
        return self.friends

# Queue structure for friendqueue
class friendNode:
    def __init__(self, data, friends):
        self.data = data
        self.friends = friends
        self.next = None

    def setNext(self, nextNode):
        self.next = nextNode

class friendQueue:
    def __init__(self):
        self.head = None

    def isEmpty(self):
        return True if self.head != None else False

    def add(self, item, friends):
        newnode = friendNode(item, friends)
        if(self.head == None):
            self.head = newnode
        else:
            nextnode = self.head
            while (nextnode.next != None):
                nextnode = nextnode.next
            nextnode.setNext(newnode)

    def pop(self):
        retnode = self.head
        self.head = self.head.next
        return retnode.data, retnode.friends

# shortestFriendPath finds the least number of people between two profiles in a given social network
def shortestFriendPath(profileId1, profileId2):
    # Perform BFS on first profile friends, checking if 2nd profile is there. If not, repeat for all
    # friends of first profile, recursively, keeping track of friend path up to that point
    friendpath = [profileId1]
    friendqueue = friendQueue()
    friendqueue.add(profileId1, friendpath)
    while(~friendqueue.isEmpty()):
        friendId, friendpath = friendqueue.pop()
        if(friendId == profileId2):
            return friendpath.append[profileId2]
        else:
            friendqueue.add(friendId,friendpath.append(friendId))
    return None

# Testing shortestPath
# TODO: create tests for shortestPath