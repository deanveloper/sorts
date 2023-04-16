fn main() {
    println!("Hello, world!");
}

pub mod sort {
    pub mod merge {
        pub fn sort<T: PartialOrd<T> + Copy>(slice: &mut [T]) {
            if slice.len() < 1 {
                return;
            }

            let mut left: Vec<T> = Vec::with_capacity(slice.len() / 2);
            let mut right: Vec<T> = Vec::with_capacity(slice.len() / 2 + 1);
            for i in 0..slice.len() / 2 {
                left.push(slice[i]);
            }
            for i in slice.len() / 2..slice.len() {
                right.push(slice[i]);
            }

            sort(&mut left);
            sort(&mut right);

            let mut l = 0;
            let mut r = 0;
            for i in 0..slice.len() {
                if r >= right.len() {
                    slice[i] = left[l];
                    l += 1;
                    continue;
                }
                if l >= left.len() {
                    slice[i] = right[r];
                    r += 1;
                    continue;
                }
                if left[l] < right[r] {
                    slice[i] = left[l];
                    l += 1;
                } else {
                    slice[i] = right[r];
                    r += 1;
                }
            }
        }
    }
    pub mod quick {
        pub fn sort<T: PartialOrd<T> + Copy>(mut slice: &mut [T]) {
            if slice.len() <= 1 {
                return;
            }

            let p = partition(&mut slice);

            let (left, right) = slice.split_at_mut(p);
            sort(left);
            sort(right);
        }

        pub fn partition<T: PartialOrd<T> + Copy>(slice: &mut [T]) -> usize {
            if slice.len() <= 1 {
                return slice.len() - 1;
            }

            let mut left = 0;
            let right = slice.len() - 1;

            for i in 0..slice.len() {
                if slice[i] < slice[right] {
                    slice.swap(i, left);
                    left += 1;
                }
            }
            slice.swap(left, right);

            left
        }
    }
}
