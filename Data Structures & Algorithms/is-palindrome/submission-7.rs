impl Solution {
    pub fn is_palindrome(s: String) -> bool {
        let mut chars: Vec<char> = Vec::new();

        for c in s.chars() {
            if c.is_alphanumeric(){
                chars.push(c.to_ascii_lowercase());
            }
        }

        if chars.is_empty() {
            return true
        }
        
        let mut left: usize = 0;
        let mut right: usize = chars.len() - 1;

        while left < right {
            if chars[left] != chars[right] {
                return false;
            }
            left += 1;
            right -= 1;

        }
        return true;
    }
}
