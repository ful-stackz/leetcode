namespace Leetcode.ThreeSum;

public class Solution {
    struct SumMember {
        public int Index { get; set; }
        public int Value { get; set; }
    }

    public IList<IList<int>> ThreeSum(int[] nums) {
        Array.Sort(nums);

        var bases = new List<(SumMember a, SumMember b)>();
        for (int i = 0; i < nums.Length - 1; i++) {
            for (int j = i + 1; j < nums.Length; j++) {
                bases.Add((
                    new SumMember { Index = i, Value = nums[i] },
                    new SumMember { Index = j, Value = nums[j] }
                ));
            }
        }

        var triplets = new HashSet<(int a, int b, int c)>(bases.Count);
        foreach (var ab in bases) {
            var cValue = (ab.a.Value + ab.b.Value) * -1;
            var cIndex = Array.BinarySearch(nums, cValue);
            if (cIndex < 0 || cIndex == ab.a.Index || cIndex == ab.b.Index) continue;

            var triplet = (ab.a.Value, ab.b.Value, cValue);
            triplets.Add(Sort(ref triplet));
        }

        IList<IList<int>> result = new List<IList<int>>();
        foreach (var triplet in triplets) {
            result.Add(new List<int> { triplet.a, triplet.b, triplet.c });
        }

        return result;

        static (int, int, int) Sort(ref (int a, int b, int c) tpl) {
            var arr = new [] { tpl.a, tpl.b, tpl.c };
            Array.Sort(arr);
            return (arr[0], arr[1], arr[2]);
        }
    }
}
