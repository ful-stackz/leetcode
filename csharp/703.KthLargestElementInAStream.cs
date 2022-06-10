namespace Leetcode.KthLargestElementInAStream;

public class KthLargest {
    private readonly int _k;
    private readonly List<int> _buffer;

    public KthLargest(int k, int[] nums) {
        _k = k - 1;
        _buffer = new List<int>(nums);
        _buffer.Sort();
        _buffer.Reverse();
    }

    public int Add(int val) {
        if (_buffer.Count == 0) { // first value
            _buffer.Add(val);
        }
        else if (_buffer[0] < val) { // new largest value
            _buffer.Insert(0, val);
        }
        else if (_buffer[^1] > val) { // new smallest value
            _buffer.Add(val);
        }
        else {
            for (int i = 0; i < _buffer.Count; i++) {
                if (_buffer[i] <= val) {
                    _buffer.Insert(i, val);
                    break;
                }
            }
        }

        return _buffer[_k];
    }
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * KthLargest obj = new KthLargest(k, nums);
 * int param_1 = obj.Add(val);
 */
