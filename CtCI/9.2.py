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

# TODO: decide what data structure to use for people & social network

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


#TODO: create queue structure for friendqueue

# shortestFriendPath finds the least number of people between two profiles in a given social network
def shortestFriendPath(profileId1, profileId2):
    # TODO: Implement


# Testing shortestPath
# TODO: create tests for shortestPath