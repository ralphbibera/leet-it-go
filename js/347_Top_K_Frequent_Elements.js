/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var topKFrequent = function (nums, k) {
  const frequencyMap = new Map();
  for (let i = 0; i < nums.length; i++) {
    frequencyMap.set(nums[i], (frequencyMap.get(nums[i]) ?? 0) + 1);
  }
  const sorted = Array.from(frequencyMap.entries()).sort((a, b) => b[1] - a[1]);
  return sorted.map(([key]) => key).slice(0, k);
};
