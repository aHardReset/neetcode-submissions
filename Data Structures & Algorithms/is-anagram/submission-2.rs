impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        if s.chars().count() != t.chars().count() {
            return false;
        }

        let mut alphabet = [0i32; 26];

        for (s_char, t_char) in s.bytes().zip(t.bytes()) {
            let s_idx = (s_char - b'a') as usize;
            let t_idx = (t_char - b'a') as usize;
            alphabet[s_idx] += 1;
            alphabet[t_idx] -= 1;
        }

        Solution::is_all_zeros(alphabet)
    }

    pub fn is_all_zeros(numbers: [i32; 26]) -> bool {
        for number in numbers {
            if number != 0 {
                return false
            }
        }
        true
    }
}
