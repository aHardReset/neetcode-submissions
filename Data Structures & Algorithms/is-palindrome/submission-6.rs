impl Solution {
    pub fn is_palindrome(s: String) -> bool {
        let mut s1 = String::from("");

        for c in s.chars() {
            if c.is_alphanumeric(){
                s1.push(c.to_lowercase().to_string().chars().next().expect("Error"));
            }
        }

        let s1_vec: Vec<char> = s1.chars().collect();

        if s1_vec.len() <= 0 {
            return true
        }
        
        let mut left: usize = 0;
        let mut right: usize = s1.len() - 1;

        while left < right {
            if s1_vec[left] != s1_vec[right] {
                return false;
            }
            left += 1;
            right -= 1;

        }
        return true;
    }
}
