/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function (nums) {
  const hs = {};
  for (const num of nums) {
    hs[num] = (hs[num] ?? 0) + 1;
    if (hs[num] === 2) {
      return true;
    }
  }
  return false;
};
