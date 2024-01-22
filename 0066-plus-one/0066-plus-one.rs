impl Solution {
  pub fn plus_one(mut digits: Vec<i32>) -> Vec<i32> {
    let mut result: Vec<i32> = digits;
    let mut carry: i32 = 1;

    for i in (0..result.len()).rev() {
        let sum: i32 = result[i] + carry;
        result[i] = sum % 10;
        carry = sum / 10;
        if carry == 0 {
            break;
        }
    }

    if carry > 0 {
        result.insert(0, carry);
    }

    result
    }
}
