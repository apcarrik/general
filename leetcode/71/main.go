func simplifyPath(path string) string {
	new := "/"
	for i:=1; i<len(path); i++ {
		switch path[i] {
		case '/' : 
			continue
		case '.' : 
			// check if double period, single, or more
			if i < len(path) - 1 {
                if path[i+1] == '.' {
                    // at this point, we know we have two periods. The two cases at this point are: 
                    //  1. it is the end of string or the next character is a slash, and 
                    //  2. there are more characters and they are not slashes                    
                    if i==len(path)-2 || path[i+2] == '/' {
                        // case 1 - move back 1 directory
                        if len(new) > 1 {
                            for j := len(new)-2; j>=0; j-- {
                                if j == 0 {
                                    new = "/"
                                    break
                                } else if new[j] == '/' {
                                    new = new[:j]
                                    break
                                }
                            }
                        }
                        i += 2
                    } else {
                        // case 2 - treat as new directory name
                        if new[len(new)-1] != '/' {
                            new += "/"
                        }
                        for j := i; j<len(path); j++{
                            if path[j] == '/' {
                                new += path[i:j]
                                i = j	
                                break					
                            } else if j == len(path) -1 {
                                new += path[i:]
                                i = j	
                                break		
                            }
                        }
                    }                
                } else if path[i+1] != '/' {
                    // treat as new directory name
                    if new[len(new)-1] != '/' {
                        new += "/"
                    }
                    for j := i; j<len(path); j++{
                        if path[j] == '/' {
                            new += path[i:j]
                            i = j		
                            break					
                        } else if j == len(path) -1 {
                            new += path[i:]
                            i = j		
                            break		
                        }
                    }
                }
            }
		default: // any alphanumeric, get new directory name
			if new[len(new)-1] != '/' {
                new += "/"
            }
			for j := i; j<len(path); j++{
				if path[j] == '/' {
		            new += path[i:j]
                    i = j		
		            break					
	            } else if j == len(path) -1 {
		            new += path[i:]
                    i = j		
		            break		
				}
			}
		}
	}

	return new

}
