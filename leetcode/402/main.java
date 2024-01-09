class Solution {
	public String removeKdigits(String num, int k) {
        int r = num.length() - k;
		String output = "0";
		int last_min_idx = -1;
		for (int i=0; i<r; i++) {
			int min = 9;
			for (int j=last_min_idx+1; j<num.length()-r+i+1; j++){
				int n = Character.getNumericValue(num.charAt(j)); // convert num[j] to int
				if (n == 0) {
					min = n;
					last_min_idx = j;
					break;
				} else if (n < min) {
					min = n;
					last_min_idx = j;
				}
			}
			// append min as char to output
			output += Integer.toString(min);
		}
        // purge leading 0's from output
        // output = Integer.toString(Integer.parseInt(output));
        int first_nonzero_idx = 0;
        for (int i=0; i<output.length(); i++) {
            if (Character.getNumericValue(output.charAt(i)) != 0) {
                break;
            }
            first_nonzero_idx++;
        }
        output = output.substring(first_nonzero_idx);
        // check if output is empty
        if (output.length() == 0 ) {
            output = "0";
        }
		return output;
	}
}
