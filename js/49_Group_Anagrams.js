/**
 * @param {string[]} strs
 * @return {string[][]}
 */
var groupAnagrams = function (strs) {
  const groups = new Map();
  for (let i = 0; i < strs.length; i++) {
    const sorted = strs[i].split("").sort().join("");
    let updatedGroup;
    if (groups.has(sorted)) {
      updatedGroup = [...groups.get(sorted), strs[i]];
    } else {
      updatedGroup = [strs[i]];
    }
    groups.set(sorted, updatedGroup);
  }
  return Array.from(groups.values());
};

/**
 * @param {string[]} strs
 * @return {string[][]}
 */
var groupAnagrams = function (strs) {
  const res = {};

  for (let i = 0; i < strs.length; i++) {
    const count = new Array(26).fill(0);
    const baseCharCode = "a".charCodeAt(0);

    for (const char of strs[i]) {
      count[char.charCodeAt(0) - baseCharCode]++;
    }

    let sortedString = "";
    for (let i = 0; i < 26; i++) {
      sortedString += String.fromCharCode(baseCharCode + i).repeat(count[i]);
    }

    if (res[sortedString]) {
      res[sortedString].push(strs[i]);
    } else {
      res[sortedString] = [strs[i]];
    }
  }

  return Array.from(Object.values(res));
};
