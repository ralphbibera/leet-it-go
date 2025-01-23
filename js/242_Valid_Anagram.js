/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function (s, t) {
  return s.split("").sort().join("") === t.split("").sort().join("");
};

/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function (s, t) {
  if (s.length !== t.length) return false;

  const count = Array(26).fill(0);

  for (const char of s) {
    count[char.charCodeAt(0) - "a".charCodeAt(0)] += 1;
  }

  for (const char of t) {
    if (count[char.charCodeAt(0) - "a".charCodeAt(0)] === 0) return false;
    count[char.charCodeAt(0) - "a".charCodeAt(0)] -= 1;
  }

  return true;
};
