use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut memory : HashMap<i32, usize> = HashMap::new();
        for (i, &num) in nums.iter().enumerate() {
            if let Some(&compl_idx) = memory.get(&num) {
                return vec![compl_idx as i32, i as i32];
            }
            memory.insert(target-num , i);
        }
        return vec![-1,-1];
    }
}
